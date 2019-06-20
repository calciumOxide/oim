package instructions

import (
	"../../utils"
	"../oil/types"
	"../runtime"
	"../variator"
	"reflect"
)

type I_fdiv struct {
}

func init() {
	INSTRUCTION_MAP[0x6f] = &I_fdiv{}
}

func (s I_fdiv) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "fdiv exce >>>>>>>>>\n")

	value2, _ := ctx.CurrentFrame.PopFrame()
	value1, _ := ctx.CurrentFrame.PopFrame()

	if value1 == types.JDN || value2 == types.JDN {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if (value1 == reflect.TypeOf(types.JDU) || value1 == reflect.TypeOf(types.JDO)) &&
		(value2 == reflect.TypeOf(types.JDO) && value2 == reflect.TypeOf(types.JDU)) {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if value1 == reflect.TypeOf(types.JDU) || value1 == reflect.TypeOf(types.JDO) {
		ctx.CurrentFrame.PushFrame(value1)
		return nil
	}

	if value2 == reflect.TypeOf(types.JDU) || value2 == reflect.TypeOf(types.JDO) {
		ctx.CurrentFrame.PushFrame(types.Jfloat(0))
		return nil
	}

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jfloat(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jfloat(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	if value1.(types.Jfloat) == 0.0 && value2.(types.Jfloat) == 0.0 {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jfloat(value1.(types.Jfloat) / value2.(types.Jfloat)))
	return nil
}

func (s I_fdiv) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jfloat(9.123456789012345))
	f.PushFrame(types.Jfloat(9.123456789012345))
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
		操作				||		float 类型数据除法
======================================================================================
						||		fdiv
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
		结构				||		fdiv = 110(0x6e)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||
						||
						||		value1 和 value2 都必须为 float 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		并且经过数值集合转换(§2.8.3)后得到值 value1’和 value2’，
						||		接着将这两个数值相除(value1’÷value2’)，结 果转换为 float 类型值 result，最后 result 被压入到操作数栈中。
						||
		描述				||		fdiv 指令的运算结果取决于 IEEE 规范中规定的运算规则:
						||			 如果 value1’和 value2’中有任意一个值为 NaN，那运算结果即为 NaN。
						||			 如果 value1’和 value2’两者都不为 NaN，那当两者符号相同时，运算结果为正，反着，当两者符号不同时，运算结果为负。
						||			 两个无穷大相除，运算结果为 NaN。
						||			 一个无穷大的数与一个有限的数相除，结果为无穷大，无穷大的符号由第2 点规则确定。
						||			 一个有限的数与一个无穷大的数相除，结果为零，零值的符号由第 2 点规则确定。
						||			 零除以零结果为 NaN，零除以任意其他非零有限值结果为零，零值的符号由第 2 点规则确定。
						||			 任意非零有限值除以零结果为无穷大，无穷大的符号由第 2 点规则确定。
						||			 对于上述情况之外的场景，即任意一个操作数都不是无穷大、零以及 NaN，就按算术求商，
						||				并以IEEE 754规范的最接近数舍入模式得到运算结果， 如果运算结果的绝对值太大以至于无法使用 float 类型来表示，换句话说就是出现了上限溢出，
						||				那结果将会使用具有适当符号的无穷大来代替。
						||				如果运算结果的绝对值太小以至于无法使用 float 类型来表示，换句话 说就是出现了下限溢出，那结果将会使用具有适当符号的零值来代替。
						||
 						||		Java 虚拟机必须支持 IEEE 754 中定义的逐级下溢(Gradual Underflow)， 尽管指令执行期间，上溢、下溢以及精度丢失等情况都有可能发生，但 fdiv 指令永远不会抛出任何运行时异常。
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
