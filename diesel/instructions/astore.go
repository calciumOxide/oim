package instructions

import (
		"../runtime"
	"../../utils"
		)

type I_astore struct {
}

func init()  {
	INSTRUCTION_MAP[0x3a] = &I_astore{}
}

func (s I_astore)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "astore exce >>>>>>>>>\n")

	index := ctx.Code[ctx.PC]
	ref, _ := ctx.CurrentFrame.PopFrame()
	ctx.PC += 1

	ctx.CurrentAborigines.SetAborigines(uint32(index), ref)

	return nil
}

func (s I_astore)Test() *runtime.Context {

	f1 := new(runtime.Frame)
	f1.PushFrame(9999)

	var a []interface{}
	a = append(append(append(a, 1), 2), 4)

	return &runtime.Context{
		PC: 0,
		Code: []byte{0x3a, 0x2},
		CurrentFrame: f1,
		CurrentAborigines: &runtime.Aborigines{
			Layers: a,
		},
	}
}
/**
======================================================================================
		操作				||		将一个 reference 类型数据保存到局部变量表中
======================================================================================
						||		astore
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
		结构				||		astore = 58(0x3a)
======================================================================================
						||		...，objectref →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
						||
						||		index 是一个无符号 byte 型整数，它必须是一个指向当前栈帧(§2.6)局 部变量表的索引值，
						||		而在操作数栈栈顶的 objectref 必须是 returnAddress 或者 reference 类型的数据，
		描述				||		这个数据将从操作数栈出栈， 然后保存到 index 所指向的局部变量表位置中。
						||
						||
						||
						||
======================================================================================
						||		
						||
						||
	   链接时异常			||
						||		
						||		
						||		
======================================================================================
						||
						||
	   运行时异常			||
						||
						||
======================================================================================
						||
						||		astore 指令可以与 returnAddress 类型的数据配合来实现 Java 语言中的 finally 子句(参见§3.13，“编译 fianlly”)。
						||		但是 aload 指令不可以用 来从局部变量表加载 returnAddress 类型的数据到操作数栈，这种 astore指令的不对称性是有意设计的。
						||
		注意				||		astore 指令可以与 wide 指令联合使用，以实现使用 2 字节宽度的无符号整 数作为索引来访问局部变量表。
						||
						||
						||
======================================================================================
 */