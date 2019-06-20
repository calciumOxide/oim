package instructions

import (
	"../../utils"
	"../oil/types"
	"../runtime"
	"../variator"
)

type I_dup2 struct {
}

func init() {
	INSTRUCTION_MAP[0x5c] = &I_dup2{}
}

func (s I_dup2) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "dup2 exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PeekFrame()
	ctx.CurrentFrame.PushFrame(value1)
	if !runtime.IsDoubleLong(value1) && !runtime.IsDoubleLong(value2) {
		ctx.CurrentFrame.PopFrame()
		ctx.CurrentFrame.PushFrame(value2)
		ctx.CurrentFrame.PushFrame(value1)
		ctx.CurrentFrame.PushFrame(value2)
		ctx.CurrentFrame.PushFrame(value1)
		return nil
	}
	if runtime.IsDoubleLong(value2) {
		ctx.CurrentFrame.PushFrame(value1)
		ctx.CurrentFrame.PushFrame(value1)
		return nil
	}
	ctx.CurrentFrame.PushFrame(value1)
	except, _ := variator.AllocExcept(variator.InstructionException)
	ctx.Throw(except)
	return nil
}

func (s I_dup2) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(8))
	f.PushFrame(types.Jdouble(9.123456789012345))
	f.PushFrame(types.Jint(6))
	f.PushFrame(types.Jint(7))

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
		操作				||		复制操作数栈栈顶 1 个或 2 个值，并插入到栈顶
======================================================================================
						||		dup2
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
		结构				||		dup2 = 92(0x5c)
======================================================================================
						||		...，value2，value1 →
	   操作数栈			||-----------------------------------------当 value1 和 value2 都是§2.11.1 的表 2.3 中列出的分类一中的数据类型 时满足结构 1。
						||		...，value2，value1，value2，value1
						==============================================================
						||		...，value →
	   操作数栈			||-----------------------------------------当 value 是§2.11.1 的表 2.3 中列出的分类二中的数据类型时满足结构 2。
						||		...，value，value
======================================================================================
						||
		描述				||		复制操作数栈栈顶 1 个或 2 个值，并将这些值按照原来的顺序压入到操作数栈 顶。
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
