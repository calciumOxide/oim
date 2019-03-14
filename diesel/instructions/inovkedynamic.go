package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	"../../loader/clazz/item"
)

type I_inovkedynamic struct {
}

func init()  {
	INSTRUCTION_MAP[0xba] = &I_inovkedynamic{}
}

func (s I_inovkedynamic)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "inovkedynamic exce >>>>>>>>>\n")

	index := utils.BigEndian2Little4U2(ctx.Code[ctx.PC : ctx.PC+2])
	ctx.PC += 2

	mrCp, _ := ctx.Clazz.GetConstant(index)
	kind := mrCp.Info.(*item.MethodHandle).ReferenceKind

	switch kind {
		case item.GET_FIELD:
		break
		case item.GET_STATIC:
		break
		case item.PUT_FIELD:
		break
		case item.PUT_STATIC:
		break
		case item.INVOKE_VIRTUAL:
		break
		case item.INVOKE_STATIC:
		break
		case item.INVOKE_SPECIAL:
		break
		case item.NEW_INVOKE_SPECIAL:
		break
		case item.INVOKE_INTERFACE:
		break
	}

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jint(value.(types.Jint) * -1))
	return nil
}

func (s I_inovkedynamic)Test(octx *runtime.Context) *runtime.Context {
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

 /**

 http://note.youdao.com/noteshare?id=8af39e568b5e8870cf1a124212ff9539


 理解 invokedynamic
96  TiouLims
 0.7 2017.10.12 17:17* 字数 2715 阅读 2268评论 2喜欢 9
inDy（invokedynamic）是 java 7 引入的一条新的虚拟机指令，这是自 1.0 以来第一次引入新的虚拟机指令。到了 java 8 这条指令才第一次在 java 应用，用在 lambda 表达式中。 indy 与其他 invoke 指令不同的是它允许由应用级的代码来决定方法解析。所谓应用级的代码其实是一个方法，在这里这个方法被称为引导方法（Bootstrap Method），简称 BSM。BSM 返回一个 CallSite（调用点） 对象，这个对象就和 inDy 链接在一起了。以后再执行这条 inDy 指令都不会创建新的 CallSite 对象。CallSite 就是一个 MethodHandle（方法句柄）的 holder。方法句柄指向一个调用点真正执行的方法。

理解 MethodHandle（方法句柄）的一种方式就是将其视为以安全、现代的方式来实现反射的核心功能。

一个 java 方法的实体有四个构成：

方法名
签名--参数列表和返回值
定义方法的类
方法体（代码）
同一个类中，方法名相同，签名不同，JVM 会视为不同的方法，不过在 Java 中只支持签名的参数列表部分，也就是重载多态。一次方法调用，除了要方法的实体外，还要调用者（caller）和接收者（receiver），调用者也就是方法调用语句所在的类。接收者是一个对象，每个方法调用都要一个接收者，它可以是隐藏的（this），也可以是类方法，比如： String.valueOf，类也是 Class 的一个实例。

MethodType 表示方法签名。

用 MethodHandle 实现的方法调用的示例如下，可以看到方法的四个构成：

Object rcvr = "a";
try {
    MethodType mt = MethodType.methodType(int.class); // 方法签名
    MethodHandles.Lookup l = MethodHandles.lookup(); // 调用者，也就是当前类。调用者决定有没有权限能访问到方法
    MethodHandle mh = l.findVirtual(rcvr.getClass(), "hashCode", mt); //分别是定义方法的类，方法名，签名

    int ret;
    try {
        ret = (int)mh.invoke(rcvr); // 代码，第一个参数就是接收者
        System.out.println(ret);
    } catch (Throwable t) {
        t.printStackTrace();
    }
} catch (IllegalArgumentException | NoSuchMethodException | SecurityException e) {
    e.printStackTrace();
} catch (IllegalAccessException x) {
    x.printStackTrace();
}
详细可参考：

Invokedynamic - Java’s Secret Weapon
和译文 Invokedynamic：Java的秘密武器 - 知乎专栏
java8 lambda 表达式
lambda 表达式 是怎么使用 inDy 呢？以一段简单的代码为例

public class LambdaTest {
    public static void main(String[] args) {
        Runnable r = () -> System.out.println(Arrays.toString(args));
        r.run();
    }
}
用 javap -v -p LambdaTest 查看字节码，可以发现寥寥几行 java 代码生成的字节码却不少，单单常量池常量就有 66 个之多。输出见 LambdaTest.class。

可以发现多出了一个新方法，方法体就是 lambda 体（lambda body），转换为源码如下：

private static void lambda$main$0(java.lang.String[] args){
    System.out.println(Arrays.toString(args));
}
主要看一下 main 方法，并没有直接调用上面的方法，而是出现一条 inDy 指令：

public static void main(java.lang.String[]);
    descriptor: ([Ljava/lang/String;)V
    flags: ACC_PUBLIC, ACC_STATIC
    Code:
      stack=1, locals=2, args_size=1
         0: aload_0
         1: invokedynamic #2,  0              // InvokeDynamic #0:run:([Ljava/lang/String;)Ljava/lang/Runnable;
         6: astore_1
         7: aload_1
         8: invokeinterface #3,  1            // InterfaceMethod java/lang/Runnable.run:()V
        13: return
可以看到 inDy 指向一个类型为 CONSTANT_InvokeDynamic_info 的常量项 #2，另外 0 是预留参数，暂时没有作用。

#2 = InvokeDynamic      #0:#30         // #0:run:([Ljava/lang/String;)Ljava/lang/Runnable;
#0 表示在 Bootstrap methods 表中的索引：

BootstrapMethods:
  0: #27 invokestatic java/lang/invoke/LambdaMetafactory.metafactory:(Ljava/lang/invoke/MethodHandles$Lookup;Ljava/lang/String;Ljava/lang/invoke/MethodType;Ljava/lang/invoke/MethodType;Ljava/lang/invoke/MethodHandle;Ljava/lang/invoke/MethodType;)Ljava/lang/invoke/CallSite;
    Method arguments:
      #28 ()V
      #29 invokestatic com/company/LambdaTest.lambda$main$0:([Ljava/lang/String;)V
      #28 ()V
#30 则是一个 CONSTANT_NameAndType_info，表示方法名和方法类型（返回值和参数列表），这个会作为参数传递给 BSM。

#30 = NameAndType        #43:#44        // run:([Ljava/lang/String;)Ljava/lang/Runnable;
再看回表中的第 0 项，#27 是一个 CONSTANT_MethodHandle_info，实际上是个 MethodHandle（方法句柄）对象，这个句柄指向的就是 BSM 方法。在这里就是:

java.lang.invoke.LambdaMetafactory.metafactory(MethodHandles.Lookup,String,MethodType,MethodType,MethodHandle,MethodType)
BSM 前三个参数是固定的，后面还可以附加任意数量的参数，但是参数的类型是有限制的，参数类型只能是

String
Class
int
long
float
double
MethodHandle
MethodType
LambdaMetafactory.metafactory 带多三个参数，这些的参数的值由 Bootstrap methods 表 提供：

Method arguments:
  #25 ()V
  #26 invokestatic com/company/LambdaTest.lambda$main$0:()V
  #25 ()V
inDy 所需要的数据大概就是这些，可参考 Java8学习笔记（2） -- InvokeDynamic指令 - CSDN博客

inDy 运行时
每一个 inDy 指令都称为 Dynamic Call Site(动态调用点)，根据 jvm 规范所说的，inDy 可以分为两步，这两步部分代码代码是在 java 层的，给 metafactory 方法设断点可以看到一些行为。

第一步 inDy 需要一个 CallSite（调用点对象），CallSite 是由 BSM 返回的，所以这一步就是调用 BSM 方法。代码可参考：java.lang.invoke.CallSite#makeSite

调用 BSM 方法可以看作 invokevirtual 指令执行一个 invoke 方法，方法签名如下：

invoke:(MethodHandle,Lookup,String,MethodType,/*其他附加静态参数*\/)CallSite
前四个参数是固定的，被依次压入操作栈里

MethodHandle，实际上这个方法句柄就是指向 BSM
Lookup, 也就是调用者，是 Indy 指令所在的类的上下文，可以通过 Lookup#lookupClass()获取这个类
name ，lambda 所实现的方法名，也就是"run"
invokedType，调用点的方法签名，这里是 methodType(Runnable.class,String[].class)
接下来就是附加参数，这些参数是灵活的，由Bootstrap methods 表提供，这里分别是：

samMethodType，其实就是 Runnable.run 的描述符: methodType(void.class)。sam 就 single public abstract method 的缩写
implMethod: 编译器给生成的 desugar 方法，是一个 MethodHandle：caller.findStatic(LambdaTest.class,"lambda$main$0",methodType(void.class))
instantiatedMethodType: Runnable.run 运行时的描述符，如果方法泛型的，那这个类型可能不一样。这里是 methodType(void.class)
上面说的固定其实应该是指 inDy 传递的实参类型是固定的，BSM 形参声明可以是随意，保证 BSM 能被调用就行，比如说 Lookup 声明为 Object 不影响调用。

接下来就是执行 LambdaMetafactory.metafactory 方法了，它会创建一个匿名类，这个类是通过 ASM 编织字节码在内存中生成的，然后直接通过 unsafe 直接加载而不会写到文件里。不过可以通过下面的虚拟机参数让它运行的时候输出到文件

-Djdk.internal.lambda.dumpProxyClasses=<path>
这个类是根据 lambda 的特点生成的，输出后可以看到，在这个例子中是这样的：

import java.lang.invoke.LambdaForm.Hidden;

// $FF: synthetic class
final class LambdaTest$$Lambda$1 implements Runnable {
private final String[] arg$1;

private LambdaTest$$Lambda$1(String[] var1) {
this.arg$1 = var1;
}

private static Runnable get$Lambda(String[] var0) {
return new LambdaTest$$Lambda$1(var0);
}

@Hidden
public void run() {
LambdaTest.lambda$main$0(this.arg$1);
}
}
然后就是创建一个 CallSite，绑定一个 MethodHandle，指向的方法其实就是生成的类中的静态方法 LambdaTest$$Lambda$1.get$Lambda(String[])Runnable。然后把调用点对象返回，到这里 BSM 方法执行完毕。

更详细的可参考：

浅谈Lambda Expression - 简书
[Java] 关于OpenJDK对Java 8 lambda表达式的运行时实现的查看方式 - 知乎专栏
第二步，就是执行这个方法句柄了，这个过程就像 invokevirtual 指令执行 MethodHandle#invokeExact 一样，

加上 inDy 上面那一条 aload_0 指令，这是操作数栈有两个分别是：

args[]，lambda 里面调用了 main 方法的参数
调用点对象（CallSite），实际上是方法句柄。如果是 CostantCallSite 的时候，inDy 会直接跟他的方法句柄链接。见代码：MethodHandleNatives.java#L255
传入 args，执行方法，返回一个 Runnable 对象，压入栈顶。到这里 inDy 就执行完毕。

接下来的指令就很好理解，astore_1 把栈顶的 Runnable 对象放到局部变量表的槽位1，也是变量 r。剩下的就是再拿出来调用 run 方法。

Groovy
接下来看一下 groovy 是如何使用 inDy 指令的。先复习一遍 groovy 的方法派发。

每当 Groovy 调用一个方法时，它不会直接调用它，而是要求一个中间层来代替它。 中间层通过钩子方法允许我们更改方法调用的行为。这个中间层就是 MOP（meta object proctol），MOP 主要承载的类就是 MetaClass 。一个简化版的 MOP 主要有这些方法：

invokeMethod(String methodName, Object args)
methodMissing(String name, Object arguments)
getProperty(String propertyName)
setProperty(String propertyName, Object newValue)
propertyMissing(String name)
可以大致认为在 Groovy 中的每个方法和属性访问调用都会转化上面的方法调用。而这些方法可以在运行时通过重写修改它的默认行为，MOP 作为方法派发的中心枢纽为 Groovy 提供了非常灵活的动态编程的能力。

现在来看一下一段简短的 groovy 代码，

class Test{
int a = 0;
static void main(args){
Test wtf = new Test()
wtf.a
wtf.doSomething()
}
}
通过 groovyc -indy Test.groovy 把它编译成字节码。 indy 选项的意思就是启用 invokedynamic 支持。

看一下编译后的 main 方法。

public static void main(java.lang.String...);
descriptor: ([Ljava/lang/String;)V
flags: ACC_PUBLIC, ACC_STATIC, ACC_VARARGS
Code:
stack=1, locals=2, args_size=1
0: ldc           #2                  // class Test
2: invokedynamic #44,  0             // InvokeDynamic #0:init:(Ljava/lang/Class;)Ljava/lang/Object;
7: invokedynamic #50,  0             // InvokeDynamic #1:cast:(Ljava/lang/Object;)LTest;
12: astore_1
13: aload_1
14: pop
15: aload_1
16: invokedynamic #56,  0             // InvokeDynamic #2:getProperty:(LTest;)Ljava/lang/Object;
21: pop
22: aload_1
23: invokedynamic #61,  0             // InvokeDynamic #3:invoke:(LTest;)Ljava/lang/Object;
28: pop
29: return
可以看到一共有 4 条 inDy 指令，包括构造函数，访问成员变量，和不存在的方法调用都是 通过 invokedynamic 实现的。

再看一下引导方法表

BootstrapMethods:
0: #38 invokestatic org/codehaus/groovy/vmplugin/v7/IndyInterface.bootstrap:(Ljava/lang/invoke/MethodHandles$Lookup;Ljava/lang/String;Ljava/lang/invoke/MethodType;Ljava/lang/String;I)Ljava/lang/invoke/CallSite;
Method arguments:
#39 <init>
#40 0
1: #38 invokestatic org/codehaus/groovy/vmplugin/v7/IndyInterface.bootstrap:(Ljava/lang/invoke/MethodHandles$Lookup;Ljava/lang/String;Ljava/lang/invoke/MethodType;Ljava/lang/String;I)Ljava/lang/invoke/CallSite;
Method arguments:
#46 ()
#40 0
2: #38 invokestatic org/codehaus/groovy/vmplugin/v7/IndyInterface.bootstrap:(Ljava/lang/invoke/MethodHandles$Lookup;Ljava/lang/String;Ljava/lang/invoke/MethodType;Ljava/lang/String;I)Ljava/lang/invoke/CallSite;
Method arguments:
#51 a
#52 4
3: #38 invokestatic org/codehaus/groovy/vmplugin/v7/IndyInterface.bootstrap:(Ljava/lang/invoke/MethodHandles$Lookup;Ljava/lang/String;Ljava/lang/invoke/MethodType;Ljava/lang/String;I)Ljava/lang/invoke/CallSite;
Method arguments:
#58 doSomething
#40 0
可以发现所有 inDy 指令的引导方法都是 IndyInterface.bootstrap

以方法调用的 inDy 指令为例，它的方法名称是 "invoke"，方法签名是 methodType(Object.class,Test.class)，BSM 方法还附带两个参数分别是实际的方法名："doSomething" 和一个标志：0

BSM 方法最终调用的是 realBootstrap 方法：

private static CallSite realBootstrap(Lookup caller, String name, int callID, MethodType type, boolean safe, boolean thisCall, boolean spreadCall) {
MutableCallSite mc = new MutableCallSite(type); //这里是 MutableCallSite，lambda 表达式用的是 ConstantCallSite
MethodHandle mh = makeFallBack(mc,caller.lookupClass(),name,callID,type,safe,thisCall,spreadCall);
mc.setTarget(mh);
return mc;
}
主要的代码是调用 makeFallBack 来获取一个临时的 MethodHandle。因为在第一步 groovy 无法确定接收者（receiver），也是就是 invoke 方法的第一个实参（Test 实例），必须要在第二步确定 CallSite 后才会传递过来。所以方法解析要放在第二步。

protected static MethodHandle makeFallBack(MutableCallSite mc, Class<?> sender, String name, int callID, MethodType type, boolean safeNavigation, boolean thisCall, boolean spreadCall) {
MethodHandle mh = MethodHandles.insertArguments(SELECT_METHOD, 0, mc, sender, name, callID, safeNavigation, thisCall, spreadCall, /*dummy receiver:*\/ 1); //MethodHandle(Object.class,Object[].class)
mh =    mh.asCollector(Object[].class, type.parameterCount()).
asType(type);
return mh;
}
这个 fallback 方法其实就是 selectMethod。insertArguments 在这里主要做了一个柯里化的操作，因为selectMethod 的方法签名是

methodType(Object.class, MutableCallSite.class, Class.class, String.class, int.class, Boolean.class, Boolean.class, Boolean.class, Object.class, Object[].class)
而 inDy 要求的方法签名却是

methodType(Object.class,Test.class)。
所以得经过 insertArguments 的变换，把确定的值填充进去，用最后的数组参数来接收 inDy 传递的参数。这样这个方法就能够被 inDy 调用了。第一步创建 CallSite 到这里就结束。

第二步，就是 selectMethod 方法的调用，这时候 groovy 已经知道方法的接收者 arguments[0]，

public static Object selectMethod(MutableCallSite callSite, Class sender, String methodName, int callID, Boolean safeNavigation, Boolean thisCall, Boolean spreadCall, Object dummyReceiver, Object[] arguments) throws Throwable {
Selector selector = Selector.getSelector(callSite, sender, methodName, callID, safeNavigation, thisCall, spreadCall, arguments);
selector.setCallSiteTarget();

MethodHandle call = selector.handle.asSpreader(Object[].class, arguments.length);
call = call.asType(MethodType.methodType(Object.class,Object[].class));
return call.invokeExact(arguments);
}
首先创建一个方法解析器，在这里是 MethodSelector。接着调用 setCallSiteTarget()，这个方法就是用来解析实际的方法。具体的过程还是很复杂的，所以也没法说清楚，大体来说就是确定接收者的 MetaClass，决定这个方法是实际的方法，还是交给 MetaClass 的钩子方法，然后就是创建这个方法的 MethodHandle，然后把这个 MethodHandle 的签名转化为要求的签名。这时 selecor.handle 就是最终调用的方法句柄了。接下来就是最终的方法调用了，到这里 inDy 指令就执行完毕了。

还有一个方法值得留意：

public void doCallSiteTargetSet() {
if (!cache) {
if (LOG_ENABLED) LOG.info("call site stays uncached");
} else {
callSite.setTarget(handle);
if (LOG_ENABLED) LOG.info("call site target set, preparing outside invocation");
}
}
这也是为什么用 MutableCallSite 的原因，如果编译器认为这个方法是可以缓存，那么就会把这个 CallSite 绑定到实际的 MethodHandle，后续的调用就不用再重新解析了。

最后
没有相关经验，inDy 还是很不好理解的，学习了 java 8 和 groovy 对 inDy 的应用才有一点大致的认识，文中如果有什么错误，还请帮忙指出。

原文链接：https://dourok.info/2017/10/08/understanding-invokedynamic/




*/