package instructions

import (
	"../runtime"
	"../../utils"
	"../oil/types"
	"reflect"
	"../variator"
)

type I_dadd struct {
}

func init()  {
	INSTRUCTION_MAP[0x63] = &I_dadd{}
}

func (s I_dadd)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "dadd exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()

	if value2 == reflect.TypeOf(types.JDN) || value1 == reflect.TypeOf(types.JDN) {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if (value2 == reflect.TypeOf(types.JDU) && value1 == reflect.TypeOf(types.JDO)) ||
		(value2 == reflect.TypeOf(types.JDO) && value1 == reflect.TypeOf(types.JDU)) {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if (value2 == reflect.TypeOf(types.JDU) && value1 == reflect.TypeOf(types.JDU)) ||
		(value2 == reflect.TypeOf(types.JDO) && value1 == reflect.TypeOf(types.JDO)) {
		ctx.CurrentFrame.PushFrame(value1)
		return nil
	}

	if value2 == reflect.TypeOf(types.JDU) || value1 == reflect.TypeOf(types.JDU) {
		ctx.CurrentFrame.PushFrame(types.JDU)
		return nil
	}
	if value2 == reflect.TypeOf(types.JDO) || value1 == reflect.TypeOf(types.JDO) {
		ctx.CurrentFrame.PushFrame(types.JDO)
		return nil
	}

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jdouble(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jdouble(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jdouble(value1.(types.Jdouble) + value2.(types.Jdouble)))
	return nil
}

func (s I_dadd)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jdouble(9.123456789012345))
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
		操作				||		double 类型数据相加
======================================================================================
						||		dadd
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
		结构				||		dadd = 99(0x63)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
						||		value1 和 value2 都必须为 double 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		并且经过数值集合转换(§2.8.3)后得到值 value1’和 value2’，接着将这两个数值相加，结果转换为 double 类型值 result，最后 result 被压入到操作数栈中。
						||
						||		dadd 指令的运算结果取决于 IEEE 规范中规定的运算规则:
		描述				||			 dadd 指令的运算结果取决于 IEEE 规范中规定的运算规则:
						||			 两个不同符号的无穷大相加，结果为 NaN。
						||			 两个相同符号的无穷大相加，结果仍然为相同符号的无穷大。
						||			 一个无穷大的数与一个有限的数相加，结果为无穷大。
						||			 两个不同符号的零值相加，结果为正零。
						||			 两个相同符号的零值相加，结果仍然为相同符号的零值。
						||			 零值与一个非零有限值相加，结果等于那个非零有限值。
						||			 两个绝对值相等、符号相反的非零有限值相加，结果为正零。
						||			 对于上述情况之外的场景，即任意一个操作数都不是无穷大、零、NaN 以及两个值具有相同符号或者不同的绝对值，就按算术求和，并以最接近数舍入模式得到运算结果。
						||
						||		Java 虚拟机必须支持 IEEE 754 中定义的逐级下溢(Gradual Underflow)， 尽管指令执行期间，上溢、下溢以及精度丢失等情况都有可能发生，但 dadd 指令永远不会抛出任何运行时异常。
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