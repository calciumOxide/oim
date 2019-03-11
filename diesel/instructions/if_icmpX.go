package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_if_icmpX struct {
}

func init()  {
	INSTRUCTION_MAP[0x9f] = &I_if_icmpX{}
	INSTRUCTION_MAP[0xa0] = &I_if_icmpX{}
	INSTRUCTION_MAP[0xa1] = &I_if_icmpX{}
	INSTRUCTION_MAP[0xa2] = &I_if_icmpX{}
	INSTRUCTION_MAP[0xa3] = &I_if_icmpX{}
	INSTRUCTION_MAP[0xa4] = &I_if_icmpX{}
}

func (s I_if_icmpX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "if_icmpX exce >>>>>>>>>\n")

	op := ctx.Code[ctx.PC - 1]
	branch := uint32(ctx.Code[ctx.PC]) << 8 | uint32(ctx.Code[ctx.PC + 1])
	ctx.PC += 2
	value2, _ := ctx.CurrentFrame.PopFrame()
	value1, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jint(0)) ||  reflect.TypeOf(value2) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	if (value1 == value2 && op == 0x9f) || (value1 != value2 && op == 0xa0) ||
		(value1.(types.Jint) < value2.(types.Jint) && op == 0xa1) || (value1.(types.Jint) <= value2.(types.Jint) && op == 0xa4) ||
		(value1.(types.Jint) > value2.(types.Jint) && op == 0xa3) || (value1.(types.Jint) >= value2.(types.Jint) && op == 0xa2) {
		ctx.PC = branch
	}
	return nil
}

func (s I_if_icmpX)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(&types.Jreference{})
	f.PushFrame(types.Jint(1))
	jr, _ := f.PeekFrame()
	f.PushFrame(jr)
	//f.PushFrame(nil)
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x9f, 0x0, 12},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		int 数值的条件分支判断
======================================================================================
						||		if_icmp<cond>
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
						||		if_icmpeq = 159(0x9f)
		结构				||------------------------------------------------------------
						||		if_icmpne = 160(0xa0)
		结构				||------------------------------------------------------------
						||		if_icmplt = 161(0xa1)
		结构				||------------------------------------------------------------
						||		if_icmpge = 162(0xa2)
		结构				||------------------------------------------------------------
						||		if_icmpgt = 163(0xa3)
		结构				||------------------------------------------------------------
						||		if_icmple = 164(0xa4)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
						||
						||		value1 和 value2 都必须为 int 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		然后进行比较运算(所有比较都是带符号的)，比较的规则如下:
						||			 eq 当且仅当 value1=value2 比较的结果为真。
		描述				||			 ne 当且仅当 value1≠value2 比较的结果为真。
		描述				||			 lt 当且仅当 value1<value2 比较的结果为真。
		描述				||			 le 当且仅当 value1≤value2 比较的结果为真。
		描述				||			 gt 当且仅当 value1>value2 比较的结果为真。
		描述				||			 ge 当且仅当 value1≥value2 比较的结果为真。
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