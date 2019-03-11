package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_i2c struct {
}

func init()  {
	INSTRUCTION_MAP[0x92] = &I_i2c{}
}

func (s I_i2c)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "i2c exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jint(types.Jchar(value.(types.Jint))))
	return nil
}

func (s I_i2c)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(-0x1234))
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
		操作				||		将 int 类型数据转换为 char 类型
======================================================================================
						||		i2c
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
		结构				||		i2c = 146(0x92)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
		描述				||		value 必须是在操作数栈栈顶的 int 类型数据，指令执行时，它将从操作数 栈中出栈，
						||		转换成 byte 类型数据，然后零位扩展(Zero-Extended)回一个 int 的结果压入到操作数栈之中。
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
		注意				||		i2c 指令执行了窄化类型转换(Narrowing Primitive Conversion，JLS3 §5.1.3)，它可能会导致 value 的数值大小发生改变，甚至导致转换结果(结果永远为正数)与原值有不同的正负号。
						||
						||
						||
======================================================================================
 */