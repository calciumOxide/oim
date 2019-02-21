package instructions

import (
	"../runtime"
	"../../utils"
	"../oil/types"
	"../variator"
	)

type I_dupX1 struct {
}

func init()  {
	INSTRUCTION_MAP[0x5a] = &I_dupX1{}
}

func (s I_dupX1)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "dupX1 exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()
	if runtime.IsDoubleLong(value1) || runtime.IsDoubleLong(value2) {
		except, _ := variator.AllocExcept(variator.InstructionException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(value1)
	ctx.CurrentFrame.PushFrame(value2)
	ctx.CurrentFrame.PushFrame(value1)
	return nil
}

func (s I_dupX1)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jdouble(8.123456789012345))
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
		操作				||		复制操作数栈栈顶的值，并插入到栈顶以下 2 个值之后
======================================================================================
						||		dup_x1
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
		结构				||		dup_x1 = 90(0x5a)
======================================================================================
						||		...，value2，value1 →
	   操作数栈			||------------------------------------------------------------
						||		...，value1，value2，value1
======================================================================================
						||
		描述				||		复制操作数栈栈顶的值，并将此值压入到操作数栈顶以下 2 个值之后。
						||		如果 value1 和 value2 不是§2.11.1 的表 2.3 中列出的分类一中的数据类型，就不能使用 dup_x1 指令来复制栈顶值
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