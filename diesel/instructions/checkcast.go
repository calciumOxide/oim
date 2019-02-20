package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
			"../variator"
	"reflect"
)

type I_checkcast struct {
}

func init()  {
	INSTRUCTION_MAP[0xc0] = &I_checkcast{}
}

func (s I_checkcast)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "checkcast exce >>>>>>>>>\n")

	index := uint16(ctx.Code[ctx.PC]) << 8 | uint16(ctx.Code[ctx.PC + 1])
	ctx.PC += 2
	ref, _ := ctx.CurrentFrame.PeekFrame()

	if ref == nil || reflect.TypeOf(ref) != reflect.TypeOf(&types.Jreference{}) || ref.(*types.Jreference).Reference == nil {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	//ctx.Clazz.(*clazz.ClassFile).GetConstant(index)
	if ref.(*types.Jreference).Reference.(*types.Jobject).ClassTypeIndex != index {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	return nil
}

func (s I_checkcast)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jreference{
		Reference: &types.Jobject{
			ClassTypeIndex: 0x0F0F,
		},
	})
	f.PushFrame(33)

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x0, 0x0F, 0x0F},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		检查对象是否符合给定的类型
======================================================================================
						||		checkcast
						||------------------------------------------------------------
						||		indexbyte1
						||------------------------------------------------------------
						||		Indexbyte2
		格式				||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		checkcast = 192(0xc0)
======================================================================================
						||		...，objectref →
	   操作数栈			||------------------------------------------------------------
						||		...，objectref
======================================================================================
						||		
						||		objectref 必须为 reference 类型的数据，indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运行时常量池的索引值，
		描述				||		构建方式为 (indexbyte1 << 8)| indexbyte2，该索引所指向的运行时常量池项应 当是一个类、接口或者数组类型的符号引用。
						||
						||		如果 objectref 为 null 的话，那操作数栈不会有任何变化。
						||
						||		否则，参数指定的类、接口或者数组类型会被虚拟机解析(§5.4.3.1)。
						||		如 果 objectref 可以转换为这个类、接口或者数组类型，那操作数栈就保持不 变，否则 checkcast 指令将抛出一个 ClassCastException 异常。
						||
						||		以下规则可以用来确定一个非空的 objectref 是否可以转换为指定的已解析 类型:
						||		假设 S 是 objectref 所指向的对象的类型，T 是进行比较的已解析的 类、接口或者数组类型，checkcast 指令根据这些规则来判断转换是否成立:
						||			 如果 S 是类类型(Class Type)，那么:
						||				 如果T也是类类型，那S必须与T是同一个类类型，或者S是T所 代表的类型的子类。
						||				 如果 T 是接口类型，那 S 必须实现了 T 的接口。
						||			 如果 S 是接口类型(Interface Type)，那么:
						||				 如果 T 是类类型，那么 T 只能是 Object。
						||				 如果T是接口类型，那么T与S应当是相同的接口，或者T是S的 父接口。
						||			 如果 S 是数组类型(Array Type)，假设为 SC[]的形式，这个数组的组 件类型为 SC，那么:
						||				 如果 T 是类类型，那么 T 只能是 Object。
						||				 如果 T 是数组类型，假设为 TC[]的形式，这个数组的组件类型为 TC，
						||					 TC 和 SC 是同一个原始类型。
						||					 TC 和 SC 都是 reference 类型，并且 SC 能与 TC 类型相匹配(以此处描述的规则来判断是否互相匹配)。
						||		如果T是接口类型，那T必须是数组所实现的接口之一(JLS3 §4.10.3)。
						||
						||
======================================================================================
						||		
						||
						||
	   链接时异常			||		在类、接口或者数组的符号解析阶段，任何在§5.4.3.1 章节中描述的异常 都可能被抛出。
						||		
						||		
						||		
======================================================================================
						||
						||
						||
	   运行时异常			||		如果 objectref 不能转换成参数指定的类、接口或者数组类型，checkcast 指令将抛出 ClassCastException 异常
						||
						||
						||
======================================================================================
						||
						||		checkcast 指令与 instanceof 指令非常类似，它们之间的区别是如何处理 null 值的情况、测试类型转换的结果反馈方式(checkcast 是抛异常，
						||		而 instanceof 是返回一个比较结果)以及指令执行后对操作数栈的影响。
		注意				||
						||
						||
						||
======================================================================================
 */