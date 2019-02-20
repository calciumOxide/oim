package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
		)

type I_if_acmpX struct {
}

func init()  {
	INSTRUCTION_MAP[0xa5] = &I_if_acmpX{}
	INSTRUCTION_MAP[0xa6] = &I_if_acmpX{}
}

func (s I_if_acmpX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "if_acmpX exce >>>>>>>>>\n")

	op := ctx.Code[ctx.PC - 1]
	branch := uint32(ctx.Code[ctx.PC]) << 8 | uint32(ctx.Code[ctx.PC + 1])
	ctx.PC += 2
	value2, _ := ctx.CurrentFrame.PopFrame()
	value1, _ := ctx.CurrentFrame.PopFrame()

	if (value1 == value2 && op == 0xa5) || (value1 != value2 && op == 0xa6) {
		ctx.PC = branch
	}
	return nil
}

func (s I_if_acmpX)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(&types.Jreference{})
	f.PushFrame(&types.Jreference{})
	jr, _ := f.PeekFrame()
	f.PushFrame(jr)
	f.PushFrame(nil)
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0xa5, 0x0, 12},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		reference 数据的据条件分支判断
======================================================================================
						||		if_acmp<cond>
						||------------------------------------------------------------
						||		branchbyte1
						||------------------------------------------------------------
						||		branchbyte2
		格式				||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
======================================================================================
						||		if_acmpeq = 165(0xa5)
		结构				||------------------------------------------------------------
						||		if_acmpne = 166(0xa6)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
						||
						||		value1 和 value2 都必须为 reference 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		然后进行比较运算，比较的规则如下:
						||			 eq 当且仅当 value1=value2 比较的结果为真。
		描述				||			 ne 当且仅当 value1≠value2 比较的结果为真。
						||
						||		如果比较结果为真，那无符号 byte 型数据 branchbyte1 和 branchbyte2 用于构建一个 16 位有符号的分支偏移量，
						||		构建方式为(branchbyte1 << 8) | branchbyte2。
						||		指令执行后，程序将会转到这个if_acmp<cond>指令之 后的，由上述偏移量确定的目标地址上继续执行。
						||		这个目标地址必须处于 if_acmp<cond>指令所在的方法之中。
						||		另外，如果比较结果为假，那程序将继续执行 if_acmp<cond>指令后面的其 他直接码指令。
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
		注意				||
						||
						||
======================================================================================
 */