package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_lstore struct {
}

func init()  {
	INSTRUCTION_MAP[0x37] = &I_lstore{}
}

func (s I_lstore)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lstore exce >>>>>>>>>\n")

	index := uint32(0)
	if ctx.PopWide() {
		index = uint32(utils.BigEndian2Little4U2(ctx.Code[ctx.PC : ctx.PC + 2]))
		ctx.PC += 2
	} else {
		index = uint32(ctx.Code[ctx.PC])
		ctx.PC += 1
	}
	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentAborigines.SetAborigines(uint32(index), value)
	return nil
}

func (s I_lstore)Test(octx *runtime.Context) *runtime.Context {
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
						||		lstore
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
		结构				||		lstore = 55(0x37)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...
======================================================================================
						||
		描述				||		index 是一个无符号 byte 型整数，它与 index 共同表示一个当前栈帧(§ 2.6)局部变量表的索引值，而在操作数栈栈顶的 value 必须是 long 类型的
数据，这个数据将从操作数栈出栈，然后保存到 index 及 index+1 所指向 的局部变量表位置中。
======================================================================================
						||		
	   运行时异常			||
						||
======================================================================================
						||
						||
		注意				||
						||	lstore 指令可以与 wide 指令联合使用，以实现使用 2 字节宽度的无符号整 数作为索引来访问局部变量表。
						||
						||
======================================================================================
 */