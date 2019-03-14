package instructions

import (
	"../runtime"
	"../../utils"
	"../variator"
	"../../types"
	"reflect"
)

type I_dstore struct {
}

func init()  {
	INSTRUCTION_MAP[0x39] = &I_dstore{}
}

func (s I_dstore)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "dstore exce >>>>>>>>>\n")

	index := uint32(0)
	if ctx.PopWide() {
		index = uint32(utils.BigEndian2Little4U2(ctx.Code[ctx.PC : ctx.PC + 2]))
		ctx.PC += 2
	} else {
		index = uint32(ctx.Code[ctx.PC])
		ctx.PC += 1
	}
	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.JDN) && reflect.TypeOf(value) != reflect.TypeOf(types.Jdouble(0)) &&
		reflect.TypeOf(value) != reflect.TypeOf(types.JDO) && reflect.TypeOf(value) != reflect.TypeOf(types.JDU){
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentAborigines.SetAborigines(uint32(index), value)
	return nil
}

func (s I_dstore)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(2))
	f.PushFrame(types.Jdouble(33))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x0, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		将一个 double 类型数据保存到局部变量表中
======================================================================================
						||		dstore
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
		结构				||		dstore = 57(0x39)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
						||		index 是一个无符号 byte 型整数，它和 index+1 共同一个指向当前栈帧(§ 2.6)局部变量表的索引值，
		描述				||		而在操作数栈栈顶的 value 必须是 double 类型 的数据，这个数据将从操作数栈出栈，
						||		并且经过数值集合转换(§2.8.3)后得到值 value’，然后保存到 index 和 index+1 所指向的局部变量表位置 中。
						||		
======================================================================================
						||		
						||
						||
	   运行时异常			||
						||		
						||		
						||		
======================================================================================
						||
		注意				||		dstore 指令可以与 wide 指令联合使用，以实现使用 2 字节宽度的无符号整 数作为索引来访问局部变量表。
						||
						||
						||
======================================================================================
 */