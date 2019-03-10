package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_irem struct {
}

func init()  {
	INSTRUCTION_MAP[0x70] = &I_irem{}
}

func (s I_irem)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "irem exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jint(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	if value2.(types.Jint) == types.Jint(0) {
		except, _ := variator.AllocExcept(variator.ArithmeticException)
		ctx.Throw(except)
		return nil
	}

	if value1.(types.Jint) == types.Jint(-0x7fffffff) && value2.(types.Jint) == types.Jint(-1) {
		ctx.CurrentFrame.PushFrame(types.Jint(0))
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jint(value1.(types.Jint) | value2.(types.Jint)))
	return nil
}

func (s I_irem)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		int 类型数据求余
======================================================================================
						||		irem
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
		结构				||		irem = 112(0x70)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||
						||		value1 和 value2 都必须为 int 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		根据 value1-(value1÷value2)×value2 计算出结果，然后把运算结果入栈回操作数栈中。
		描述				||		irem 指令的运算结果就是保证(a÷b)×b +(a%b)=a 能够成立，
						||		唯一的特 殊情况是当被除数是 int 类型绝对值最大的负数，并且除数为-1 的时候(这 时候余数值为 0)。
						||		irem 运算指令执行时会遵循当被除数为负数时余数才能是 负数，当被除数为正数时余数才能是正数的规则。
						||		另外，irem 运算结果的绝 对值永远小于除数的绝对值。
						||
						||
						||
						||
======================================================================================
						||		
						||
						||
	   运行时异常			||		如果除数为 0，irem 指令将会抛出一个 ArithmeticException 异常。
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