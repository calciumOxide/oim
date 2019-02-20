package instructions

import (
	"../runtime"
	"../../utils"
)

type I_aconstNull struct {
}

func init()  {
	INSTRUCTION_MAP[0x01] = &I_aconstNull{}
}

func (s I_aconstNull)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "aconstNull exce >>>>>>>>>\n")

	frame := ctx.CurrentFrame
	frame.PushFrame(nil)

	return nil
}

func (s I_aconstNull)Test() *runtime.Context {
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
		操作				||		将一个 null 值压入到操作数栈中
======================================================================================
						||		aconst_null
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
		结构				||		aconst_null = 1(0x1)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，null
======================================================================================
						||		
						||		
						||
		描述				||		将一个 null 值压入到操作数栈中。
						||		
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
		注意				||		《Java 虚拟机规范》并没有强制规定 null 值的在虚拟机的内存中是如何实 际表示的。
						||
						||
						||
======================================================================================
 */