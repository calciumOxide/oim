package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_lloadX struct {
}

func init()  {
	INSTRUCTION_MAP[0x1e] = &I_lloadX{}
	INSTRUCTION_MAP[0x1f] = &I_lloadX{}
	INSTRUCTION_MAP[0x20] = &I_lloadX{}
	INSTRUCTION_MAP[0x21] = &I_lloadX{}
}

func (s I_lloadX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lloadX exce >>>>>>>>>\n")

	index := uint32(ctx.Code[ctx.PC - 1]) - 0x1e
	value, _ := ctx.CurrentAborigines.GetAborigines(uint32(index))
	if reflect.TypeOf(value) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(value)
	return nil
}

func (s I_lloadX)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jchar{1, 2, 3, 4},
	})
	f.PushFrame(types.Jlong(2))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, types.Jlong(1234))
	return &runtime.Context{
		Code: []byte{0x1a, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		从局部变量表加载一个 long 类型值到操作数栈中
======================================================================================
						||		lload_<n>
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
						||		lload_0 = 30(0x1e)
						||------------------------------------------------------------
						||		lload_1 = 31(0x1f)
		结构				||------------------------------------------------------------
						||		lload_2 = 32(0x20)
						||------------------------------------------------------------
						||		lload_3 = 33(0x21)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||		
		描述				||		<n>与<n>+1 共同构成一个当前栈帧(§2.6)中局部变量表的索引的，<n> 作为索引定位的局部变量必须为 long 类型，记为 value。指令执行后，value
将会压入到操作数栈栈顶
						||
======================================================================================
						||		
						||
	   运行时异常			||
						||		
						||		
						||		
======================================================================================
						||
		注意				||	lload_<n>指令族中的每一条指令都与使用<n>作为 index 参数的 lload 指
令作的作用一致，仅仅除了操作数<n>是隐式包含在指令中这点不同而已。
						||
======================================================================================
 */