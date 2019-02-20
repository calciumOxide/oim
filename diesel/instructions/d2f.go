package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_d2f struct {
}

func init()  {
	INSTRUCTION_MAP[0x90] = &I_d2f{}
}

func (s I_d2f)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "d2f exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) == reflect.TypeOf(types.JDN) {
		ctx.CurrentFrame.PushFrame(types.JFN)
		return nil
	}
	if reflect.TypeOf(value) == reflect.TypeOf(types.JDU) {
		ctx.CurrentFrame.PushFrame(types.JFU)
		return nil
	}
	if reflect.TypeOf(value) == reflect.TypeOf(types.JDO) {
		ctx.CurrentFrame.PushFrame(types.JFO)
		return nil
	}
	if reflect.TypeOf(value) != reflect.TypeOf(types.Jdouble(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jfloat(value.(types.Jdouble)))
	return nil
}

func (s I_d2f)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jdouble(9.123456789012345))
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
		操作				||		将 double 类型数据转换为 float 类型
======================================================================================
						||		d2f
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
		结构				||		d2f = 144(0x90)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
						||		在操作数栈栈顶的值 value 必须为 double 类型的数据，指令执行时，value 从操作数栈中出栈，并且经过数值集合转换(§2.8.3)后得到值 value’，
						||		value’再通过 IEEE 754 的向最接近数舍入模式(§2.8.1)转换为 float类型值 result。然后 result 被压入到操作数栈中。
						||
						||		如果 d2f 指令运行在 FP-strict(§2.8.2)模式下，那转换的结果永远是 转换为单精度浮点值集合中与原值最接近的可表示值。
		描述				||		如果 d2f 指令运行在非 FP-strict 模式下，那转换结果可能会从单精度扩展 指数集合(§2.3.2)中选取，也就是说并非一定会转换为单精度浮点值集合 中与原值最接近的可表示值的。
						||
						||		当有限值 value’太小以至于无法使用 float 类型数据来表示时，将会被转 换为与原值符号相同的零值。
						||		同样，当有限值 value’太大以至于无法使用 float 类型数据来表示时，将会被转换为与原值符号相同的无穷大。
						||		double 类型的 NaN 值永远转换为 float 类型的 NaN 值。
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
						||		d2f 指令执行了窄化类型转换(Narrowing Primitive Conversion，JLS3 §5.1.3)，它可能会导致 value’的数值大小和精度发生丢失。
		注意				||
						||
						||
						||
======================================================================================
 */