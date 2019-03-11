package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_saload struct {
}

func init()  {
	INSTRUCTION_MAP[0x35] = &I_saload{}
}

func (s I_saload)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "saload exce >>>>>>>>>\n")

	index, _ := ctx.CurrentFrame.PopFrame()
	array, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(index) != reflect.TypeOf(types.Jint(0)) || reflect.TypeOf(array) != reflect.TypeOf(types.Jarray{}) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	if array.(*types.Jarray).Reference == nil {
		except, _ := variator.AllocExcept(variator.NullPointerException)
		ctx.Throw(except)
		return nil
	}

	if index.(types.Jint) < 0 || len(array.(*types.Jarray).Reference.([]types.Jbyte)) <= int(index.(types.Jint)) {
		except, _ := variator.AllocExcept(variator.ArrayIndexOutOfBoundsException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jint(array.(*types.Jarray).Reference.([]types.Jbyte)[index.(types.Jint)]))
	return nil
}

func (s I_saload)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		从数组中加载一个 short 类型数据到操作数栈
======================================================================================
						||		saload
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
		结构				||		saload = 53(0x35)
======================================================================================
						||		...，arrayref，index  →
	   操作数栈			||------------------------------------------------------------
						||		„，value
======================================================================================
						||
		描述				||		arrayref 必须是一个 reference 类型的数据，它指向一个组件类型为 int 的数组，index 必须为 int 类型。
指令执行后，arrayref 和 index 同时从 操作数栈出栈，
index 作为索引定位到数组中的 short 类型值先被零位扩展 (Zero-Extended)为一个 int 类型数据 value，然后再将 value 压入到
操作数栈中。
						||
======================================================================================
						||		
	   运行时异常			||
						||如果 arrayref 为 null，saload 指令将抛出 NullPointerException 异 常
						||另外，如果 index 不在 arrayref 所代表的数组上下界范围中，saload 指 令将抛出 ArrayIndexOutOfBoundsException 异常。
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