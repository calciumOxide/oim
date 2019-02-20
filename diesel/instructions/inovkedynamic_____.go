package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_inovkedynamic struct {
}

func init()  {
	INSTRUCTION_MAP[0x74] = &I_inovkedynamic{}
}

func (s I_inovkedynamic)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "inovkedynamic exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jint(value.(types.Jint) * -1))
	return nil
}

func (s I_inovkedynamic)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jfloat(9.123456789012345))
	f.PushFrame(types.Jint(9))
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		调用动态方法
======================================================================================
						||		inovkedynamic
						||------------------------------------------------------------
						||		indexbyte1
						||------------------------------------------------------------
						||		indexbyte2
		格式				||------------------------------------------------------------
						||		0
						||------------------------------------------------------------
						||		0
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		invokedynamic = 186 (0xba)
======================================================================================
						||		...，[arg1， [arg2 ...]]→
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
		描述				||		代码中每条 invokedynamic 指令出现的位置都被称为一个动态调用点 (Dynamic Call Site)。
						||		首先。无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6) 的运行时常量池的索引值，
						||		构建方式为(indexbyte1 << 8)| indexbyte2,
						||		该索引所指向的运行时常量池项应当是一个调用点限定符(§5.1)的符号引 用。指令第 3、4 个操作数固定为 0。
						||		调用点限定符会被解析(§5.4.3.6)为一个动态调用点，从中可以获取到
						||		java.lang.invoke.MethodHandle 实例的引用、java.lang.invoke.MethodType 实例的引用，和所涉及的静态参数的引用。
						||		接着，作为调用点限定符解析过程的一部分，引导方法将会被执行。
						||		如同使用 invokevirtual 指令调用普通方法那样，会包含一个运行时常量池的索引指 向一个带有如下属性的方法(§5.1):
						||			 此方法名为 invoke。
						||			 此方法描述符中的返回值是 java.lang.invoke.CallSite。
						||			 此方法描述符中的参数来源自操作数栈中的元素，包括如下顺序排列的 4 个参数:
						||				 java.lang.invoke.MethodHandle
						||				 java.lang.invoke.MethodHandles.Lookup
						||				 String
						||				 java.lang.invoke.MethodType
						||			 如果调用点限定符有静态参数，那么这些参数的参数类型应附加在方法描 述符的参数类型中，以便在调用时按顺序入栈至操作数栈。静态参数的参 数类型可以包括:
						||				 Class
						||				 java.lang.invoke.MethodHandle
						||				 java.lang.invoke.MethodType
						||				 String
						||				 int
						||				 long
						||				 float
						||				 double
						||			 在 java.lang.invoke.MethodHandle 之中可以提供关于在哪个类中 能找到方法的符号引用所对应的实际方法的信息。
						||		引导方法执行前，下面各项内容将会按顺序压入到操作数栈中:
						||			 用于代表引导方法的 java.lang.invoke.MethodHandle 对象的引用。
						||			 用于确定动态调用点发生位置的java.lang.invoke.MethodHandles.Lookup 对象的引用。
						||			 用于确定调用点限定符中方法名的 String 对象引用。
						||			 在方法限定符中出现的各种静态参数，包括:类、方法类型、方法句柄、字符串以及各种数值类型(§2.3.1, §2.3.2)都必须按照他们在方法 限定符中出现的顺序依次入栈(此处基本类型不会发生自动装箱)。
						||		只要引导方法能够被正确调用，它的描述符可以是不精确的。举个例子，引导 方法的第一个参数应该是 java.lang.invoke.MethodHandles.Lookup， 但是在可以使用 Object 来代替，返回值应该是 java.lang.invoke.CallSite，也可以使用 Object 来代替。
						||		如果引导方法是一个变长参数方法(Variable Arity Method)，那某些(甚 至是全部)上面描述中要压入到操作数栈的参数会被包含在一个数组参数之 中。
						||		引导方法的调用发生在试图解析动态方法调用点的调用点限定符的那条线程 上，如果同时有多条线程进行此操作，那引导方法将会被并发调用。因此，如 果引导方法中有访问公有数据的话，需要注意多线程竞争问题，对公有数据访 问施行适当的保护措施。
						||		引导方法执行后的返回值是一个 java.lang.invoke.CallSite 或其子类 的实例，这个对象被称为调用点对象(Call Site Object)，此对象的引用 将会从操作数栈中出栈，就像 invokevirtual 指令执行过程一样。
						||		如果多条线程同时执行了一个动态调用点的引导方法，那 Java 虚拟机必须选 择其中一个引导方法的返回值作为调用点对象，并将其发布到所有线程之中。
						||		此动态调用点中其余的引导方法会完成整个执行过程，但是它们的返回结果将 被忽略掉，转为使用哪个被 Java 虚拟机选中的调用点对象来继续执行。
						||		调用点对象拥有一个类型描述符(一个 java.lang.invoke.MethodType 的实例)，它必须语义上等同于调用点限定符中方法描述符内所包含的 java.lang.invoke.MethodType 对象。
						||		调用点限定符解析的结果是一个调用点对象，此对象将会与它的动态调用点永 久绑定起来。
						||		绑定于动态调用点的调用点对象所表示的方法句柄将会被调用，这次调用就和 执行 invokevirtual 指令一样，
						||		会带有一个指向运行时常量池的索引，它指 向的常量池项是一个方法的符号引用，此方法具备如下属性:
						||			 方法名为 invokeExact。
						||			 方法描述符为调用点限定符中包含的描述符。
						||			 由 java.lang.invoke.MethodHandle 来确定在哪个类中查找方法的 符号引用所对应的方法。
						||		指令执行时，操作数栈中的内容会被虚拟机解释为包含一个调用点对象的引用 以及跟随 nargs 个参数值，这些参数的数量、类型和顺序都必须与调用点限 定符中的方法描述符保持一致。
						||
						||
						||
						||
						||
						||
======================================================================================
						||		
						||		如果调用点限定符的符号引用解析过程中抛出了异常 E，那 invokedynamic 指令必须抛出包装着异常 E 的 BootstrapMethodError 异常。
						||		另外，在调用点限定符的后续解析过程中，如果引导方法执行过程因异常 E 而 异常退出(§2.6.5)，那 invokedynamic 指令必须抛出包装着异常 E 的 BootstrapMethodError 异常。
	   链接时异常			||		(这可能是由于引导方式有错误的参数长度、参数类型或者返回值而导致 java.lang.invoke.MethodHandle.invoke 方法抛出了 java.lang.invoke.WrongMethodTypeException 异常。)
						||		另外，在调用点限定符的后续解析过程中，如果引导方法的返回值不是一个 java.lang.invoke.CallSite 的实例，那 invokedynamic 指令必须抛 出 BootstrapMethodError 异常。
						||		另外，在调用点限定符的后续解析过程中，如果调用点对象的目标的类型描述 符与方法限定符中所包括的方法描述符不一致，那 invokedynamic 指令必须 抛出 BootstrapMethodError 异常。
						||
======================================================================================
						||
						||		如果动态调用点的调用点限定符解析过程成功完成，那就意味着将有一个非空 的 java.lang.invoke.CallSite 的实例绑定到该动态调用点之上。
		运行时异常		||		因此， 操作数栈中表示调用点目标的对象不会为空，这也意味着，调用点限定符中的 方法描述符与等效于被 invokevirtual 指令所调用方法句柄那个方法句柄 的类型描述符语义上是一致的。
	   					||		上面描述的意思是已经绑定了调用点对象的 invokedynamic 指令，永远不可 能抛出 NullPointerException 异常或者 java.lang.invoke.WrongMethodTypeException 异常。
						||
======================================================================================
						||
						||
						||
		注意				||
						||
						||
						||
======================================================================================
 */