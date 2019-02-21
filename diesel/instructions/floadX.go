package instructions

import (
	"../runtime"
	"../../utils"
	"../oil/types"
	"reflect"
	"../variator"
	)

type I_floadX struct {
}

func init()  {
	INSTRUCTION_MAP[0x22] = &I_floadX{}
	INSTRUCTION_MAP[0x23] = &I_floadX{}
	INSTRUCTION_MAP[0x24] = &I_floadX{}
	INSTRUCTION_MAP[0x25] = &I_floadX{}
}

func (s I_floadX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "floadX exce >>>>>>>>>\n")

	index := uint32(ctx.Code[ctx.PC - 1]) - 0x22
	value, _ := ctx.CurrentAborigines.GetAborigines(uint32(index))
	if reflect.TypeOf(value) != reflect.TypeOf(types.Jfloat(0)) &&
		reflect.TypeOf(value) != reflect.TypeOf(types.JDO) &&
		reflect.TypeOf(value) != reflect.TypeOf(types.JDU) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jfloat(value.(types.Jfloat)))
	return nil
}

func (s I_floadX)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jchar{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(2))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, types.Jfloat(1234.5678))
	return &runtime.Context{
		Code: []byte{0x22, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		从局部变量表加载一个 float 类型值到操作数栈中
======================================================================================
						||		fload_<n>
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
						||		fload_0 = 34(0x22)
						||------------------------------------------------------------
						||		fload_1 = 35(0x23)
		结构				||------------------------------------------------------------
						||		fload_2 = 36(0x24)
						||------------------------------------------------------------
						||		fload_3 = 37(0x25)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||		
						||		<n>和<n>+1 共同代表一个当前栈帧(§2.6)中局部变量表的索引值，<n> 作为索引定位的局部变量必须为 float 类型，
		描述				||		记作 value。指令执行后，value 将会压入到操作数栈栈顶
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
		注意				||	fload_<n>指令族中的每一条指令都与使用<n>作为 index 参数的 fload 指令作的作用一致，仅仅除了操作数<n>是隐式包含在指令中这点不同而已。
						||
						||
						||
======================================================================================
 */