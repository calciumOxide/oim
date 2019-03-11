package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_multianewarray struct {
}

func init()  {
	INSTRUCTION_MAP[0xc2] = &I_multianewarray{}
}

func (s I_multianewarray)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "multianewarray exce >>>>>>>>>\n")

	index := utils.BigEndian2Little4U2(ctx.Code[ctx.PC : ctx.PC+2])
	dimes := ctx.Code[ctx.PC+2]
	typeCp, _ := ctx.Clazz.GetConstant(index)
	sizeArray := make([]types.Jint, dimes)

	for i := uint8(0); i < dimes; i++ {
		size, _ := ctx.CurrentFrame.PopFrame()
		sizeArray = append(sizeArray, size.(types.Jint))
	}

	obj, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(obj) != reflect.TypeOf(types.Jreference{}) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	if obj.(*types.Jreference).Reference == nil {
		except, _ := variator.AllocExcept(variator.NullPointerException)
		ctx.Throw(except)
		return nil
	}

	if !obj.(*types.Jreference).multianewarray(ctx.ThreadId) {
		except, _ := variator.AllocExcept(variator.IllegalMonitorStateException)
		ctx.Throw(except)
		return nil
	}

	return nil
}

func (s I_multianewarray)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		创建一个新的多维数组
======================================================================================
						||		multianewarray
						||------------------------------------------------------------
						||		indexbyte1
						||------------------------------------------------------------
						||		indexbyte2
		格式				||------------------------------------------------------------
						||		dimensions
						||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		multianewarray = 197(0xc5)
======================================================================================
						||		...，count1，[count2，...]  →
	   操作数栈			||------------------------------------------------------------
						||		„，arrayref
======================================================================================
						||
		描述				||		dimensions 操作数是一个无符号 byte 类型数据，它必须大于或等于 1，代 表创建数组的维度值。相应地，操作数栈中必须包含 dimensions 个数值，
数组每一个值代表每个维度中需要创建的元素数量。这些值必须为非负数 int 类型数据。count1 描述第一个维度的长度，count2 描述第二个维度的长度， 依此类推。

指令执行时，所有 count 都将从操作数栈中出栈，无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运行时常量池的索引值，
构建 方式为(indexbyte1 << 8)| indexbyte2，该索引所指向的运行时常量 池项应当是一个类、接口或者数组类型的符号引用，
这个类、接口或者数组类 型应当是已被解析(§5.4.3.1)的。指令执行产生的结果将会是一个维度不 小于 dimensions 的数组。

一个新的多维数组将会被分配在 GC 堆中，如果任何一个 count 值为 0，那就 不会分配维度。
数组第一维的元素被初始化为第二维的子数组，后面每一维都 依此类推。
数组的最后一个维度的元素将会被分配为数组元素类型的初始值 (§2.3，§2.4)。并且一个代表该数组的 reference 类型数据 arrayref 压入到操作数栈中。

						||
======================================================================================
						||		
	   链接时异常			||
						||	在类、接口或者数组的符号解析阶段，任何在§5.4.3.1 章节中描述的异常 都可能被抛出。
另外，如果当前类没有权限访问数组的元素类型，multianewarray 指令将 会抛出 IllegalAccessError 异常。
						||
======================================================================================
						||
	   运行时异常			||
						||	另外，如果 dimensions 值小于 0 的话，multianewarray 指令将会抛出一 个 NegativeArraySizeException 异常。
						||
======================================================================================
						||
		注意				||
						||对于一维数组来说，使用 newarray 或者 anewarray 指令创建会更加高效。
在运行时常量池中确定的数组类型维度可能比操作数栈中 dimensions 所代 表的维度更高，在这种情况下，multianewarray 指令只会创建数组的第一 个维度。
						||
						||
======================================================================================
 */