package instructions

import (
	"../runtime"
	"../../utils"
	"../variator"
	"../../types"
	"reflect"
)

type I_lcmp struct {
}

func init()  {
	INSTRUCTION_MAP[0x94] = &I_lcmp{}
}

func (s I_lcmp)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lcmp exce >>>>>>>>>\n")

	value2, _ := ctx.CurrentFrame.PopFrame()
	value1, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jlong(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	if value1.(types.Jlong) > value2.(types.Jlong) {
		ctx.CurrentFrame.PushFrame(types.Jint(1))
	} else if value1.(types.Jlong) > value2.(types.Jlong) {
		ctx.CurrentFrame.PushFrame(types.Jint(-1))
	} else {
		ctx.CurrentFrame.PushFrame(types.Jint(0))
	}

	return nil
}

func (s I_lcmp)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		ElementJype: reflect.TypeOf(types.Jlong(0)),
		Reference: []interface{}{1, 2, 3, 4},
	})
	f.PushFrame(types.Jlong(2))
	f.PushFrame(types.Jlong(6))

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
		操作				||		比较 2 个 long 类型数据的大小
======================================================================================
						||		lcmp
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
		结构				||		lcmp = 148(0x94)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
		描述				||		value1 和 value2 都必须为 long 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，使用一个 int 数值作为比较结果:如果 value1 大于
value2，结果为 1;如果 value1 等于 value2，结果为 0;如果 value1 小于 value2，结果为-1，最后比较结果被压入到操作数栈中。
======================================================================================
						||		
						||
	   运行时异常			||
						||		
						||		
						||		
======================================================================================
						||
						||		//lcmp 指令可以用来从数组中读取 byte 或者 boolean 的数据，在 Oracle 的虚拟机实现中，
						||		//布尔类型的数组(T_BOOLEAN 类型的数组可参见§2.2 和本章中对 newarray 指令的介绍)被实现为 8 位宽度的数值，
		注意				||		//而其他的虚拟机 实现很可能使用其他方式实现 boolean 数组，那其他虚拟机实现的 lcmp 就必须能正确访问相应实现的数组。
						||
						||
						||
======================================================================================
 */