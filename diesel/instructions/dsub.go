package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_dsub struct {
}

func init()  {
	INSTRUCTION_MAP[0x67] = &I_dsub{}
}

func (s I_dsub)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "dsub exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value2) == reflect.TypeOf(types.JDN) || reflect.TypeOf(value1) == reflect.TypeOf(types.JDN) {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if (reflect.TypeOf(value1) == reflect.TypeOf(types.JDO) || reflect.TypeOf(value1) == reflect.TypeOf(types.JDU)) &&
		(reflect.TypeOf(value2) == reflect.TypeOf(types.JDU) && reflect.TypeOf(value2) == reflect.TypeOf(types.JDO)) {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if reflect.TypeOf(value1) == reflect.TypeOf(types.JDO) || reflect.TypeOf(value1) == reflect.TypeOf(types.JDU) {
		ctx.CurrentFrame.PushFrame(value1)
		return nil
	}

	if reflect.TypeOf(value2) == reflect.TypeOf(types.JDO) {
		ctx.CurrentFrame.PushFrame(types.JDU)
		return nil
	}

	if reflect.TypeOf(value2) == reflect.TypeOf(types.JDU) {
		ctx.CurrentFrame.PushFrame(types.JDO)
		return nil
	}

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jdouble(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jdouble(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jdouble(value1.(types.Jdouble) - value2.(types.Jdouble)))
	return nil
}

func (s I_dsub)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jdouble(8.123456789012345))
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
		操作				||		double 类型数据相减
======================================================================================
						||		dsub
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
		结构				||		dsub = 103(0x67)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
						||		value1 和 value2 都必须为 double 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		并且经过数值集合转换(§2.8.3)后得到值 value1’和 value2’，
						||		接着将这两个数值相减(result = value1’− value2’)，结果转换为 double 类型值 result，最后 result 被压入到操 作数栈中。
						||
		描述				||		对于一般 double 类型数据的减法来说，a-b 与 a+(-b)的结果永远是一 致的，但是对于 dsub 指令来说，与零相减的符号则会相反，因为如果 x 是+ 0.0 的话，那 0.0-x 等于+0.0，但-x 等于-0.0。
						||
						||		Java 虚拟机必须支持 IEEE 754 中定义的逐级下溢(Gradual Underflow)， 尽管指令执行期间，上溢、下溢以及精度丢失等情况都有可能发生，但 dsub 指令永远不会抛出任何运行时异常。
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