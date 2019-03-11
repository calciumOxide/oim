package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_istore struct {
}

func init()  {
	INSTRUCTION_MAP[0x36] = &I_istore{}
}

func (s I_istore)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "istore exce >>>>>>>>>\n")

	index := ctx.Code[ctx.PC]
	ctx.PC += 1
	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentAborigines.SetAborigines(uint32(index), value)
	return nil
}

func (s I_istore)Test(octx *runtime.Context) *runtime.Context {
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
						||		istore
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
		结构				||		istore = 54(0x36)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...
======================================================================================
						||
						||		index 是一个无符号 byte 型整数，它指向当前栈帧(§2.6)局部变量表的 索引值，
		描述				||		而在操作数栈栈顶的 value 必须是 int 类型的数据，这个数据将从操作数栈出栈，然后保存到 index 所指向的局部变量表位置中。
						||
======================================================================================
						||		
	   运行时异常			||
						||
======================================================================================
						||
						||
						||		istore 指令可以与 wide 指令联合使用，以实现使用 2 字节宽度的无符号整数作为索引来访问局部变量表。
		注意				||
						||
						||
						||
======================================================================================
 */