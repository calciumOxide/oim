package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_f2l struct {
}

func init()  {
	INSTRUCTION_MAP[0x8c] = &I_f2l{}
}

func (s I_f2l)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "f2l exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) == reflect.TypeOf(types.JDN) {
		ctx.CurrentFrame.PushFrame(types.Jlong(0))
		return nil
	}
	if reflect.TypeOf(value) == reflect.TypeOf(types.JDO) {
		ctx.CurrentFrame.PushFrame(types.Jlong(0x7fffffffffffffff))
		return nil
	}
	if reflect.TypeOf(value) == reflect.TypeOf(types.JDU) {
		ctx.CurrentFrame.PushFrame(types.Jlong(-0x7fffffffffffffff))
		return nil
	}
	if reflect.TypeOf(value) != reflect.TypeOf(types.Jfloat(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jlong(value.(types.Jfloat)))
	return nil
}

func (s I_f2l)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jfloat(9.123456789012345))
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
		操作				||		将 float 类型数据转换为 long 类型
======================================================================================
						||		f2l
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
		结构				||		f2l = 140(0x8c)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
						||		在操作数栈栈顶的值 value 必须为 float 类型的数据，指令执行时，value 从操作数栈中出栈，
						||		并且经过数值集合转换(§2.8.3)后得到值 value’，value’
						||		再转换为 long 类型值 result。然后 result 被压入到操作数栈中。
						||
						||		 如果 value’是 NaN 值，那 result 的转换结果为 long 类型的零值。
		描述				||		 另外，如果 value’不是无穷大，那将会使用 IEEE 754 标准中的向零舍入模式(§2.8.1)转换成整数值 V，如果这个整数 V 在 long 类型的可表示范围之内，那么 result 的转换结果就是这个整数 V。
						||		 另外，如果 value’太小(绝对值很大的负数或者负无穷大)以至于超过了 long 类型可表示的下限，那将转换为 long 类型中最小的可表示数。同 样地，如果 value’太大(很大的正数或者无穷大)以至于超过了 long 类型可表示的上限，那将转换为 long 类型中最大的可表示数。
						||
						||
						||
						||
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
						||		f2l 指令执行了窄化类型转换(Narrowing Primitive Conversion，JLS3 §5.1.3)，它可能会导致 value’的数值大小和精度发生丢失。
		注意				||
						||
						||
						||
======================================================================================
 */