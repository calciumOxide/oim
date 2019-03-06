package instructions

import (
	"../runtime"
	"../../utils"
	"../variator"
	"../../types"
	"reflect"
)

type I_laload struct {
}

func init()  {
	INSTRUCTION_MAP[0x2f] = &I_laload{}
}

func (s I_laload)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "laload exce >>>>>>>>>\n")

	index, _ := ctx.CurrentFrame.PopFrame()
	array, _ := ctx.CurrentFrame.PopFrame()
	if array == nil {
		except, _ := variator.AllocExcept(variator.NullPointerException)
		ctx.Throw(except)
		return nil
	}
	if array.(*types.Jarray).ElementJype == nil || array.(*types.Jarray).ElementJype != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	bytes := array.(*types.Jarray).Reference.([]interface{})
	if len(bytes) <= int(index.(types.Jlong)) {
		except, _ := variator.AllocExcept(variator.ArrayIndexOutOfBoundsException)
		ctx.Throw(except)
		return nil
	}
	value := bytes[index.(types.Jlong)]
	ctx.CurrentFrame.PushFrame(value)
	return nil
}

func (s I_laload)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		ElementJype: reflect.TypeOf(types.Jlong(0)),
		Reference: []interface{}{1, 2, 3, 4},
	})
	f.PushFrame(types.Jlong(2))
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
		操作				||		从数组中加载一个 long 类型数据到操作数栈
======================================================================================
						||		laload
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
		结构				||		laload = 47(0x2f)
======================================================================================
						||		...，arrayref，index →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||		
		描述				||		arrayref 必须是一个 reference 类型的数据，它指向一个组件类型为 long 的数组，index 必须为 int 类型。指令执行后，arrayref 和 index 同时从 操作数栈出栈，index 作为索引定位到数组中的 long 类型值将压入到操作数
栈中。
======================================================================================
						||		
						||		如果 arrayref 为 null，laload 指令将抛出 NullPointerException 异 常
						||
	   运行时异常			||		另外，如果 index 不在 arrayref 所代表的数组上下界范围中，laload 指 令将抛出 ArrayIndexOutOfBoundsException 异常。
						||		
						||		
						||		
======================================================================================
						||
						||
		注意				||
						||
						||
======================================================================================
 */