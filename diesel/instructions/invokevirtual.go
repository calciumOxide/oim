package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
		"../variator"
	"../../loader/clazz"
	"../../loader/clazz/item"
	"reflect"
)

type I_invokevirtual struct {
}

func init()  {
	INSTRUCTION_MAP[0xb6] = &I_invokevirtual{}
}

func (s I_invokevirtual)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "invokevirtual exce >>>>>>>>>\n")

	methodIndex := (uint16(ctx.Code[ctx.PC]) << 8) | uint16(ctx.Code[ctx.PC + 1])
	ctx.PC += 2

	methodInterface, _ := ctx.Clazz.GetConstant(methodIndex)
	classCp, _ := ctx.Clazz.GetConstant(methodInterface.Info.(*item.MethodRef).ClassIndex)
	classNameCp, _ := ctx.Clazz.GetConstant(classCp.Info.(*item.Class).NameIndex)
	class := clazz.GetClass(classNameCp.Info.(*item.Utf8).Str)
	if !ctx.Cinit(class) {
		return nil
	}

	andType, _ := ctx.Clazz.GetConstant(methodInterface.Info.(*item.MethodRef).NameAndTypeIndex)
	nameCp, _ := ctx.Clazz.GetConstant(andType.Info.(*item.NameAndType).NameIndex)
	method := &clazz.Method{}
	method = nil
	desc := ""
	if nameCp.Info.(*item.Utf8).Str == "<init>" || nameCp.Info.(*item.Utf8).Str == " <clinit>" {
		except, _ := variator.AllocExcept(variator.AbstractMethodError)
		ctx.Throw(except)
		return nil
	}

	descCp, _ := ctx.Clazz.GetConstant(andType.Info.(*item.NameAndType).DescriptorIndex)
	desc = descCp.Info.(*item.Utf8).Str

	method = class.GetMethod(nameCp.Info.(*item.Utf8).Str, descCp.Info.(*item.Utf8).Str)
	if method == nil {
		except, _ := variator.AllocExcept(variator.MethodNotFindException)
		ctx.Throw(except)
		return nil
	}

	if method.IsPrivate() {
		except, _ := variator.AllocExcept(variator.AbstractMethodError)
		ctx.Throw(except)
		return nil
	}

	if method.IsStatic() {
		except, _ := variator.AllocExcept(variator.IncompatibleClassChangeError)
		ctx.Throw(except)
		return nil
	}

	parmsType := method.GetParmsType(desc)
	var args []interface{}
	for i := 0; i < len(parmsType); i++ {
		top, _ := ctx.CurrentFrame.PopFrame()
		args = append(args, top)
	}

	object, _ := ctx.CurrentFrame.PopFrame()

	if object != nil && reflect.TypeOf(object) != reflect.TypeOf(&types.Jreference{}) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	if object == nil || object.(*types.Jreference).Reference == nil {
		except, _ := variator.AllocExcept(variator.NullPointerException)
		ctx.Throw(except)
		return nil
	}
	objectClass := object.(*types.Jreference).ElementType.(*clazz.ClassFile)

	if !objectClass.ExtendOf(class) {
		except, _ := variator.AllocExcept(variator.AbstractMethodError)
		ctx.Throw(except)
		return nil
	}

	if method.IsProtected() {
		if !ctx.Clazz.EqualsPackage(objectClass) && !ctx.Clazz.ExtendOf(objectClass){
			except, _ := variator.AllocExcept(variator.AbstractMethodError)
			ctx.Throw(except)
			return nil
		}
	}

	if !method.CheckParams(parmsType, args) {
		except, _ := variator.AllocExcept(variator.MethodParamsNoMatchExcept)
		ctx.Throw(except)
		return nil
	}
	args = append(args, object)
	ctx.InvokeMethod(class, method, args)
	return nil
}

func (s I_invokevirtual)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jreference{
		Reference: types.Jobject{},
		ElementType: clazz.CLASS_MAP["com/oxide/A"],
	})
	f.PushFrame(types.Jint(1))
	//f.PushFrame(types.Jdouble(9.123456789012345))
	//f.PushFrame(types.Jfloat(9.123456789012345))
	//f.PushFrame(&types.Jreference{
	//	Reference: types.Jobject{},
	//	ElementType: clazz.CLASS_MAP["com/oxide/A"],
	//})
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x0, 0x0, 14, 0x1, 0x0},
		//Code: []byte{0x0, 0x0, 0x1, 0x1, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		调用实例方法，依据实例的类型进行分派
======================================================================================
						||		invokevirtual
						||------------------------------------------------------------
						||		indexbyte1
						||------------------------------------------------------------
						||		indexbyte2
		格式				||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		invokevirtual = 182(0xb6)
======================================================================================
						||		...，objectref，[arg1，[arg2 ...]] →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
		描述				||		无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运 行时常量池的索引值，
						||		构建方式为(indexbyte1 << 8)| indexbyte2，
该索引所指向的运行时常量池项应当是一个方法(§5.1)的符号引用，它包 含了方法的名称和描述符，以及包含该方法的接口的符号引用。
这个方法的符 号引用是已被解析过的(§5.4.3.3)，
而且这个方法不能是实例初始化方法 (§2.9)和类或接口的初始化方法(§2.9)。
最后，如果调用的方法是 protected 的(§4.6)，并且这个方法是当前类的父类成员，
并且这个方法 没有在同一个运行时包(§5.3)中定义过，那 objectref 所指向的对象的 类型必须为当前类或者当前类的子类。

假设 C 是 objectref 所对应的类，虚拟机将按下面规则查找实际执行的方法:
 如果 C 中定义了一个实例方法 M，它了重写(override，§5.4.3.5)
了符号引用中表示的方法，那方法 M 就会被调用，查找过程终止。
 另外，如果 C 有父类，查找过程将按第一点的方式顺序递归搜索 C 的直接
           父类，如果超类中能搜索到符合的方法，那这个方法就会被调用。
 否则，抛出 AbstractMethodError 异常。

在操作数栈中，objectref 之后必须跟随 N 个参数，它们的数量、类型和顺 序都必须与方法描述符所描述的保持一致。

如果要调用的是同步方法，那与 objectref 相关的管程(monitor)将会进 入或者重入，就如当前线程中同执行了 monitorenter 指令一般。
如果要调用的不是本地方法，n 个 args 参数和 objectref 将从操作数栈中 出栈。
法调用的时候，一个新的栈帧将在 Java 虚拟机栈中被创建出来， objectref 和连续的 n 个参数将存放到新栈帧的局部变量表中，
objectref 存为局部变量 0，arg1 存为局部变量 1(如果 arg1 是 long 或 double 类型， 那将占用局部变量 1 和 2 两个位置)，依此类推。
参数中的浮点类型数据在存 入局部变量之前会先进行数值集合转换(§2.8.3)。
新栈帧创建后就成为当 前栈帧，Java 虚拟机的 PC 寄存器被设置成指向调用方法的首条指令，程序就 从这里开始继续执行。

如果要调用的是本地方法，要是这些平台相关的代码尚未绑定(§5.6)到虚 拟机中的话，绑定动作先要完成。
指令执行时，n 个 args 参数和 objectref 将从操作数栈中出栈并作为参数传递给实现此方法的代码。
参数中的浮点类型 数据在传递给调用方法之前会先进行数值集合转换(§2.8.3)。
参数传递和 代码执行都会以具体虚拟机实现相关的方式进行。当这些平台相关的代码返回 时:
 如果这个本地方法是同步方法，那与 objectref 相关的管程状态将会被 更新，也可能退出了，就如当前线程中同执行了 monitorexit 指令一般。
 如果这个本地方法有返回值，那平台相关的代码返回的数据必须通过某种 实现相关的方式转换成本地方法所定义的 Java 类型，并压入到操作数栈 中。










						||
						||
						||
						||
======================================================================================
						||		
	   链接时异常			||		在类、接口或者数组的符号解析阶段，任何在§5.4.3.4 章节中描述的异常 都可能被抛出。
另外，如果调用方法是实例方法，那 invokevirtual 指令将抛出 IncompatibleClassChangeError 异常。


						||
======================================================================================
						||
		运行时异常		||		如果 objectref 为 null，invokevirtual 指令将抛出 NullPointerException 异常。
另外，如果没有找到任何名称和描述符都与要调用的接口方法一致的方法，那 invokevirtual 指令将抛出 AbstractMethodError 异常。
另外，如果被调用方法是 abstract 的，那 invokevirtual 指令将抛出 AbstractMethodError 异常。
另外，如果搜索到的方法是 abstract 的话，那 invokevirtual 指令将抛 出 AbstractMethodError 异常。
另外，如果搜索到的方法是 native 的话，当实现代码实现代码无法绑定到虚 拟机中，那 invokevirtual 指令将抛出 UnsatisfiedLinkError 异常。






						||
======================================================================================
						||
						||objectref 和随后的 n 个参数并不一定与局部变量表的数量一一对应，因为 参数中的 long 和 double 类型参数需要使用 2 个连续的局部变量来存储，因此在参数传递的时候，可能需要比参数个数更多的局部变量。
		注意				||
						||
						||
						||
======================================================================================
 */