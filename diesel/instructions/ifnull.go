package instructions

import (
	"../../utils"
	"../oil/types"
	"../runtime"
	"../variator"
	"reflect"
)

type I_ifnull struct {
}

func init() {
	INSTRUCTION_MAP[0xc6] = &I_ifnull{}

}

func (s I_ifnull) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "ifnull exce >>>>>>>>>\n")

	branch := uint32(ctx.Code[ctx.PC])<<8 | uint32(ctx.Code[ctx.PC+1])
	ctx.PC += 2
	value, _ := ctx.CurrentFrame.PopFrame()

	if value != nil && reflect.TypeOf(value) != reflect.TypeOf(&types.Jreference{}) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	if value == nil || value.(*types.Jreference).Reference == nil {
		ctx.PC = branch
	}
	return nil
}

func (s I_ifnull) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(&types.Jreference{Reference: types.Jobject{}})
	f.PushFrame(nil)
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code:              []byte{0x99, 0x0, 12},
		CurrentFrame:      f,
		CurrentAborigines: a,
	}
}

/**
======================================================================================
		操作				||		引用为空的条件分支判断
======================================================================================
						||		ifnull
						||------------------------------------------------------------
						||		branchbyte1
						||------------------------------------------------------------
						||		branchbyte2
		格式				||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||
======================================================================================
		结构				||		fnull = 198(0xc6)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||
						||
						||		value 必须为 reference 类型数据，指令执行时，value 从操作数栈中出栈，
						||		然后判断是否为 null，如果 value 为 null，那无符号 byte 型数据 branchbyte1 和 branchbyte2 用于构建一个 16 位有符号的分支偏移量，
		描述				||		构建方式为(branchbyte1 << 8)| branchbyte2。
						||		指令执行后，程序将 会转到这个 ifnull 指令之后的，由上述偏移量确定的目标地址上继续执行。这个目标地址必须处于 ifnull 指令所在的方法之中。
						||		另外，如果比较结果为假，那程序将继续执行 ifnull 指令后面的其他直接码 指令。
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
		注意				||
						||
						||
======================================================================================
*/
