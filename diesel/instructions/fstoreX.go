package instructions

import (
	"../../utils"
	"../oil/types"
	"../runtime"
	"../variator"
	"reflect"
)

type I_fstoreX struct {
}

func init() {
	INSTRUCTION_MAP[0x43] = &I_fstoreX{}
	INSTRUCTION_MAP[0x44] = &I_fstoreX{}
	INSTRUCTION_MAP[0x45] = &I_fstoreX{}
	INSTRUCTION_MAP[0x46] = &I_fstoreX{}
}

func (s I_fstoreX) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "fstoreX exce >>>>>>>>>\n")

	index := uint32(ctx.Code[ctx.PC-1]) - 0x43
	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.JFN) && reflect.TypeOf(value) != reflect.TypeOf(types.Jfloat(0)) &&
		reflect.TypeOf(value) != reflect.TypeOf(types.JFO) && reflect.TypeOf(value) != reflect.TypeOf(types.JFU) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentAborigines.SetAborigines(index, value)
	return nil
}

func (s I_fstoreX) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(2))
	f.PushFrame(types.Jfloat(33))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, 1, 2, 3, 4)
	return &runtime.Context{
		Code:              []byte{0x43, 0x0},
		CurrentFrame:      f,
		CurrentAborigines: a,
	}
}

/**
======================================================================================
		操作				||		将一个 float 类型数据保存到局部变量表中
======================================================================================
						||		fstore_<n>
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
						||		fstore_0 = 67(0x43)
						||------------------------------------------------------------
						||		fstore_1 = 68(0x44)
		结构				||------------------------------------------------------------
						||		fstore_2 = 69(0x45)
						||------------------------------------------------------------
						||		fstore_3 = 70(0x46)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||
						||		<n>必须是一个指向当前栈帧(§2.6)局部变量表的索引值，而在操作数栈 栈顶的 value 必须是 float 类型的数据，
		描述				||		这个数据将从操作数栈出栈，并且经过数值集合转换(§2.8.3)后得到值 value’，
						||		然后保存到<n>所指向的 局部变量表位置中。
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
		注意				||		fstore_<n>指令族中的每一条指令都与使用<n>作为 index 参数的 fstore 指令作的作用一致，仅仅除了操作数<n>是隐式包含在指令中这点不同而已。
						||
						||
						||
======================================================================================
*/
