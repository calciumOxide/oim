package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_lushr struct {
}

func init()  {
	INSTRUCTION_MAP[0x7d] = &I_lushr{}
}

func (s I_lushr)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lushr exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jlong(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jlong(uint32(value1.(types.Jlong)) >> uint32(value2.(types.Jlong) & types.Jlong(0x3f))))
	return nil
}

func (s I_lushr)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		long 数值逻辑右移运算
======================================================================================
						||		lushr
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
		结构				||		lushr = 125(0x7d)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		„，result
======================================================================================
						||
		描述				||		value1 和 value2 都必须为 long 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，然后将 value1 右移 s 位，s 是 value2 低 6 位所表示的
         值，计算后把运算结果入栈回操作数栈中。

						||
======================================================================================
						||		
	   运行时异常			||
						||
======================================================================================
						||
		注意				||		假设 value1 是正数并且 s 为 value2 与 0x3f 算术与运算后的结果，那 lushr 指令的运算结果与 value1 >> s 的结果是一致的;假设 value1 是负数，那
lushr指令的运算结果与表达式(value1 >> s)+(2L << ~s)一致。 附加的(2L << ~s)操作用于取消符号位的移动。位移的距离实际上被限制 在 0 到 63 之间，
						||
						||
======================================================================================
 */