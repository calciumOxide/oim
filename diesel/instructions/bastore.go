package instructions

import (
	"../runtime"
	"../../utils"
	"../variator"
	"../../types"
)

type I_bastore struct {
}

func init()  {
	INSTRUCTION_MAP[0x54] = &I_bastore{}
}

func (s I_bastore)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "bastore exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()
	index, _ := ctx.CurrentFrame.PopFrame()
	array, _ := ctx.CurrentFrame.PopFrame()
	if array == nil {
		except, _ := variator.AllocExcept(variator.NullPointerException)
		ctx.Throw(except)
		return nil
	}
	bytes := array.(*types.Jarray).Reference.([]types.Jbyte)
	if len(bytes) <= int(index.(types.Jint)) {
		except, _ := variator.AllocExcept(variator.ArrayIndexOutOfBoundsException)
		ctx.Throw(except)
		return nil
	}
	bytes[index.(types.Jint)] = value.(types.Jbyte)
	return nil
}

func (s I_bastore)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(2))
	f.PushFrame(types.Jbyte(33))

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
		操作				||		从操作数栈读取一个 byte 或 boolean 类型数据存入到数组中
======================================================================================
						||		bastore
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
		结构				||		bastore = 84(0x54)
======================================================================================
						||		...，arrayref，index，value →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
						||		arrayref 必须是一个 reference 类型的数据，它指向一个组件类型为 byte 或 boolean 的数组，
		描述				||		index 和 value 都必须为 int 类型。指令执行后，arrayref、index 和 value 同时从操作数栈出栈，
						||		value 将被转换为 byte 类型，然后存储到 index 作为索引定位到数组元素中。
						||		
======================================================================================
						||		
						||		如果 arrayref 为 null，bastore 指令将抛出 NullPointerException 异 常。
						||
	   运行时异常			||		另外，如果 index 不在数组的上下界范围之内，bastore 指令将抛出 ArrayIndexOutOfBoundsException 异常。
						||		
						||		
						||		
======================================================================================
						||
						||		bastore 指令可以用来从数组中读取 byte 或者 boolean 的数据，在 Oracle 的虚拟机实现中，
						||		布尔类型的数组(T_BOOLEAN 类型的数组可参见§2.2 和本章中对 newarray 指令的介绍)被实现为 8 位宽度的数值，
		注意				||		而其他的虚拟机 实现很可能使用其他方式实现 boolean 数组，那其他虚拟机实现的 bastore 就必须能正确访问相应实现的数组。
						||
						||
						||
======================================================================================
 */