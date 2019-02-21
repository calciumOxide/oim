package instructions

import (
	"../runtime"
	"../../utils"
	"../oil/types"
	"reflect"
	"../variator"
)

type I_f2d struct {
}

func init()  {
	INSTRUCTION_MAP[0x8d] = &I_f2d{}
}

func (s I_f2d)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "f2d exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) == reflect.TypeOf(types.JFN) {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}
	if reflect.TypeOf(value) == reflect.TypeOf(types.JFU) {
		ctx.CurrentFrame.PushFrame(types.JDU)
		return nil
	}
	if reflect.TypeOf(value) == reflect.TypeOf(types.JFO) {
		ctx.CurrentFrame.PushFrame(types.JDO)
		return nil
	}
	if reflect.TypeOf(value) != reflect.TypeOf(types.Jfloat(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jdouble(value.(types.Jfloat)))
	return nil
}

func (s I_f2d)Test() *runtime.Context {
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
		操作				||		将 float 类型数据转换为 double 类型
======================================================================================
						||		f2d
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
		结构				||		f2d = 141(0x8d)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
						||		在操作数栈栈顶的值 value 必须为 float 类型的数据，指令执行时，value 从操作数栈中出栈，
						||		并且经过数值集合转换(§2.8.3)后得到值 value’， value’
						||		再通过 IEEE 754 的向最接近数舍入模式(§2.8.1)转换为 double类型值 result。然后 result 被压入到操作数栈中。
						||
						||		如果 d2f 指令运行在 FP-strict(§2.8.2)模式下，那指令执行过程就是 一种宽化类型转换(Widening Primitive Conversion，JLS3 §5.1.2)。
		描述				||		因为所有单精度浮点数集合(§2.3.2)都可以在双精度浮点数集合(§ 2.3.2)中找到精确对应的数值，因此这种转换是精确的。
						||
						||		如果 d2f 指令运行在非 FP-strict 模式下，那转换结果就可能会从双精度扩 展指数集合(§2.3.2)中选取，
						||		并不需要舍入为双精度浮点数集合中最接近 的可表示值。不过，如果操作数 value 是单精度扩展指数集合中的数值，那 把结果舍入为双精度浮点数集合则是必须的。
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
		注意				||
						||
						||
						||
======================================================================================
 */