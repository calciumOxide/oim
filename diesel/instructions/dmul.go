package instructions

import (
	"../../utils"
	"../oil/types"
	"../runtime"
	"../variator"
	"reflect"
)

type I_dmul struct {
}

func init() {
	INSTRUCTION_MAP[0x6b] = &I_dmul{}
}

func (s I_dmul) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "dmul exce >>>>>>>>>\n")

	value2, _ := ctx.CurrentFrame.PopFrame()
	value1, _ := ctx.CurrentFrame.PopFrame()

	if value2 == types.JDN || value1 == types.JDN {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if (reflect.TypeOf(value1) == reflect.TypeOf(types.JDO) && reflect.TypeOf(value2) == reflect.TypeOf(types.JDU)) ||
		reflect.TypeOf(value2) == reflect.TypeOf(types.JDO) && reflect.TypeOf(value1) == reflect.TypeOf(types.JDU) {
		ctx.CurrentFrame.PushFrame(types.JDO)
		return nil
	}

	if reflect.TypeOf(value1) == reflect.TypeOf(types.JDO) || reflect.TypeOf(value1) == reflect.TypeOf(types.JDU) {
		if reflect.TypeOf(value2) == reflect.TypeOf(types.Jdouble(0)) {
			if value2.(types.Jdouble) == 0 {
				ctx.CurrentFrame.PushFrame(types.JDN)
				return nil
			}
			ctx.CurrentFrame.PushFrame(value1)
			return nil
		}
	}

	if reflect.TypeOf(value2) == reflect.TypeOf(types.JDO) || reflect.TypeOf(value2) == reflect.TypeOf(types.JDU) {
		if reflect.TypeOf(value2) == reflect.TypeOf(types.Jdouble(0)) {
			if value2.(types.Jdouble) == 0 {
				ctx.CurrentFrame.PushFrame(types.JDN)
				return nil
			}
			ctx.CurrentFrame.PushFrame(value2)
			return nil
		}
	}

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jdouble(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jdouble(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jdouble(value1.(types.Jdouble) * value2.(types.Jdouble)))
	return nil
}

func (s I_dmul) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jdouble(9.123456789012345))
	f.PushFrame(types.Jdouble(9.123456789012345))
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code:              []byte{0x0},
		CurrentFrame:      f,
		CurrentAborigines: a,
	}
}

/**
======================================================================================
		操作				||		double 类型数据乘法
======================================================================================
						||		dmul
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
		结构				||		dmul = 107(0x6b)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||
						||
						||		value1 和 value2 都必须为 double 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		并且经过数值集合转换(§2.8.3)后得到值 value1’和 value2’，
						||		接着将这两个数值相乘(value1’×value2’)，结 果转换为 double 类型值 result，最后 result 被压入到操作数栈中。
						||
						||		dmul 指令的运算结果取决于 IEEE 规范中规定的运算规则:
		描述				||			 如果 value1’和 value2’中有任意一个值为 NaN，那运算结果即为 NaN。
						||			 如果 value1’和 value2’两者都不为 NaN，那当两者符号相同时，运算结果为正，反着，当两者符号不同时，运算结果为负。
						||			 无穷大与零值相乘，运算结果为 NaN。
						||			 一个无穷大的数与一个有限的数相乘，结果为无穷大，无穷大的符号由第2 点规则确定。
						||			 对于上述情况之外的场景，即任意一个操作数都不是无穷大或者 NaN，就按算术求积，
						||				并以 IEEE 754 规范的最接近数舍入模式得到运算结果，如 果运算结果的绝对值太大以至于无法使用 double 类型来表示，
						||				换句话说 就是出现了上限溢出，那结果将会使用具有适当符号的无穷大来代替。
						||				如 果运算结果的绝对值太小以至于无法使用 double 类型来表示，换句话说 就是出现了下限溢出，那结果将会使用具有适当符号的零值来代替。
						||
						||		Java 虚拟机必须支持 IEEE 754 中定义的逐级下溢(Gradual Underflow)， 尽管指令执行期间，上溢、下溢以及精度丢失等情况都有可能发生，但 dmul 指令永远不会抛出任何运行时异常。
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
