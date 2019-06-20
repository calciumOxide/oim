package instructions

import (
	"../../utils"
	"../oil/types"
	"../runtime"
	"../variator"
	"reflect"
)

type I_frem struct {
}

func init() {
	INSTRUCTION_MAP[0x72] = &I_frem{}
}

func (s I_frem) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "frem exce >>>>>>>>>\n")

	value2, _ := ctx.CurrentFrame.PopFrame()
	value1, _ := ctx.CurrentFrame.PopFrame()
	if (reflect.TypeOf(value1) != reflect.TypeOf(types.Jfloat(0)) && reflect.TypeOf(value1) != reflect.TypeOf(types.JDN) &&
		reflect.TypeOf(value1) != reflect.TypeOf(types.JDO) && reflect.TypeOf(value1) != reflect.TypeOf(types.JDU)) ||
		(reflect.TypeOf(value2) != reflect.TypeOf(types.Jfloat(0)) && reflect.TypeOf(value2) != reflect.TypeOf(types.JDN) &&
			reflect.TypeOf(value2) != reflect.TypeOf(types.JDO) && reflect.TypeOf(value2) != reflect.TypeOf(types.JDU)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	if value1 == types.JDN || value2 == types.JDN {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if value1 == types.JDO || value1 == types.JDU {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if reflect.TypeOf(value2) == reflect.TypeOf(types.JDO) || reflect.TypeOf(value2) == reflect.TypeOf(types.JDU) {
		ctx.CurrentFrame.PushFrame(value1)
		return nil
	}

	if reflect.TypeOf(value2) == reflect.TypeOf(types.Jfloat(0)) && value2.(types.Jfloat) == 0.0 {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if reflect.TypeOf(value1) == reflect.TypeOf(types.Jfloat(0)) && value1.(types.Jfloat) == 0.0 {
		ctx.CurrentFrame.PushFrame(value1)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jfloat(value1.(types.Jfloat) - value2.(types.Jfloat)*types.Jfloat(int64(value1.(types.Jfloat)/value2.(types.Jfloat)))))
	return nil
}

func (s I_frem) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jfloat(9.99))
	f.PushFrame(types.Jfloat(2.2))
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
		操作				||		float 类型数据求余
======================================================================================
						||		frem
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
		结构				||		frem = 114(0x72)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||
						||
						||		value1 和 value2 都必须为 float 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		并且经过数值集合转换(§2.8.3)后得到值 value1’和 value2’，
						||		接着将这两个数值求余，结果转换为 float 类型值 result，最后 result 被压入到操作数栈中。
						||
						||		frem 指令的运算结果与 IEEE 754 中定义的 remainder 操作并不相同，
						||		IEEE 754 中的 remainder 操作使用舍入除法(Rounding Division)而不是去 尾触发(Truncating Division)来获得求余结果，
						||		因此这种运算与通常对 整数的求余方式并不一致。
						||		Java 虚拟机中定义的 frem 则是与虚拟机中整数 求余指令(irem 和 lrem)保持了一致的行为，
						||		这可以与 C 语言中的 fmod 函 数互相比较。
						||
						||
						||
						||
		描述				||		frem 指令的运算结果通过以下规则获得:
						||			 如果 value1’和 value2’中有任意一个值为 NaN，那运算结果即为 NaN。
						||			 如果 value1’和 value2’两者都不为 NaN，那运算结果的符号被除数的符号一致。
						||			 如果被除数是无穷大，或者除数为零，那运算结果为 NaN。
						||			 如果被除数是有限值，而除数是无穷大，那运算结果等于被除数。
						||			 如果被除数为零，而除数是有限值，那运算结果等于被除数。
						||			 对于上述情况之外的场景，即任意一个操作数都不是无穷大、零以及 NaN，就以 value1’为被除数、value2’为除数使用浮点算术规则求余:
						||				result = value1’−(value2’∗ q)，这里的 q 是一个整数，其符号与 value1’÷value2’的符号相同，大小与他们的商相同。
						||
 						||		尽管除数为零的情况可能发生，但是 frem 指令永远不会抛出任何运行时异常， 上限溢出、下限溢出和进度丢失的情况也不会出现。
						||		IEEE 754 规范中定义的 remainder 操作可以使用 Math.IEEEremainder 来完成。
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
						||
		注意				||
						||
						||
						||
======================================================================================
*/
