package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_lload struct {
}

func init()  {
	INSTRUCTION_MAP[0x16] = &I_lload{}
}

func (s I_lload)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lload exce >>>>>>>>>\n")

	index := ctx.Code[ctx.PC]
	ctx.PC += 1

	value, _ := ctx.CurrentAborigines.GetAborigines(uint32(index))
	if reflect.TypeOf(value) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(value)
	return nil
}

func (s I_lload)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jchar{1, 2, 3, 4},
	})
	f.PushFrame(types.Jlong(2))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, types.Jlong(1234))
	return &runtime.Context{
		Code: []byte{0x15, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		从局部变量表加载一个 long 类型值到操作数栈中
======================================================================================
						||		lload
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
		结构				||		iload = 22(0x16)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||		
		描述				||		index 是一个无符号 byte 类型整数，它与 index+1 共同构成一个当前栈帧 (§2.6)中局部变量表的索引的，index 作为索引定位的局部变量必须为
long 类型，记为 value。指令执行后，value 将会压入到操作数栈栈顶
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
		注意				||	lload 操作码可以与 wide 指令联合一起实现使用 2 个字节长度的无符号byte 型数值作为索引来访问局部变量表。
						||
						||
======================================================================================
 */