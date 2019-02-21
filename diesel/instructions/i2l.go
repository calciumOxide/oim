package instructions

import (
	"../runtime"
	"../../utils"
	"../oil/types"
	"reflect"
	"../variator"
)

type I_i2l struct {
}

func init()  {
	INSTRUCTION_MAP[0x85] = &I_i2l{}
}

func (s I_i2l)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "i2l exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jint(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jlong(value.(types.Jint)))
	return nil
}

func (s I_i2l)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jint(-1234))
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
		操作				||		将 int 类型数据转换为 long 类型
======================================================================================
						||		i2l
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
		结构				||		i2l = 133(0x85)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
		描述				||		value 必须是在操作数栈栈顶的 int 类型数据，指令执行时，它将从操作数 栈中出栈，
						||		并带符号扩展(Sign-Extended)成 long 类型数据，然后压入到操作数栈之中。
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
		注意				||		i2l 指令执行了宽化类型转换(Widening Primitive Conversion，JLS3 §5.1.2)，因为所有 int 类型的数据都可以精确表示为 long 类型的数据，所以转换是精确的。
						||
						||
						||
======================================================================================
 */