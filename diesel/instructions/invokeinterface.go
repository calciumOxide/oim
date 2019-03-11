package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	"../../loader/clazz"
	"../../loader/clazz/item"
)

type I_invokeinterface struct {
}

func init()  {
	INSTRUCTION_MAP[0x74] = &I_invokeinterface{}
}

func (s I_invokeinterface)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "invokeinterface exce >>>>>>>>>\n")

	methodIndex := (uint16(ctx.Code[ctx.PC]) << 8) | uint16(ctx.Code[ctx.PC + 1])
	object, _ := ctx.CurrentFrame.PopFrame()
	count := int(ctx.Code[ctx.PC + 2])
	ctx.PC += 4

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
	methodInterface, _ := ctx.Clazz.GetConstant(methodIndex)
	objectClass := object.(*types.Jreference).ElementType.(*clazz.ClassFile)
	andType, _ := ctx.Clazz.GetConstant(methodInterface.Info.(*item.MethodRef).NameAndTypeIndex)
	nameCp, _ := ctx.Clazz.GetConstant(andType.Info.(*item.NameAndType).NameIndex)
	descCp, _ := ctx.Clazz.GetConstant(andType.Info.(*item.NameAndType).DescriptorIndex)
	method := objectClass.GetMethod(nameCp.Info.(*item.Utf8).Str, descCp.Info.(*item.Utf8).Str)
	if method == nil {
		except, _ := variator.AllocExcept(variator.MethodNotFindException)
		ctx.Throw(except)
		return nil
	}

	if method.IsAbstarct() {
		except, _ := variator.AllocExcept(variator.AbstractMethodError)
		ctx.Throw(except)
		return nil
	}

	var args []interface{}
	for i := 0; i < count; i++ {
		top, _ := ctx.CurrentFrame.PopFrame()
		args = append(args, top)
	}
	ctx.InvokeMethod(method, args)
	return nil
}

func (s I_invokeinterface)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jfloat(9.123456789012345))
	f.PushFrame(&types.Jreference{
		Reference: types.Jobject{},
		ElementType: clazz.CLASS_MAP["com/oxide/A.class"],
	})
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x0, 0x0, 0xC, 0x1, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		调用接口方法
======================================================================================
						||		invokeinterface
						||------------------------------------------------------------
						||		indexbyte1
						||------------------------------------------------------------
						||		indexbyte2
		格式				||------------------------------------------------------------
						||		count
						||------------------------------------------------------------
						||		0
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		invokeinterface = 185(0xb9)
======================================================================================
						||		...，objectref，[arg1，[arg2 ...]] →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
		描述				||		无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运 行时常量池的索引值，
						||		构建方式为(indexbyte1 << 8)| indexbyte2，
						||		该索引所指向的运行时常量池项应当是一个接口方法(§5.1)的符号引用，
						||		它包含了方法的名称和描述符，以及包含该方法的接口的符号引用。
						||		这个方法 的符号引用是已被解析过的(§5.4.3.2)，而且这个方法不能是实例初始化 方法(§2.9)和类或接口的初始化方法(§2.9)。
						||		操作数 count 是一个无符号 byte 型数据，而且不能为零。
						||		objectref 必须 是一个 reference 类型的数据。
						||		在操作数栈中，objectref 之后还跟随着 连续 n 个参数值，这些参数的数值、数据类型和顺序都必须遵循接口方法的描 述符中的描述。
						||		invokeinterface 指令的第四个参数规定永远为 byte 类型 的 0。
						||		假设 C 是 objectref 所对应的类，虚拟机将按下面规则查找实际执行的方法:
						||			 如果 C 中包含了名称和描述符都和要调用的接口方法一致的方法，那这个方法就会被调用，查找过程终止。
						||			 另外，如果 C 有父类，查找过程将按顺序递归搜索 C 的直接父类，如果超 类中能搜索到名称和描述符都和要调用的接口方法一致的方法，那这个方 法就会被调用。
						||			 另外，如果抛出 AbstractMethodError 异常。
						||		如果要调用的是同步方法，那与 objectref 相关的管程(monitor)将会进 入或者重入，就如当前线程中同执行了 monitorenter 指令一般。
						||		如果要调用的不是本地方法，n 个 args 参数和 objectref 将从操作数栈中 出栈。
						||		方法调用的时候，一个新的栈帧将在 Java 虚拟机栈中被创建出来， objectref 和连续的 n 个参数将存放到新栈帧的局部变量表中，
						||		，objectref 存为局部变量 0，arg1 存为局部变量 1(如果 arg1 是 long 或 double 类型， 那将占用局部变量 1 和 2 两个位置)，依此类推。
						||		参数中的浮点类型数据在存 入局部变量之前会先进行数值集合转换(§2.8.3)。
						||		新栈帧创建后就成为当 前栈帧，Java 虚拟机的 PC 寄存器被设置成指向调用方法的首条指令，程序就 从这里开始继续执行。
						||
						||		如果要调用的是本地方法，要是这些平台相关的代码尚未绑定(§5.6)到虚 拟机中的话，绑定动作先要完成。
						||		指令执行时，n 个 args 参数和 objectref 将从操作数栈中出栈并作为参数传递给实现此方法的代码。
						||		参数中的浮点类型 数据在传递给调用方法之前会先进行数值集合转换(§2.8.3)。
						||		参数传递和 代码执行都会以具体虚拟机实现相关的方式进行。
						||		当这些平台相关的代码返回 时:
						||			 如果这个本地方法是同步方法，那与 objectref 相关的管程状态将会被更新，也可能退出了，就如当前线程中同执行了 monitorexit 指令一般。
						||			 如果这个本地方法有返回值，那平台相关的代码返回的数据必须通过某种 实现相关的方式转换成本地方法所定义的 Java 类型，并压入到操作数栈中。
						||
						||
						||
						||
======================================================================================
						||		
	   链接时异常			||		在类、接口或者数组的符号解析阶段，任何在§5.4.3.4 章节中描述的异常 都可能被抛出。
						||
======================================================================================
						||
		运行时异常		||		如果 objectref 为 null，invokeinterface 指令将抛出 NullPointerException 异常。
						||		另外，如果 objectref 所对应的对象并未实现接口方法中所需的接口，那 invokeinterface 指令将抛出 IncompatibleClassChangeError 异常。
						||		另外，如果没有找到任何名称和描述符都与要调用的接口方法一致的方法，那 invokeinterface 指令将抛出 AbstractMethodError 异常。
						||		另外，如果搜索到的方法不是 public 的话，那 invokeinterface 指令将 抛出 IllegalAccessError 异常。
						||		另外，如果搜索到的方法是 abstract 的话，那 invokeinterface 指令将 抛出 AbstractMethodError 异常。
						||		另外，如果搜索到的方法是 native 的话，当实现代码实现代码无法绑定到虚 拟机中，那 invokeinterface 指令将抛出 UnsatisfiedLinkError 异常。
						||
						||		invokeinterface 指令的 count 操作数用于确定参数的数量，long 和 double 类型的参数占用 2 个数量单位，而其他类型的参数占用 1 个数量单位。
						||		其实这些信息完全可以从方法的描述符中获取到，有这个参数完全是历史原 因。
						||
						||		invokeinterface 指令的第四个操作数是为了给 Oracle 实现的虚拟机的 额外操作数而预留的空间，
						||		invokeinterface 指令会在运行时被替换为特殊 的其他指令，这必须要保持向后兼容性。
						||
						||		objectref 和随后的 n 个参数并不一定与局部变量表的数量一一对应，
						||		因为 参数中的 long 和 double 类型参数需要使用 2 个连续的局部变量来存储，
						||		因 此在参数传递的时候，可能需要比参数个数更多的局部变量。
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