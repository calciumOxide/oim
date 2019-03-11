package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_l2f struct {
}

func init()  {
	INSTRUCTION_MAP[0x89] = &I_l2f{}
}

func (s I_l2f)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "l2f exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(types.Jfloat(value.(types.Jlong)))
	return nil
}

func (s I_l2f)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		将 long 类型数据转换为 float 类型
======================================================================================
						||		l2f
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
		结构				||		l2f = 137(0x89)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
		描述				||		value 必须是在操作数栈栈顶的 long 类型数据，指令执行时，它将从操作数 栈中出栈，使用IEEE 754规范的向最接近数舍入模式转换成float类型数
         据，然后压入到操作数栈之中。

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
		注意				||		i2d 指令执行了宽化类型转换(Widening Primitive Conversion，JLS3 §5.1.2)，由于 double 类型只有 24 位有效位数，所以转换可能会产生精度
丢失。
						||
						||
======================================================================================
 */