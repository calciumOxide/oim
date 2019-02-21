package instructions

import (
	"../runtime"
	"../../utils"
	"../oil/types"
				)

type I_gotoW struct {
}

func init()  {
	INSTRUCTION_MAP[0xc8] = &I_gotoW{}
}

func (s I_gotoW)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "gotoW exce >>>>>>>>>\n")

	offset := (uint32(ctx.Code[ctx.PC]) << 24) | (uint32(ctx.Code[ctx.PC + 1]) << 16) | (uint32(ctx.Code[ctx.PC + 2]) << 8) | uint32(ctx.Code[ctx.PC + 3])
	ctx.PC += 4
	ctx.PC = uint32(offset)

	return nil
}

func (s I_gotoW)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	fields := make(map[string]interface{})
	fields["app"] = "bb"
	f.PushFrame(types.Jreference{
		Reference: types.Jobject{
			ClassTypeIndex: 4,
			Fileds: fields,
		},
	})
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x0, 0x0, 0x0, 0xF, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		分支跳转(宽范围)
======================================================================================
						||		gotoW
						||------------------------------------------------------------
						||		branchbyte1
						||------------------------------------------------------------
						||		branchbyte2
		格式				||------------------------------------------------------------
						||		branchbyte3
						||------------------------------------------------------------
						||		branchbyte4
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		goto_w = 200(0xc8)
======================================================================================
						||		...， →   无改变
	   操作数栈			||------------------------------------------------------------
						||		...， →
======================================================================================
						||		
						||		无符号 byte 型数据 branchbyte1、branchbyte2、branchbyte3 和 branchbyte4 用于构建一个 32 位有符号的分支偏移量，构建方式为(branchbyte1 << 24)|(branchbyte1 << 16)|(branchbyte1 << 8)|branchbyte2。
		描述				||		指令执行后，程序将会转到这个 goto_w 指令之后的， 由上述偏移量确定的目标地址上继续执行。这个目标地址必须处于 goto_w 指 令所在的方法之中。
						||
======================================================================================
						||		
						||
						||
	   链接时异常			||
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
		注意				||尽管 goto_w 指令拥有 4 字节宽度的分支偏移量，但是还受到方法最大字节码 长度为 65535 字节(§4.11)的限制，这个限制值可能会在未来的 Java 虚拟机版本中增大。
						||
						||
						||
======================================================================================
 */