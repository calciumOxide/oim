package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_lneg struct {
}

func init()  {
	INSTRUCTION_MAP[0x75] = &I_lneg{}
}

func (s I_lneg)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lneg exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jlong(value.(types.Jlong) * -1))
	return nil
}

func (s I_lneg)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jfloat(9.123456789012345))
	f.PushFrame(types.Jlong(9))
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
		操作				||		long 类型数据取负运算
======================================================================================
						||		lneg
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
		结构				||		lneg = 117(0x75)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
		描述				||		value 必须为 long 类型数据，指令执行时，value 从操作数栈中出栈，接着 对这个数进行算术取负运算，运算结果-value 被压入到操作数栈中。
对于 long 类型数据，取负运算等同于与零做减法运算。因为 Java 虚拟机使 用二进制补码来表示整数，而且二进制补码值的范围并不是完全对称的，long 类型中绝对值最大的负数取反的结果也依然是它本身。尽管指令执行过程中可 能发生上限溢出，但是不会抛出任何异常。
对于所有的 long 类型值 x 来说，-x 等于(~x)+1
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