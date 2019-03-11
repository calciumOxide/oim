package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_ifX struct {
}

func init()  {
	INSTRUCTION_MAP[0x99] = &I_ifX{}
	INSTRUCTION_MAP[0x9a] = &I_ifX{}
	INSTRUCTION_MAP[0x9b] = &I_ifX{}
	INSTRUCTION_MAP[0x9c] = &I_ifX{}
	INSTRUCTION_MAP[0x9d] = &I_ifX{}
	INSTRUCTION_MAP[0x9e] = &I_ifX{}
}

func (s I_ifX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "ifX exce >>>>>>>>>\n")

	op := ctx.Code[ctx.PC - 1]
	branch := uint32(ctx.Code[ctx.PC]) << 8 | uint32(ctx.Code[ctx.PC + 1])
	ctx.PC += 2
	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	if (value.(types.Jint) == 0 && op == 0x99) || (value != 0 && op == 0x9a) ||
		(value.(types.Jint) < 0 && op == 0x9b) || (value.(types.Jint) <= 0 && op == 0x9e) ||
		(value.(types.Jint) > 0 && op == 0x9d) || (value.(types.Jint) >= 0 && op == 0x9c) {
		ctx.PC = branch
	}
	return nil
}

func (s I_ifX)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(&types.Jreference{})
	f.PushFrame(types.Jint(0))
	jr, _ := f.PeekFrame()
	f.PushFrame(jr)
	//f.PushFrame(nil)
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x99, 0x0, 12},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		整数与零比较的条件分支判断
======================================================================================
						||		if<cond>
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
						||		ifeq = 153(0x99)
		结构				||------------------------------------------------------------
						||		ifne = 154(0x9a)
		结构				||------------------------------------------------------------
						||		iflt = 155(0x9b)
		结构				||------------------------------------------------------------
						||		ifge = 156(0x9c)
		结构				||------------------------------------------------------------
						||		ifgt = 157(0x9d)
		结构				||------------------------------------------------------------
						||		ifle = 158(0x9e)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
						||
						||		value 必须为 int 类型数据，指令执行时，value 从操作数栈中出栈，
						||		然后进行比较运算(所有比较都是带符号的)，比较的规则如下:
						||			 eq 当且仅当 value=0 比较的结果为真。
		描述				||			 ne 当且仅当 value≠0 比较的结果为真。
		描述				||			 lt 当且仅当 value<0 比较的结果为真。
		描述				||			 le 当且仅当 value≤0 比较的结果为真。
		描述				||			 ge 当且仅当 value>0 比较的结果为真。
		描述				||			 ge 当且仅当 value≥0 比较的结果为真。
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