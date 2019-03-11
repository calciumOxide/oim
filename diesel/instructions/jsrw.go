package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
			)

type I_jsrw struct {
}

func init()  {
	INSTRUCTION_MAP[0xa8] = &I_jsrw{}
}

func (s I_jsrw)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "jsrw exce >>>>>>>>>\n")

	index := uint32(ctx.Code[ctx.PC]) << 24 | uint32(ctx.Code[ctx.PC + 1]) << 16 | uint32(ctx.Code[ctx.PC + 2]) << 8 | uint32(ctx.Code[ctx.PC + 3])
	ctx.PC += 4

	ctx.CurrentFrame.PushFrame(types.Jaddress(index))
	return nil
}

func (s I_jsrw)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(9))
	f.PushFrame(types.Jint(9))
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
		操作				||		程序段落跳转
======================================================================================
						||		jsr_w
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
		结构				||		jsr_w = 201(0xc9)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，address
======================================================================================
						||
						||		address 是一个 returnAddress 类型的数据，它由 jsr_w 指令推入操作数 栈中。无符号 byte 型数据 branchbyte1、branchbyte2、branchbyte3
和 branchbyte4 用于构建一个 32 位有符号的分支偏移量，构建方式为 (branchbyte1 << 24)|(branchbyte1 << 16)|(branchbyte1 << 8)|branchbyte2。指令执行时，
将产生一个当前位置的偏移坐标，并压入 到操作数栈中。跳转目标地址必须在 jsr_w 指令所在的方法之内。
		描述				||
						||
======================================================================================
						||		
	   运行时异常			||
						||
======================================================================================
						||
						||
						||jsr_w指令被用来与ret指令一同实现Java语言中的finally语句块(参 见§3.13“编译 fianlly 语句块”)。请注意，jsr_w 指令推送 address 到 操作数栈，ret 指令从局部变量表中把它取出，这种不对称的操作是故意设计的。
虽然 jsr_w 指令拥有 4 个字节的分支偏移量，但是其他因素限定了一个方法 的最大长度不能超过 65535 个字节(§4.11)。这个上限值可能会在将来发布 的 Java 虚拟机中被提升。
		注意				||
						||
						||
						||
======================================================================================
 */