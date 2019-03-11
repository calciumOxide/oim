package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_l2i struct {
}

func init()  {
	INSTRUCTION_MAP[0x88] = &I_l2i{}
}

func (s I_l2i)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "l2i exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jint(value.(types.Jlong) & 0xFFFFFFFF))
	return nil
}

func (s I_l2i)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jlong(-1234))
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
		操作				||		将 long 类型数据转换为 int 类型
======================================================================================
						||		l2i
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
		结构				||		l2i = 136(0x88)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
		描述				||		value 必须是在操作数栈栈顶的 long 类型数据，指令执行时，它将从操作数 栈中出栈，使用保留低 32 位、丢弃高 32 位的方式转换为 int 类型数据，然
后压入到操作数栈之中。
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
		注意				||		i2d 指令执行了窄化类型转换(Narrowing Primitive Conversion，JLS3 §5.1.3)，它可能会导致 value 的数值大小发生改变，甚至导致转换结果与
原值有不同的正负号。
						||
						||
======================================================================================
 */