package instructions

import (
	"../runtime"
	"../../utils"
	"../variator"
	"../../types"
	"reflect"
)

type I_dstoreX struct {
}

func init()  {
	INSTRUCTION_MAP[0x47] = &I_dstoreX{}
	INSTRUCTION_MAP[0x48] = &I_dstoreX{}
	INSTRUCTION_MAP[0x49] = &I_dstoreX{}
	INSTRUCTION_MAP[0x4a] = &I_dstoreX{}
}

func (s I_dstoreX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "dstoreX exce >>>>>>>>>\n")

	index := uint32(ctx.Code[ctx.PC - 1]) - 0x47
	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.JDN) && reflect.TypeOf(value) != reflect.TypeOf(types.Jdouble(0)) &&
		reflect.TypeOf(value) != reflect.TypeOf(types.JDO) && reflect.TypeOf(value) != reflect.TypeOf(types.JDU){
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentAborigines.SetAborigines(index, value)
	return nil
}

func (s I_dstoreX)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(2))
	f.PushFrame(types.Jdouble(33))

	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, 1, 2, 3, 4)
	return &runtime.Context{
		Code: []byte{0x4a, 0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		将一个 double 类型数据保存到局部变量表中
======================================================================================
						||		dstore_<n>
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
						||		dstore_0 = 71(0x47)
						||------------------------------------------------------------
						||		dstore_1 = 72(0x48)
		结构				||------------------------------------------------------------
						||		dstore_2 = 73(0x49)
						||------------------------------------------------------------
						||		dstore_3 = 74(0x4a)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
						||		<n>和<n>+1 必须是一个指向当前栈帧(§2.6)局部变量表的索引值，而在 操作数栈栈顶的 value 必须是 double 类型的数据，
		描述				||		这个数据将从操作数栈出栈，并且经过数值集合转换(§2.8.3)后得到值 value’，
						||		然后保存到<n> 和<n>+1 所指向的局部变量表位置中。
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
		注意				||		dstore_<n>指令族中的每一条指令都与使用<n>作为 index 参数的 dstore 指令作的作用一致，仅仅除了操作数<n>是隐式包含在指令中这点不同而已。
						||
						||
						||
======================================================================================
 */