package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_lor struct {
}

func init()  {
	INSTRUCTION_MAP[0x81] = &I_lor{}
}

func (s I_lor)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lor exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jlong(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jlong(value1.(types.Jlong) | value2.(types.Jlong)))
	return nil
}

func (s I_lor)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		long 类型数值的布尔或运算
======================================================================================
						||		lor
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
		结构				||		lor = 129(0x81)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||
						||		value1、value2 必须为 long 类型数据，指令执行时，它们从操作数栈中出 栈，接着对这2个数进行按位或(Bitwise Inclusive OR)运算，long
类型的运算结果 result 被压入到操作数栈中。
		描述				||
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
						||
						||
		注意				||
						||
						||
						||
======================================================================================
 */