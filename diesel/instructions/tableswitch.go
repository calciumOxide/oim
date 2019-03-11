package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_tableswitch struct {
}

func init()  {
	INSTRUCTION_MAP[0xaa] = &I_tableswitch{}
}

func (s I_tableswitch)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "tableswitch exce >>>>>>>>>\n")

	switchPc := ctx.PC
	for ; ctx.PC % 4 != 0; ctx.PC++ {}

	defaultPc := uint32(ctx.Code[ctx.PC])<<24 | uint32(ctx.Code[ctx.PC + 1])<<16 | uint32(ctx.Code[ctx.PC+2])<<16 | uint32(ctx.Code[ctx.PC+3])
	ctx.PC += 4

	low := uint32(ctx.Code[ctx.PC])<<24 | uint32(ctx.Code[ctx.PC + 1])<<16 | uint32(ctx.Code[ctx.PC+2])<<16 | uint32(ctx.Code[ctx.PC+3])
	ctx.PC += 4
	height := uint32(ctx.Code[ctx.PC])<<24 | uint32(ctx.Code[ctx.PC + 1])<<16 | uint32(ctx.Code[ctx.PC+2])<<16 | uint32(ctx.Code[ctx.PC+3])
	ctx.PC += 4

	key, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(key) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	offset := defaultPc
	for i := low; i < height; i++ {
		if uint32(key.(types.Jint)) == i {
			offset = utils.BigEndian2Little4U4(ctx.Code[ctx.PC + i * 8 : ctx.PC + 4 + i * 8])
			break
		}
	}
	ctx.PC = switchPc + offset
	return nil
}

func (s I_tableswitch)Test(octx *runtime.Context) *runtime.Context {
	octx.PC += 1
	octx.CurrentFrame.PushFrame(types.Jint(1))
	return octx
}
/**
======================================================================================
		操作				||		根据索引值在跳转表中寻找配对的分支并进行跳转
======================================================================================
						||		tableswitch
						||------------------------------------------------------------
						||		<0-3 byte pad>
						||------------------------------------------------------------
						||		defaultbyte1
						||------------------------------------------------------------
						||		defaultbyte2
						||------------------------------------------------------------
						||		defaultbyte3
						||------------------------------------------------------------
						||		defaultbyte4
		格式				||------------------------------------------------------------
						||		lowbyte1
						||------------------------------------------------------------
						||		lowbyte2
						||------------------------------------------------------------
						||		lowbyte3
						||------------------------------------------------------------
						||		lowbyte4
						||------------------------------------------------------------
						||		hightbyte1
						||------------------------------------------------------------
						||		hightbyte2
						||------------------------------------------------------------
						||		hightbyte3
						||------------------------------------------------------------
						||		hightbyte4
						||------------------------------------------------------------
						||		jump offsets...
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		tableswitch = 170(0xaa)
======================================================================================
						||		...，index →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
		描述				||		tableswitch 是一条变长指令。紧跟 tableswitch 之后的 0 至 3 个字节作 为空白填充，
而后面 defaultbyte1 至 defaultbyte4 等代表了一个个由 4 个字节组成的、从当前方法开始(第一条操作码指令)计算的地址，
即紧跟随 空白填充的是一系列 32 位有符号整数值: 包括默认跳转地址 default、高位值 high 以及低位值 low。
在此之后，是 high-low+1 个有符号 32 位偏移 量 offset，其中要求 low 小于或等于 high。
这 high-low+1 个 32 位有 符号数值形成一张零基址跳转表(0-Based Jump Table)，
所有上述的32 位有符号数都以(byte1 << 24)|(byte2 << 16)|(byte3 << 8)| byte4 方式构成。
指令执行时，int 型的 index 从操作数栈中出栈，如果 index 比 low 值小或 者比 high 值大，那就是用 default 作为目标地址进行跳转。
否则，在跳转 表中第 index-low 个地址值将作为目标地址进行跳转，程序从目标地址开始 继续执行。
目标地址既可能从跳转表匹配坐标中得出，也可能从 default 中得出，但无 论如何，最终的目标地址必须在包含 tableswitch 指令的那个方法之内。


当且仅当包含 tableswitch 指令的方法刚好位于 4 字节边界上， lookupswitch 指令才能确保它的所有操作数都是 4 直接对齐的。
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
						||当且仅当包含 tableswitch 指令的方法刚好位于 4 字节边界上， tableswitch 指令才能确保它的所有操作数都是 4 直接对齐的。
         所有的匹配坐标以有序方式存储是为了查找时的效率考虑。
						||
						||
======================================================================================
 */