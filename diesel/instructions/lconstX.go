package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
		)

type I_lconstX struct {
}

func init()  {
	INSTRUCTION_MAP[0x9] = &I_lconstX{}
	INSTRUCTION_MAP[0xA] = &I_lconstX{}
}

func (s I_lconstX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lconstX exce >>>>>>>>>\n")

	op := ctx.Code[ctx.PC - 1]

	ctx.CurrentFrame.PushFrame(types.Jlong(uint32(op) - 0x9))
	return nil
}

func (s I_lconstX)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jfloat(9.123456789012345))
	f.PushFrame(types.Jfloat(9.123456789012343))
	//f.PushFrame(nil)
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, types.Jfloat(1), types.JDO, types.JDU, types.JDN)
	return &runtime.Context{
		Code: []byte{0x2},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		将 long 类型数据压入到操作数栈中
======================================================================================
						||		iconst_<i>
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
						||		lconst_0 = 9(0x9)
		结构				||------------------------------------------------------------
						||		lconst_1 = 10(0xa)
=====================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，<l>
======================================================================================
						||		
						||
		描述				||			将 long 类型的常量<l>(0 或者 1)压入到操作数栈中。
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
		注意				||		iconst_<i>指令族中的每一条指令都与使用<i>作为参数的 bipush 指令作的作用一致，仅仅除了操作数<i>是隐式包含在指令中这点不同而已。
						||
======================================================================================
 */