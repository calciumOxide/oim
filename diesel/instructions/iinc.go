package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
		)

type I_iinc struct {
}

func init()  {
	INSTRUCTION_MAP[0x84] = &I_iinc{}

}

func (s I_iinc)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "iinc exce >>>>>>>>>\n")

	index := uint32(0)
	addor := uint32(0)
	if ctx.PopWide() {
		index = uint32(utils.BigEndian2Little4U2(ctx.Code[ctx.PC : ctx.PC + 2]))
		ctx.PC += 2
		addor = uint32(utils.BigEndian2Little4U2(ctx.Code[ctx.PC : ctx.PC + 2]))
		ctx.PC += 2
	} else {
		index = uint32(ctx.Code[ctx.PC])
		addor = uint32(ctx.Code[ctx.PC + 1])
		ctx.PC += 2
	}

	value := ctx.CurrentAborigines.Layers[index]
	ctx.CurrentAborigines.Layers[index] = value.(types.Jint) + types.Jint(addor)
	return nil
}

func (s I_iinc)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(&types.Jreference{Reference: types.Jobject{}})
	f.PushFrame(nil)
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, types.Jint(5))
	return &runtime.Context{
		Code: []byte{0x84, 0x0, 12},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		局部变量自增
======================================================================================
						||		iinc
						||------------------------------------------------------------
						||		index
						||------------------------------------------------------------
						||		const
		格式				||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		iinc = 132(0x84)
======================================================================================
						||		...， →
	   操作数栈			||---------------- 无改变  ------------------------------------
						||		...，
======================================================================================
						||
						||		index 是一个代表当前栈帧(§2.6)中局部变量表的索引的无符号 byte 类 型整数，const 是一个有符号的 byte 类型数值。
		描述				||		由 index 定位到的局部变量必须是 int 类型，const 首先带符号扩展成一个 int 类型数值，
						||		然后加到 由 index 定位到的局部变量中。
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
		注意				||		iinc 操作码可以与 wide 指令联合一起实现使用 2 个字节长度的无符号 byte 型数值作为索引来访问局部变量表以及令局部变量增加 2 个字节长度的有符号数值。
						||
						||
======================================================================================
 */