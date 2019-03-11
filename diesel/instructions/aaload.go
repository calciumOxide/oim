package instructions

import (
	"../runtime"
	"../../utils"
)

type I_aaload struct {
}

func init()  {
	INSTRUCTION_MAP[0x32] = &I_aaload{}
}

func (s I_aaload)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "aaload exce >>>>>>>>>\n")

/*
	frame := ctx.CurrentFrame
	index := frame.Layers[frame.Depth].Value
	arrayRef := frame.Layers[frame.Depth - 1].Value
	frame.Depth -= 2

	aborigines := ctx.CurrentAborigines
	layer := aborigines.Layers[arrayRef.(uint32)]
	value := layer.Value.([]uint32)[index.(uint32)]

	frame.Layers[frame.Depth] = &runtime.Layer{
		Type: 0,
		Value: value,
	}
	frame.Depth += 1
*/
	frame := ctx.CurrentFrame
	index, _ := frame.PopFrame()
	arrayRef, _ := frame.PopFrame()

	aborigines := ctx.CurrentAborigines
	layer, _ := aborigines.GetAborigines(arrayRef.(uint32))
	value := (*(layer.(*[]uint32)))[index.(uint32)]

	frame.PushFrame(value)

	return nil
}

func (s I_aaload)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.Depth = 2
	f.Layers = append(f.Layers, uint32(0))
	f.Layers = append(f.Layers, uint32(0))
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x32},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		从数组中加载一个 reference 类型数据到操作数栈
======================================================================================
						||		aaload
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
		结构				||		aaload = 50(0x32)
======================================================================================
						||		...，arrayref，index →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||		
						||		
						||		arrayref 必须是一个 reference 类型的数据，它指向一个组件类型为 reference 的数组，index 必须为 int 类型。
		描述				||		指令执行后，arrayref 和 index 同时从操作数栈出栈，index 作为索引定位到数组中的 reference类型值将压入到操作数栈中。
						||		
						||		
						||		
======================================================================================
						||		
						||		
						||		如果 arrayref 为 null，aaload 指令将抛出 NullPointerException 异 常。
	   运行时异常			||		另外，如果 index 不在 arrayref 所代表的数组上下界范围中，aaload 指 令将抛出 ArrayIndexOutOfBoundsException 异常
						||		
						||		
						||		
======================================================================================
 */