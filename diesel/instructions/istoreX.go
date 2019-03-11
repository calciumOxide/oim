package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_istoreX struct {
}

func init()  {
	INSTRUCTION_MAP[0x3b] = &I_istoreX{}
	INSTRUCTION_MAP[0x3c] = &I_istoreX{}
	INSTRUCTION_MAP[0x3d] = &I_istoreX{}
	INSTRUCTION_MAP[0x3e] = &I_istoreX{}
}

func (s I_istoreX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "istoreX exce >>>>>>>>>\n")

	index := ctx.Code[ctx.PC - 1] - 0x3b
	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentAborigines.SetAborigines(uint32(index), value)
	return nil
}

func (s I_istoreX)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		将一个 int 类型数据保存到局部变量表中
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
		结构				||		istore_0 = 59(0x3b)
						||------------------------------------------------------------
						||		istore_1 = 60(0x3c)
						||------------------------------------------------------------
						||		istore_2 = 61(0x3d)
						||------------------------------------------------------------
						||		istore_3 = 62(0x3e)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...
======================================================================================
						||
						||		<n>必须是一个指向当前栈帧(§2.6)局部变量表的索引值，
		描述				||		而在操作数栈 栈顶的 value 必须是 int 类型的数据，这个数据将从操作数栈出栈，然后保存到<n>所指向的局部变量表位置中。
						||
======================================================================================
						||		
	   运行时异常			||
						||
======================================================================================
						||
						||
						||		istore_<n>指令族中的每一条指令都与使用<n>作为 index 参数的 istore指令作的作用一致，仅仅除了操作数<n>是隐式包含在指令中这点不同而已。
		注意				||
						||
						||
						||
======================================================================================
 */