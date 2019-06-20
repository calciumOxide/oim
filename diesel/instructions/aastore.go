package instructions

import (
	"../../utils"
	"../runtime"
)

type I_aastore struct {
}

func init() {
	INSTRUCTION_MAP[0x53] = &I_aastore{}
}

func (s I_aastore) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "aastore exce >>>>>>>>>\n")

	frame := ctx.CurrentFrame
	value, _ := frame.PopFrame()
	index, _ := frame.PopFrame()
	arrayRef, _ := frame.PopFrame()

	aborigines := ctx.CurrentAborigines
	layer, _ := aborigines.GetAborigines(arrayRef.(uint32))
	(*(layer.(*[]uint32)))[index.(uint32)] = value.(uint32)

	return nil
}

func (s I_aastore) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.Depth = 3
	f.Layers = append(f.Layers, uint32(0))
	f.Layers = append(f.Layers, uint32(0))
	f.Layers = append(f.Layers, uint32(5678))
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code:              []byte{0x32},
		CurrentFrame:      f,
		CurrentAborigines: a,
	}
}

/**
======================================================================================
		操作				||		从操作数栈读取一个 reference 类型数据存入到数组中
======================================================================================
						||		aastore
						||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||
		格式				||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||
======================================================================================
		结构				||		aastore = 83(0x53)
======================================================================================
						||		...，arrayref，index，value →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||
						||
						||		arrayref 必须是一个 reference 类型的数据，它指向一个组件类型为 reference 的数组，index 必须为 int 类型，
						||		value 必须为 reference类型。指令执行后，arrayref、index 和 value 同时从操作数栈出栈，
						||		value 存储到 index 作为索引定位到数组元素中。
						||
						||		在运行时，value 的实际类型必须与 arrayref 所代表的数组的组件类型相 匹配。
						||		具体地说，reference 类型值 value(记作 S)能匹配组件类型为 reference(记作 T)的数组的前提是
						||
						||		 如果 S 是类类型(Class Type)，那么: 如果 S 是类类型(Class Type)，那么:
						||			 如果T也是类类型，那S必须与T是同一个类类型，或者S是T所代表的类型的子类。
		描述				||			 如果 T 是接口类型，那 S 必须实现了 T 的接口。
						||
						||		 如果 S 是接口类型(Interface Type)，那么:
						||			 如果 T 是类类型，那么 T 只能是 Object。
						||			 如果T是接口类型，那么T与S应当是相同的接口，或者T是S的父接口。
						||
						||		 如果 S 是数组类型(Array Type)，假设为 SC[]的形式，这个数组的组件类型为 SC，那么:
						||			 如果 T 是类类型，那么 T 只能是 Object。
						||			 如果 T 是数组类型，假设为 TC[]的形式，这个数组的组件类型为 TC，那么下面两条规则之一必须成立:
						||				 TC 和 SC 是同一个原始类型。
						||				 TC 和 SC 都是 reference 类型，并且 SC 能与 TC 类型相匹配(以此处描述的规则来判断是否互相匹配)。
						||
						||		 如果 T 是接口类型，那 T 必须是数组类型所实现的接口之一(JLS3§4.10.3)。
						||
						||
======================================================================================
						||
						||
						||		如果 arrayref 为 null，aastore 指令将抛出 NullPointerException 异常
	   运行时异常			||		另外，如果 index 不在 arrayref 所代表的数组上下界范围中，aastore 指 令将抛出 ArrayIndexOutOfBoundsException 异常。
						||		另外，如果 arrayref 不为 null，并且 value 的实际类型与数组组件类型 不能互相匹配(JLS3 §5.2)，aastore 指令将抛出 ArrayStoreException 异常。
						||
						||
======================================================================================
*/
