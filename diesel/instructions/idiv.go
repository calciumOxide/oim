package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_idiv struct {
}

func init()  {
	INSTRUCTION_MAP[0x6c] = &I_idiv{}
}

func (s I_idiv)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "idiv exce >>>>>>>>>\n")

	value2, _ := ctx.CurrentFrame.PopFrame()
	value1, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jint(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	if value2.(types.Jint) == 0 {
		except, _ := variator.AllocExcept(variator.ArithmeticException)
		ctx.Throw(except)
		return nil
	}
 	if value1.(types.Jint) == -2147483648 && value2.(types.Jint) == -1 {
		ctx.CurrentFrame.PushFrame(value1)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jint(value1.(types.Jint) / value2.(types.Jint)))
	return nil
}

func (s I_idiv)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(-2147483648))
	f.PushFrame(types.Jint(3))
	f.PushFrame(types.Jint(2))
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
		操作				||		int 类型数据除法
======================================================================================
						||		idiv
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
		结构				||		idiv = 108(0x6c)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
						||		value1 和 value2 都必须为 int 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		并且将这两个数值相除(value1÷value2)，结果转换为 int 类型值 result，最后 result 被压入到操作数栈中。
						||
		描述				||		int 类型的除法结果都是向零舍入的，这意味着 n÷d 的商 q 会在满足 |d|×|q|≤|n|的前提下取尽可能大的整数值。
						||		另外，当|n|≥|d|并且 n 和 d 符号相同时，q 的符号为正。
						||		而当|n|≥|d|并且 n 和 d 的符号相反时，q 的符 号为负。
						||
						||		有一种特殊情况不适合上面的规则:如果被除数是 int 类型中绝对值最大的负 数，除数为-1。
						||		那运算时将会发生溢出，运算结果就等于被除数本身。尽管 这里发生了溢出，但是依然不会有异常抛出。
						||
======================================================================================
						||		
						||
						||
	   运行时异常			||		如果除数为零，idiv 指令将抛出 ArithmeticException 异常。
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