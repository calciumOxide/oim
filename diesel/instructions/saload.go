package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
			)

type I_swap struct {
}

func init()  {
	INSTRUCTION_MAP[0x5f] = &I_swap{}
}

func (s I_swap)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "swap exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()

	//if reflect.TypeOf(value1) != reflect.TypeOf(types.Jlong(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jlong(0)) {
	//	except, _ := variator.AllocExcept(variator.ClassCastException)
	//	ctx.Throw(except)
	//	return nil
	//}

	ctx.CurrentFrame.PushFrame(value2)
	ctx.CurrentFrame.PushFrame(value1)
	return nil
}

func (s I_swap)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jlong(9))
	f.PushFrame(types.Jlong(9))
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
		操作				||		交换操作数栈顶的两个值
======================================================================================
						||		swap
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
		结构				||		swap = 95(0x5f)
======================================================================================
						||		...，value2，value1  →
	   操作数栈			||------------------------------------------------------------
						||		„，value1，value2
======================================================================================
						||
		描述				||		交换操作数栈顶的两个值。
swap 指令只有在 value1 和 value2 都是(§2.11.1)中定义的分类一的运算类型才能使用。
Java 虚拟机未提供交换操作数栈中两个分类二数值的指令。
						||
======================================================================================
						||		
	   运行时异常			||
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