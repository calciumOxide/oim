package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
			"../../loader/clazz"
	"../../loader/clazz/item"
	"../oli"
	)

type I_new struct {
}

func init()  {
	INSTRUCTION_MAP[0xbb] = &I_new{}
}

func (s I_new)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "new exce >>>>>>>>>\n")

	index := utils.BigEndian2Little4U2(ctx.Code[ctx.PC : ctx.PC+2])
	classCp, _ := ctx.Clazz.GetConstant(index)
	classNameIndex := classCp.Info.(*item.Class).NameIndex
	classNameCp, _ := ctx.Clazz.GetConstant(classNameIndex)
	class := clazz.GetClass(classNameCp.Info.(*item.Utf8).Str)
	if !ctx.Cinit(class) {
		return nil
	}

	jobject := oli.AllocJobject(class)

	ctx.CurrentFrame.PushFrame(&types.Jreference{
		ElementType: class,
		Reference: jobject,
	})

	return nil
}

func (s I_new)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		创建一个对象
======================================================================================
						||		new
						||------------------------------------------------------------
						||		indexbyte1
						||------------------------------------------------------------
						||		indexbyte2
		格式				||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		new = 187(0xbb)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		„，objectref
======================================================================================
						||
		描述				||		无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运 行时常量池的索引值，构建方式为(indexbyte1 << 8)| indexbyte2，
该索引所指向的运行时常量池项应当是一个类或接口的符号引用，这个类或接 口类型应当是已被解析(§5.4.3.1)并且最终解析结果为某个具体的类型。
一个以此为类型为对象将会被分配在 GC 堆中，并且它所有的实例变量都会进 行初始化为相应类型的初始值(§2.3，§2.4)。一个代表该对象实例的 reference 类型数据 objectref 将压入到操作数栈中。
对于一个已成功解析但是未初始化(§5.5)的类型，在这时将会进行初始化。

						||
======================================================================================
						||		
	   链接时异常			||
						||在类、接口或者数组的符号解析阶段，任何在§5.4.3.1 章节中描述的异常 都可能被抛出。
另外，如果在类、接口或者数组的符号引用最终被解析为一个接口或抽象类， new 指令将抛出 InstantiationError 异常。
						||
======================================================================================
						||
	   运行时异常			||
						||
						||
======================================================================================
						||
		注意				||
						||new 指令执行后并没有完成一个对象实例创建的全部过程，只有实例初始化方法被执行并完成后，实例才算完全创建。
						||
						||
======================================================================================
 */