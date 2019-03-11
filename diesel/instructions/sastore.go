package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_sastore struct {
}

func init()  {
	INSTRUCTION_MAP[0x56] = &I_sastore{}
}

func (s I_sastore)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "sastore exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()
	index, _ := ctx.CurrentFrame.PopFrame()
	array, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jint(0)) || reflect.TypeOf(index) != reflect.TypeOf(types.Jint(0)) ||
		reflect.TypeOf(array) != reflect.TypeOf(&types.Jarray{}) {
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

	array.(*types.Jarray).Reference.([]types.Jbyte)[index.(types.Jint)] = types.Jbyte(value.(types.Jint))
	return nil
}

func (s I_sastore)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		从操作数栈读取一个 short 类型数据存入到数组中
======================================================================================
						||		sastore
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
		结构				||		sastore = 86(0x56)
======================================================================================
						||		...，arrayref，index，value →
	   操作数栈			||------------------------------------------------------------
						||		„，
======================================================================================
						||
		描述				||		arrayref 必须是一个 reference 类型的数据，它指向一个组件类型为 short 的数组，
index 和 value 都必须为 int 类型。指令执行后，arrayref、index 和 value 同时从操作数栈出栈，
value 将被转换为 short 类型，然 后存储到 index 作为索引定位到数组元素中。
						||
======================================================================================
						||		
	   运行时异常			||
						||如果 arrayref 为 null，sastore 指令将抛出 NullPointerException 异常
						||
						||另外，如果 index 不在 arrayref 所代表的数组上下界范围中，sastore 指 令将抛出 ArrayIndexOutOfBoundsException 异常。
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