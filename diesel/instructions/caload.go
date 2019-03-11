package instructions

import (
	"../runtime"
	"../../utils"
	"../variator"
	"../../types"
)

type I_caload struct {
}

func init()  {
	INSTRUCTION_MAP[0x34] = &I_caload{}
}

func (s I_caload)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "caload exce >>>>>>>>>\n")

	index, _ := ctx.CurrentFrame.PopFrame()
	array, _ := ctx.CurrentFrame.PopFrame()
	if array == nil {
		except, _ := variator.AllocExcept(variator.NullPointerException)
		ctx.Throw(except)
		return nil
	}
	bytes := array.(*types.Jarray).Reference.([]types.Jchar)
	if len(bytes) <= int(index.(types.Jint)) {
		except, _ := variator.AllocExcept(variator.ArrayIndexOutOfBoundsException)
		ctx.Throw(except)
		return nil
	}
	value := bytes[index.(types.Jint)]
	ctx.CurrentFrame.PushFrame(types.Jchar(value))
	return nil
}

func (s I_caload)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jchar{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(2))

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
		操作				||		从数组中加载一个 char 类型数据到操作数栈
======================================================================================
						||		caload
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
		结构				||		caload = 52(0x34)
======================================================================================
						||		...，arrayref，index →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||		
						||		arrayref 必须是一个 reference 类型的数据，它指向一个组件类型为 char 的数组，
		描述				||		index 和 value 都必须为 int 类型。指令执行后，arrayref、index 同时从操作数栈出栈，
						||		index 作为索引定位到数组中的 char 类型值先被零位扩展 (Zero-Extended)为一个 int 类型数据 value，然后再将 value 压入到 操作数栈中。
						||		
======================================================================================
						||		
						||		如果 arrayref 为 null，caload 指令将抛出 NullPointerException 异 常。
						||
	   运行时异常			||		另外，如果 index 不在数组的上下界范围之内，caload 指令将抛出 ArrayIndexOutOfBoundsException 异常。
						||		
						||		
						||		
======================================================================================
						||
		注意				||		
						||
						||
						||
======================================================================================
 */