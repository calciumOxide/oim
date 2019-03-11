package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_lshl struct {
}

func init()  {
	INSTRUCTION_MAP[0x79] = &I_lshl{}
}

func (s I_lshl)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lshl exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jlong(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jlong(value1.(types.Jlong) << uint32(value2.(types.Jlong) & types.Jlong(0x3f))))
	return nil
}

func (s I_lshl)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		long 数值左移运算
======================================================================================
						||		lshl
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
		结构				||		lshl = 121(0x79)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		„，result
======================================================================================
						||
		描述				||		value1 和 value2 都必须为 long 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，然后将 value1 左移 s 位，s 是 value2 低 6 位所表示的
         值，计算后把运算结果入栈回操作数栈中。
						||
======================================================================================
						||		
	   运行时异常			||
						||
======================================================================================
						||
		注意				||这个操作(即使是出现了溢出的情况下)等同于把 value1 乘以 2 的 s 次方， 位移的距离实际上被限制在 0 到 63 之间，相当于指令执行时会把 value2 与
0x3f 做一遍算术与操作。
						||
						||
						||
======================================================================================
 */