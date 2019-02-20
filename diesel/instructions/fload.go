package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_fload struct {
}

func init()  {
	INSTRUCTION_MAP[0x17] = &I_fload{}
}

func (s I_fload)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "fload exce >>>>>>>>>\n")

	index := ctx.Code[ctx.PC]
	ctx.PC += 1

	value, _ := ctx.CurrentAborigines.GetAborigines(uint32(index))
	if reflect.TypeOf(value) != reflect.TypeOf(types.Jfloat(0)) &&
		reflect.TypeOf(value) != reflect.TypeOf(types.JDO) &&
		reflect.TypeOf(value) != reflect.TypeOf(types.JDU) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(value)
	return nil
}

func (s I_fload)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jchar{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(2))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, types.Jfloat(1234.5678))
	return &runtime.Context{
		Code: []byte{0x17, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		从局部变量表加载一个 float 类型值到操作数栈中
======================================================================================
						||		fload
						||------------------------------------------------------------
						||		index
						||------------------------------------------------------------
						||		
		格式				||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		fload = 23(0x17)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||		
						||		index 是一个代表当前栈帧(§2.6)中局部变量表的索引的无符号 byte 类 型整数，
		描述				||		index 作为索引定位的局部变量必须为 float 类型(占用 index和 index+1 两个位置)，记为 value。
						||		指令执行后，value 将会压入到操作 数栈栈顶
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
		注意				||	fload 操作码可以与 wide 指令联合一起实现使用 2 个字节长度的无符号 byte 型数值作为索引来访问局部变量表。
						||
						||
						||
======================================================================================
 */