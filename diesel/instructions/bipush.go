package instructions

import (
	"../../utils"
	"../oil/types"
	"../runtime"
)

type I_bipush struct {
}

func init() {
	INSTRUCTION_MAP[0x10] = &I_bipush{}
}

func (s I_bipush) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "bipush exce >>>>>>>>>\n")

	value := ctx.Code[ctx.PC]
	ctx.PC += 1
	ctx.CurrentFrame.PushFrame(types.Jbyte(value))
	return nil
}

func (s I_bipush) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(2))
	f.PushFrame(types.Jbyte(33))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code:              []byte{0x0, 0x0F},
		CurrentFrame:      f,
		CurrentAborigines: a,
	}
}

/**
======================================================================================
		操作				||		将一个 byte 类型数据入栈
======================================================================================
						||		bipush
						||------------------------------------------------------------
						||		byte
						||------------------------------------------------------------
						||
		格式				||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||
======================================================================================
		结构				||		bipush = 16(0x10)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||
						||
		描述				||		将 byte 带符号扩展为一个 int 类型的值 value，然后将 value 压入到操作 数栈中。
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
