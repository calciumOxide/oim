package instructions

import (
	"../runtime"
	"../../utils"
	"../oil/types"
	"reflect"
	"../variator"
	)

type I_dloadX struct {
}

func init()  {
	INSTRUCTION_MAP[0x26] = &I_dloadX{}
	INSTRUCTION_MAP[0x27] = &I_dloadX{}
	INSTRUCTION_MAP[0x28] = &I_dloadX{}
	INSTRUCTION_MAP[0x29] = &I_dloadX{}
}

func (s I_dloadX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "dloadX exce >>>>>>>>>\n")

	index := uint32(ctx.Code[ctx.PC - 1]) - 0x26
	value, _ := ctx.CurrentAborigines.GetAborigines(uint32(index))
	if reflect.TypeOf(value) != reflect.TypeOf(types.Jdouble(0)) &&
		reflect.TypeOf(value) != reflect.TypeOf(types.JDO) &&
		reflect.TypeOf(value) != reflect.TypeOf(types.JDU) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jdouble(value.(types.Jdouble)))
	return nil
}

func (s I_dloadX)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jchar{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(2))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, types.Jdouble(1234.5678))
	return &runtime.Context{
		Code: []byte{0x26, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		从局部变量表加载一个 double 类型值到操作数栈中
======================================================================================
						||		dload_<n>
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
						||		dload_0 = 38(0x26)
						||------------------------------------------------------------
						||		dload_1 = 39(0x27)
		结构				||------------------------------------------------------------
						||		dload_2 = 40(0x28)
						||------------------------------------------------------------
						||		dload_3 = 41(0x29)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||		
						||		<n>和<n>+1 共同代表一个当前栈帧(§2.6)中局部变量表的索引值，<n> 作为索引定位的局部变量必须为 double 类型，
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
		注意				||	dload_<n>指令族中的每一条指令都与使用<n>作为 index 参数的 dload 指令作的作用一致，仅仅除了操作数<n>是隐式包含在指令中这点不同而已。
						||
						||
						||
======================================================================================
 */