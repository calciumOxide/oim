package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"../oli"
	"reflect"
	"../variator"
)

type I_newarray struct {
}

func init()  {
	INSTRUCTION_MAP[0xbc] = &I_newarray{}
}

func (s I_newarray)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "newarray exce >>>>>>>>>\n")

	typer := ctx.Code[ctx.PC]
	ctx.PC++
	count, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(count) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	if count.(types.Jint) < 0 {
		except, _ := variator.AllocExcept(variator.NegativeArraySizeException)
		ctx.Throw(except)
		return nil
	}

	jarray := oli.AllocJarray(count.(types.Jint), int(typer), 1)

	ctx.CurrentFrame.PushFrame(jarray)

	return nil
}

func (s I_newarray)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		创建一个数组
======================================================================================
						||		newarray
						||------------------------------------------------------------
						||		atype
						||------------------------------------------------------------
						||
		格式				||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		newarray = 188(0xbc)
======================================================================================
						||		...，count →
	   操作数栈			||------------------------------------------------------------
						||		„，arrayref
======================================================================================
						||
		描述				||		count 为 int 类型的数据，指令执行时它将从操作数栈中出栈，它代表了要 创建多大的数组。
atype 为要创建数组的元素类型，它将为以下值之一:

数组类型			atype
T_BOOLEAN		4
T_CHAR			5
T_FLOAT			6
T_DOUBLE		7
T_BYTE			8
T_SHORT			9
T_INT			10
T_LONG			11

一个以 atype 为组件类型、以 count 值为长度的数组将会被分配在 GC 堆中，
并且一个代表该数组的 reference 类型数据 arrayref 压入到操作数栈中。 这个新数组的所有元素将会被分配为相应类型的初始值(§2.3，§2.4)

						||
======================================================================================
						||		
	   链接时异常			||
						||
						||
======================================================================================
						||
	   运行时异常			||
						||如果 count 值小于 0 的话，newarray 指令将会抛出一个 NegativeArraySizeException 异常。
						||
======================================================================================
						||
		注意				||
						||在 Oracle 实现的 Java 虚拟机中，布尔类型(atype 值为 T_BOOLEAN)是 以 8 位储存，并使用 baload 和 bastore 指令操作，这些指令也可以操作 byte
类型的数组。其他 Java 虚拟机可能有自己的 boolean 型数组实现方式，但 必须保证 baload 和 bastore 指令依然使用于它们的 boolean 类型数组。
						||
						||
======================================================================================
 */