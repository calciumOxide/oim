package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_lsub struct {
}

func init()  {
	INSTRUCTION_MAP[0x65] = &I_lsub{}
}

func (s I_lsub)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lsub exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jlong(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jlong(value1.(types.Jlong) - value2.(types.Jlong)))
	return nil
}

func (s I_lsub)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jlong(9))
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
		操作				||		long 类型数据相减
======================================================================================
						||		lsub
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
		结构				||		lsub = 101(0x65)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		„，result
======================================================================================
						||
		描述				||		value1 和 value2 都必须为 long 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，将这两个数值相减(result = value1’− value2’)，
结果转换为 long 类型值 result，最后 result 被压入到操作数栈中。
对于 long 类型数据的减法来说，a-b 与 a+(-b)的结果永远是一致的，0 减去某个 long 值相当于对这个 long 值进行取负运算。
运算的结果使用低位在高地址(Low-Order Bites)的顺序、按照二进制补 码形式存储在 64 位空间中，其数据类型为 long。如果发生了上限溢出，那结 果的符号可能与真正数学运算结果的符号相反。
尽管可能发生上限溢出，但是 lsub 指令的执行过程中不会抛出任何运行时异 常。
						||
						||
======================================================================================
						||		
	   运行时异常			||
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