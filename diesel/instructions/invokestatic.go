package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
		"../variator"
	"../../loader/clazz"
	"../../loader/clazz/item"
)

type I_invokestatic struct {
}

func init()  {
	INSTRUCTION_MAP[0xb8] = &I_invokestatic{}
}

func (s I_invokestatic)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "invokestatic exce >>>>>>>>>\n")

	methodIndex := (uint16(ctx.Code[ctx.PC]) << 8) | uint16(ctx.Code[ctx.PC + 1])
	ctx.PC += 2

	methodInterface, _ := ctx.Clazz.GetConstant(methodIndex)
	classCp, _ := ctx.Clazz.GetConstant(methodInterface.Info.(*item.MethodRef).ClassIndex)
	classNameCp, _ := ctx.Clazz.GetConstant(classCp.Info.(*item.Class).NameIndex)
	class := clazz.GetClass(classNameCp.Info.(*item.Utf8).Str)

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
	if method.IsPrivate() {
		except, _ := variator.AllocExcept(variator.AbstractMethodError)
		ctx.Throw(except)
		return nil
	}


	if method == nil {
		except, _ := variator.AllocExcept(variator.MethodNotFindException)
		ctx.Throw(except)
		return nil
	}

	if !method.IsStatic() {
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
	if !method.CheckParams(parmsType, args) {
		except, _ := variator.AllocExcept(variator.MethodParamsNoMatchExcept)
		ctx.Throw(except)
		return nil
	}
	ctx.InvokeMethod(method, args)
	return nil
}

func (s I_invokestatic)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(types.Jint(1))
	f.PushFrame(types.Jdouble(9.123456789012345))
	f.PushFrame(&types.Jreference{
		Reference: types.Jobject{},
		ElementType: clazz.CLASS_MAP["com/oxide/A"],
	})
	f.PushFrame(types.Jfloat(9.123456789012345))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x0, 0x0, 12, 0x1, 0x0},
		//Code: []byte{0x0, 0x0, 0x1, 0x1, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		调用实例方法;专门用来处理调用超类方法、私有方法和实例初始化方法方法
======================================================================================
						||		invokestatic
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
		结构				||		invokestatic = 184(0xb8)
======================================================================================
						||		...，[arg1，[arg2 ...]] →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
		描述				||		无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运 行时常量池的索引值，
						||		构建方式为(indexbyte1 << 8)| indexbyte2，
						||		该索引所指向的运行时常量池项应当是一个方法(§5.1)的符号引用，
						||		它包 含了方法的名称和描述符(§4.3.3)，以及包含该方法的接口的符号引用。
此方法应是已被解析过(§5.4.3.3)的，而且不能是实例的初始化方法和类 或接口的初始化方法。
这个方法必须被声明为 static，因此它也不能是 abstract 方法。

在方法被成功解析之后，如果方法所在的类没有被初始化过(§5.5)，那指令 执行时将会触发其初始化过程。

在操作数栈中必须包含连续 n 个参数值，这些参数的数值、数据类型和顺序都 必须遵循实例方法的描述符中的描述。

如果要调用的是同步方法，那与这个类的 Class 对象相关的管程(monitor) 将会进入或者重入，就如当前线程中同执行了 monitorenter 指令一般。

如果要调用的不是本地方法，n 个 args 参数将从操作数栈中出栈。
方法调用 的时候，一个新的栈帧将在 Java 虚拟机栈中被创建出来，连续的 n 个参数将 存放到新栈帧的局部变量表中，
arg1 存为局部变量 1(如果 arg1 是 long 或 double 类型，那将占用局部变量 1 和 2 两个位置)，依此类推。
参数中的浮点类型数据在存入局部变量之前会先进行数值集合转换(§2.8.3)。
新栈帧 创建后就成为当前栈帧，Java 虚拟机的 PC 寄存器被设置成指向调用方法的首 条指令，程序就从这里开始继续执行。

如果要调用的是本地方法，要是这些平台相关的代码尚未绑定(§5.6)到虚 拟机中的话，绑定动作先要完成。
指令执行时，n 个 args 参数将从操作数栈 中出栈并作为参数传递给实现此方法的代码。
参数中的浮点类型数据在传递给 调用方法之前会先进行数值集合转换(§2.8.3)。
参数传递和代码执行都会 以具体虚拟机实现相关的方式进行。当这些平台相关的代码返回时:
 如果这个本地方法是同步方法，那与它所属类的 Class 对象相关的管程 状态将会被更新，也可能退出了，就如当前线程中同执行了 monitorexit 指令一般。
 如果这个本地方法有返回值，那平台相关的代码返回的数据必须通过某种 实现相关的方式转换成本地方法所定义的 Java 类型，并压入到操作数栈 中。







						||
						||
						||
						||
======================================================================================
						||		
	   链接时异常			||		在类、接口或者数组的符号解析阶段，任何在§5.4.3.3 章节中描述的异常 都可能被抛出。
另外，如果调用方法是实例方法，那 invokestatic 指令将抛出 IncompatibleClassChangeError 异常。


						||
======================================================================================
						||
		运行时异常		||		如果invokestatic指令执行时触发了类的初始化过程，那invokestatic 指令有可能所有在§5.5 中描述过的异常。
另外，如果执行的方法是 native 方法的话，当实现代码实现代码无法绑定到 虚拟机中，那 invokestatic 指令将抛出 UnsatisfiedLinkError 异常。







						||
======================================================================================
						||
						||方法调用使用到的 n 个参数的总个数并非与局部变量表的个数是一一对应的， 因为参数中的 long 和 double 类型参数需要使用 2 个连续的局部变量来存储， 因此在参数传递的时候，可能需要比参数个数更多的局部变量。						||
		注意				||
						||
						||
						||
======================================================================================
 */