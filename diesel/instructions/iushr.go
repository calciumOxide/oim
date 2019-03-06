package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_iushr struct {
}

func init()  {
	INSTRUCTION_MAP[0x7c] = &I_iushr{}
}

func (s I_iushr)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "iushr exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jint(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jint(uint32(value1.(types.Jint)) >> uint32(value2.(types.Jint) & types.Jint(0x1f))))
	return nil
}

func (s I_iushr)Test() *runtime.Context {
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
		操作				||		int 数值逻辑右移运算
======================================================================================
						||		iushr
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
		结构				||		iushr = 124(0x7c)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		„，result
======================================================================================
						||
						||		value1 和 value2 都必须为 int 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，然后将 value1 右移 s 位，
		描述				||		s 是 value2 低 5 位所表示的值，计算后把运算结果入栈回操作数栈中。
						||
======================================================================================
						||		
	   运行时异常			||
						||
======================================================================================
						||
						||
						||		假设 value1 是正数并且 s 为 value2 与 0x1f 算术与运算后的结果，那 iushr 指令的运算结果与 value1 >> s 的结果是一致的;
		注意				||		假设 value1 是负数，那 iushr指令的运算结果与表达式(value1 >> s)+(2 << ~s)一致。附 加的(2 << ~s)操作用于取消符号位的移动。位移的距离实际上被限制在 0到 31 之间，
						||
						||
						||
======================================================================================
 */