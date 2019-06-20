package instructions

import (
	"../../utils"
	"../oil/types"
	"../runtime"
	"../variator"
	"reflect"
)

type I_iloadX struct {
}

func init() {
	INSTRUCTION_MAP[0x1a] = &I_iloadX{}
	INSTRUCTION_MAP[0x1b] = &I_iloadX{}
	INSTRUCTION_MAP[0x1c] = &I_iloadX{}
	INSTRUCTION_MAP[0x1d] = &I_iloadX{}
}

func (s I_iloadX) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "iloadX exce >>>>>>>>>\n")

	index := uint32(ctx.Code[ctx.PC-1]) - 0x1a
	value, _ := ctx.CurrentAborigines.GetAborigines(uint32(index))
	if reflect.TypeOf(value) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(value)
	return nil
}

func (s I_iloadX) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jchar{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(2))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, types.Jint(1234))
	return &runtime.Context{
		Code:              []byte{0x1a, 0x0},
		CurrentFrame:      f,
		CurrentAborigines: a,
	}
}

/**
======================================================================================
		操作				||		从局部变量表加载一个 int 类型值到操作数栈中
======================================================================================
						||		iload_<n>
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
						||		iload_0 = 26(0x1a)
						||------------------------------------------------------------
						||		iload_1 = 27(0x1b)
		结构				||------------------------------------------------------------
						||		iload_2 = 28(0x1c)
						||------------------------------------------------------------
						||		iload_3 = 29(0x1d)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||
						||		<n>代表一个当前栈帧(§2.6)中局部变量表的索引值，<n>作为索引定位的 局部变量必须为 int 类型，记作 value。
		描述				||		指令执行后，value 将会压入到操作数栈栈顶
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
		注意				||	iload_<n>指令族中的每一条指令都与使用<n>作为 index 参数的 iload 指令作的作用一致，仅仅除了操作数<n>是隐式包含在指令中这点不同而已。
						||
======================================================================================
*/
