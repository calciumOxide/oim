package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_lstoreX struct {
}

func init()  {
	INSTRUCTION_MAP[0x3f] = &I_lstoreX{}
	INSTRUCTION_MAP[0x40] = &I_lstoreX{}
	INSTRUCTION_MAP[0x41] = &I_lstoreX{}
	INSTRUCTION_MAP[0x42] = &I_lstoreX{}
}

func (s I_lstoreX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lstoreX exce >>>>>>>>>\n")

	index := ctx.Code[ctx.PC - 1] - 0x3f
	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentAborigines.SetAborigines(uint32(index), value)
	return nil
}

func (s I_lstoreX)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jlong(9))
	f.PushFrame(types.Jlong(9))
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
		操作				||		将一个 long 类型数据保存到局部变量表中
======================================================================================
						||		istore_<n>
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
		结构				||		lstore_0 = 63(0x3f)
						||------------------------------------------------------------
						||		lstore_1 = 64(0x40)
						||------------------------------------------------------------
						||		lstore_2 = 65(0x41)
						||------------------------------------------------------------
						||		lstore_3 = 66(0x42)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...
======================================================================================
						||
		描述				||		<n>与<n>+1 共同表示一个当前栈帧(§2.6)局部变量表的索引值，而在操 作数栈栈顶的 value 必须是 long 类型的数据，这个数据将从操作数栈出栈，
然后保存到<n>及<n>+1 所指向的局部变量表位置中。
======================================================================================
						||		
	   运行时异常			||
						||
======================================================================================
						||
		注意				||
						||lstore_<n>指令族中的每一条指令都与使用<n>作为 index 参数的 lstore
指令作的作用一致，仅仅除了操作数<n>是隐式包含在指令中这点不同而已。
						||
						||
======================================================================================
 */