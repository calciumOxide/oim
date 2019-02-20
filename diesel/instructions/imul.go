package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_imul struct {
}

func init()  {
	INSTRUCTION_MAP[0x68] = &I_imul{}
}

func (s I_imul)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "imul exce >>>>>>>>>\n")

	value2, _ := ctx.CurrentFrame.PopFrame()
	value1, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jint(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jfloat(value1.(types.Jint) * value2.(types.Jint)))
	return nil
}

func (s I_imul)Test() *runtime.Context {
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
		操作				||		int 类型数据乘法
======================================================================================
						||		imul
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
		结构				||		imul = 104(0x68)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||
						||		value1 和 value2 都必须为 int 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		接着将这两个数值相乘(value1×value2)，结果压入到操作数栈中。
		描述				||		运算的结果使用低位在高地址(Low-Order Bites)的顺序、按照二进制补 码形式存储在 32 位空间中，其数据类型为 int。
						||		如果发生了上限溢出，那结 果的符号可能与真正数学运算结果的符号相反。
						||		尽管可能发生上限溢出，但是 imul 指令的执行过程中不会抛出任何运行时异 常。
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