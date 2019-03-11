package instructions

import (
	"../runtime"
	"../../utils"
)

type I_aloadX struct {
}

func init()  {
	INSTRUCTION_MAP[0x2A] = &I_aloadX{}
	INSTRUCTION_MAP[0x2B] = &I_aloadX{}
	INSTRUCTION_MAP[0x2C] = &I_aloadX{}
	INSTRUCTION_MAP[0x2D] = &I_aloadX{}
}

func (s I_aloadX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "aloadX exce >>>>>>>>>\n")

	value, _ := ctx.CurrentAborigines.GetAborigines(uint32(ctx.Code[ctx.PC  - 1]) - 0x2A)
	ctx.CurrentFrame.PushFrame(value)
	return nil
}

func (s I_aloadX)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.Depth = 0
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{12})
	a.Layers = append(a.Layers, &[]uint32{34})
	a.Layers = append(a.Layers, &[]uint32{56})
	a.Layers = append(a.Layers, &[]uint32{78})
	return &runtime.Context{
		Code: []byte{0x2C},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		从局部变量表加载一个 reference 类型值到操作数栈中
======================================================================================
						||		aload_<n>
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
						||		aload_0 = 42(0x2a)
						||------------------------------------------------------------
		结构				||		aload_1 = 43(0x2b)
						||------------------------------------------------------------
						||		aload_2 = 44(0x2c)
						||------------------------------------------------------------
						||		aload_3 = 45(0x2d)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，objectref
======================================================================================
						||		
						||		<n>代表当前栈帧(§2.6)中局部变量表的索引值，<n>作为索引定位的局部 变量必须为 reference 类型，
		描述				||		称为 objectref。指令执行后，objectref 将会压入到操作数栈栈顶
						||
						||		
======================================================================================
						||		
						||		
						||
	   运行时异常			||
						||		
						||		
						||		
======================================================================================
						||
						||
						||		aloadX 指令无法被用于加载类型为 returnAddress 类型的数据到操作数栈 中，
		注意				||		这点是特意设计成与 astore 指令不相对称的(astore 指令可以操作returnAddress 类型的数据)。
						||		aloadX 操作码可以与 wide 指令联合一起实现使用 2 个字节长度的无符号 byte 型数值作为索引来访问局部变量表。
						||
						||
======================================================================================
 */