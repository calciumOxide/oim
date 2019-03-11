package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_lookupswitch struct {
}

func init()  {
	INSTRUCTION_MAP[0xab] = &I_lookupswitch{}
}

func (s I_lookupswitch)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lookupswitch exce >>>>>>>>>\n")

	switchPc := ctx.PC
	for ; ctx.PC % 4 != 0; ctx.PC++ {}

	defaultPc := uint32(ctx.Code[ctx.PC])<<24 | uint32(ctx.Code[ctx.PC + 1])<<16 | uint32(ctx.Code[ctx.PC+2])<<16 | uint32(ctx.Code[ctx.PC+3])
	ctx.PC += 4

	count := uint32(ctx.Code[ctx.PC])<<24 | uint32(ctx.Code[ctx.PC + 1])<<16 | uint32(ctx.Code[ctx.PC+2])<<16 | uint32(ctx.Code[ctx.PC+3])
	ctx.PC += 4
	beginPc := ctx.PC
	offset := defaultPc
	key, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(key) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	beginCase := utils.BigEndian2Little4U4(ctx.Code[beginPc : beginPc + 4])
	endCase := utils.BigEndian2Little4U4(ctx.Code[beginPc - 8 + count * 8 : beginPc - 4 + count * 8])

	if uint32(key.(types.Jint)) <= endCase && uint32(key.(types.Jint)) >= beginCase {
		for i := uint32(0); i < count; i++ {
			curCase := utils.BigEndian2Little4U4(ctx.Code[beginPc + i * 8 : beginPc + 4 + i * 8])
			if uint32(key.(types.Jint)) == curCase {
				offset = utils.BigEndian2Little4U4(ctx.Code[beginPc + 4 + i * 8 : beginPc + 8 + i * 8])
				break
			}
		}
	}
	ctx.PC = switchPc + offset
	return nil
}

func (s I_lookupswitch)Test(octx *runtime.Context) *runtime.Context {
	octx.PC += 1
	octx.CurrentFrame.PushFrame(types.Jint(1))
	return octx
}
/**
======================================================================================
		操作				||		根据键值在跳转表中寻找配对的分支并进行跳转
======================================================================================
						||		lookupswitch
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
						||		npairs1
						||------------------------------------------------------------
						||		npairs2
						||------------------------------------------------------------
						||		npairs3
						||------------------------------------------------------------
						||		npairs4
						||------------------------------------------------------------
						||		match-offset pairs...
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		lookupswitch = 171(0xab)
======================================================================================
						||		...，key →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
		描述				||		lookupswitch 是一条变长指令。紧跟 lookupswitch 之后的 0 至 3 个字节 作为空白填充，
而后面 defaultbyte1 至 defaultbyte4 等代表了一个个由 4 个字节组成的、从当前方法开始(第一条操作码指令)计算的地址，
即紧跟 随空白填充的是一系列 32 位有符号整数值:包括默认跳转地址 default、匹 配坐标的数量 npairs 以及 npairs 组匹配坐标。
其中，npairs 的值应当大于或等于 0，每一组匹配坐标都包含了一个整数值 match 以及一个有符号 32 位偏移量 offset。
上述所有的 32 位有符号数值都由以下形式构成:(byte1 << 24)|(byte2 << 16)|(byte3 << 8)| byte4。

lookupswitch 指令之后所有的匹配坐标必须以其中的 match 值排序，按照 升序储存。
指令执行时，int 型的 key 从操作数栈中出栈，与每一个 match 值相互比较。 如果能找到一个与之相等的 match 值，
那就就以这个 match 所配对的偏移量 offset 作为目标地址进行跳转。如果没有配对到任何一个 match 值，
那就是 用 default 作为目标地址进行跳转。程序从目标地址开始继续执行。
目标地址既可能从 npairs 组匹配坐标中得出，也可能从 default 中得出， 但无论如何，最终的目标地址必须在包含 lookupswitch 指令的那个方法之 内。

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
						||当且仅当包含 lookupswitch 指令的方法刚好位于 4 字节边界上， lookupswitch 指令才能确保它的所有操作数都是 4 直接对齐的。
         所有的匹配坐标以有序方式存储是为了查找时的效率考虑。
						||
						||
======================================================================================
 */