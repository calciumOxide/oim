package instructions

import (
	"../../utils"
	"../oil/types"
	"../runtime"
	"../variator"
)

type I_daload struct {
}

func init() {
	INSTRUCTION_MAP[0x31] = &I_daload{}
}

func (s I_daload) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "daload exce >>>>>>>>>\n")

	index, _ := ctx.CurrentFrame.PopFrame()
	array, _ := ctx.CurrentFrame.PopFrame()
	if array == nil {
		except, _ := variator.AllocExcept(variator.NullPointerException)
		ctx.Throw(except)
		return nil
	}
	bytes := array.(*types.Jarray).Reference.([]interface{})
	if len(bytes) <= int(index.(types.Jint)) {
		except, _ := variator.AllocExcept(variator.ArrayIndexOutOfBoundsException)
		ctx.Throw(except)
		return nil
	}
	value := bytes[index.(types.Jint)]
	ctx.CurrentFrame.PushFrame(value)
	return nil
}

func (s I_daload) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []interface{}{1.11, 2.22, 3.33, 4.44},
	})
	f.PushFrame(types.Jint(2))
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code:              []byte{0x0},
		CurrentFrame:      f,
		CurrentAborigines: a,
	}
}

/**
======================================================================================
		操作				||		从数组中加载一个 double 类型数据到操作数栈
======================================================================================
						||		daload
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
		结构				||		daload = 49(0x31)
======================================================================================
						||		...，arrayref，index →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||
						||		arrayref 是一个 reference 类型的数据，它指向一个以 double 为组件类型的数组对象，
		描述				||		index 是一个 int 型的数据。在指令执行 时，arrayref 和 index 都从操作数栈中出栈，
						||		在数组中使用 index 为索引 index 作为索引定位到数组中的 double 类型值将压入到操作数栈中。
						||
======================================================================================
						||
						||		如果 arrayref 为 null，daload 指令将抛出 NullPointerException 异 常。
						||
	   运行时异常			||		另外，如果 index 不在数组的上下界范围之内，daload 指令将抛出 ArrayIndexOutOfBoundsException 异常。
						||
						||
						||
======================================================================================
						||
						||		daload 指令可以用来从数组中读取 byte 或者 boolean 的数据，在 Oracle 的虚拟机实现中，
						||		布尔类型的数组(T_BOOLEAN 类型的数组可参见§2.2 和本章中对 newarray 指令的介绍)被实现为 8 位宽度的数值，
		注意				||		而其他的虚拟机 实现很可能使用其他方式实现 boolean 数组，那其他虚拟机实现的 daload 就必须能正确访问相应实现的数组。
						||
						||
						||
======================================================================================
*/
