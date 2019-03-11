package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"../variator"
	)

type I_dup2X1 struct {
}

func init()  {
	INSTRUCTION_MAP[0x5d] = &I_dup2X1{}
}

func (s I_dup2X1)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "dup2X1 exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()
	value3, _ := ctx.CurrentFrame.PeekFrame()
	if !runtime.IsDoubleLong(value3) || !runtime.IsDoubleLong(value2) || !runtime.IsDoubleLong(value1) {
		ctx.CurrentFrame.PopFrame()
		ctx.CurrentFrame.PushFrame(value2)
		ctx.CurrentFrame.PushFrame(value1)
		ctx.CurrentFrame.PushFrame(value3)
		ctx.CurrentFrame.PushFrame(value2)
		ctx.CurrentFrame.PushFrame(value1)
		return nil
	}
	if !runtime.IsDoubleLong(value2) || runtime.IsDoubleLong(value1) {
		ctx.CurrentFrame.PushFrame(value1)
		ctx.CurrentFrame.PushFrame(value2)
		ctx.CurrentFrame.PushFrame(value1)
		return nil
	}
	ctx.CurrentFrame.PushFrame(value2)
	ctx.CurrentFrame.PushFrame(value1)
	except, _ := variator.AllocExcept(variator.InstructionException)
	ctx.Throw(except)
	return nil
}

func (s I_dup2X1)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jdouble(8.123456789012345))
	f.PushFrame(types.Jdouble(9.123456789012345))
	f.PushFrame(types.Jint(3))
	f.PushFrame(types.Jint(2))
	f.PushFrame(types.Jint(1))
	f.PushFrame(types.Jdouble(9.123456789012345))

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
		操作				||		复制操作数栈栈顶 1 个或 2 个值，并插入到栈顶以下 2 个或 3 个值之后
======================================================================================
						||		dup2_x1
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
		结构				||		dup2_x1 = 93(0x5d)
======================================================================================
						||		...，value3，value2，value1 →
	   操作数栈			||-------------------------------------------------当 value1、value2 和 value3 都是§2.11.1 的表 2.3 中列出的分类一中 的数据类型时满足结构 1。
						||		...，value2，value1，value3，value2，value1
======================================================================================
						||		...，value2，value1 →
	   操作数栈			||-------------------------------------------------当 value1 是§2.11.1 的表 2.3 中列出的分类二中的数据类型，而 value2 是分类一的数据类型时满足结构 2。
						||		...，value1，value2，value1
======================================================================================
						||
		描述				||		复制操作数栈栈顶 1 个或 2 个值，并按照原有的顺序插入到栈顶以下 2 个或 3 个值之后。
						||
						||
======================================================================================
						||		
						||
						||
	   运行时异常			||
						||		
						||		
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