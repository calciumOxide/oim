package instructions

import (
	"../runtime"
	"../../utils"
	"../variator"
	"../oil/types"
	"reflect"
)

type I_iaload struct {
}

func init()  {
	INSTRUCTION_MAP[0x2e] = &I_iaload{}
}

func (s I_iaload)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "iaload exce >>>>>>>>>\n")

	index, _ := ctx.CurrentFrame.PopFrame()
	array, _ := ctx.CurrentFrame.PopFrame()
	if array == nil {
		except, _ := variator.AllocExcept(variator.NullPointerException)
		ctx.Throw(except)
		return nil
	}
	if array.(*types.Jarray).ElementJype == nil || array.(*types.Jarray).ElementJype != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	bytes := array.(*types.Jarray).Reference.([]interface{})
	if len(bytes) <= int(index.(types.Jint)) {
		except, _ := variator.AllocExcept(variator.ArrayIndexOutOfBoundsException)
		ctx.Throw(except)
		return nil
	}
	value := bytes[index.(types.Jint)]
	ctx.CurrentFrame.PushFrame(value)
	return nil
}

func (s I_iaload)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		ElementJype: reflect.TypeOf(types.Jint(0)),
		Reference: []interface{}{1, 2, 3, 4},
	})
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
		操作				||		从数组中加载一个 int 类型数据到操作数栈
======================================================================================
						||		iaload
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
		结构				||		iaload = 46(0x2e)
======================================================================================
						||		...，arrayref，index →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||		
						||		arrayref 必须是一个 reference 类型的数据，它指向一个组件类型为 int 的数组，index 必须为 int 类型。
		描述				||		指令执行后，arrayref 和 index 同时从操作数栈出栈，
						||		ndex 作为索引定位到数组中的 int 类型值将压入到操作数 栈中。
						||		
======================================================================================
						||		
						||		如果 arrayref 为 null，iaload 指令将抛出 NullPointerException 异 常
						||
	   运行时异常			||		另外，如果 index 不在 arrayref 所代表的数组上下界范围中，iaload 指 令将抛出 ArrayIndexOutOfBoundsException 异常。
						||		
						||		
						||		
======================================================================================
						||
						||		//iaload 指令可以用来从数组中读取 byte 或者 boolean 的数据，在 Oracle 的虚拟机实现中，
						||		//布尔类型的数组(T_BOOLEAN 类型的数组可参见§2.2 和本章中对 newarray 指令的介绍)被实现为 8 位宽度的数值，
		注意				||		//而其他的虚拟机 实现很可能使用其他方式实现 boolean 数组，那其他虚拟机实现的 iaload 就必须能正确访问相应实现的数组。
						||
						||
						||
======================================================================================
 */