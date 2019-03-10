package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_isub struct {
}

func init()  {
	INSTRUCTION_MAP[0x64] = &I_isub{}
}

func (s I_isub)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "isub exce >>>>>>>>>\n")

	value1, _ := ctx.CurrentFrame.PopFrame()
	value2, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jint(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jint(value1.(types.Jint) - value2.(types.Jint)))
	return nil
}

func (s I_isub)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		int 类型数据相减
======================================================================================
						||		isub
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
		结构				||		isub = 100(0x64)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		„，result
======================================================================================
						||
						||		value1 和 value2 都必须为 int 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
		描述				||		将这两个数值相减(result = value1’− value2’)，结果转换为 int 类型值 result，最后 result 被压入到操作数栈中。
						||		对于 int 类型数据的减法来说，a-b 与 a+(-b)的结果永远是一致的，0 减去某个 int 值相当于对这个 int 值进行取负运算。
						||		运算的结果使用低位在高地址(Low-Order Bites)的顺序、按照二进制补 码形式存储在 32 位空间中，其数据类型为 int。
						||		如果发生了上限溢出，那结 果的符号可能与真正数学运算结果的符号相反。
						||		尽管可能发生上限溢出，但是 isub 指令的执行过程中不会抛出任何运行时异 常。
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