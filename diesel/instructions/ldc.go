package instructions

import (
	"../runtime"
	"../../utils"
		"../../types"
	"reflect"
)

type I_ldc struct {
}

func init()  {
	INSTRUCTION_MAP[0x12] = &I_ldc{}
}

func (s I_ldc)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "ldc exce >>>>>>>>>\n")

	index := uint16(ctx.Code[ctx.PC])
	ctx.PC += 1

	ctx.CurrentFrame.PushFrame(ctx.Clazz.GetConstantValue(index))
	return nil
}

func (s I_ldc)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		从运行时常量池中提取数据推入操作数栈
======================================================================================
						||		ldc
						||------------------------------------------------------------
						||		index
						||------------------------------------------------------------
						||		
		格式				||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		ldc = 18(0x12)
======================================================================================
						||		... →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
		描述				||		index 是一个无符号 byte 型数据，它作为当前类(§2.6)的运行时常量池 的索引使用。index 指向的运行时常量池项必须是一个 int 或者 float 类型
的运行时常量，或者是一个类的符号引用(§5.4.3.1)或者字符串字面量 (§5.1)。

如果运行时常量池项必须是一个 int 或者 float 类型的运行时常量，那数值 这个常量所对应的数值 value 将被压入到操作数栈之中。 另外，如果运行时常量池项必须是一个代表字符串字面量(§5.1)的 String 类的引用，那这个实例的引用所对应的 reference 类型数据 value 将被压入 到操作数栈之中。

另外，如果运行时常量池项必须是一个类的符号引用(§4.4.1)，这个符号 引用是已被解析(§5.4.3.1)过的，那这个类的 Class 对象所对应的 reference 类型数据 value 将被压入到操作数栈之中。

======================================================================================
						||		
						||
	   运行时异常			||
						||	在类的符号解析阶段，任何在§5.4.3.1 章节中描述的异常都可能被抛出。
						||		
						||		
======================================================================================
						||
						||
	   运行时异常			||
						||
						||
						||
======================================================================================
						||
		注意				||
						||		ldc 指令只能用来处理单精度浮点集合(§2.3.2)中的 float 类型数据，因为常量池(§4.4.4)中 float 类型的常量必须从单精度浮点集合中选取。
						||
						||
======================================================================================
 */