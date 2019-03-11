package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
)

type I_fneg struct {
}

func init()  {
	INSTRUCTION_MAP[0x76] = &I_fneg{}
}

func (s I_fneg)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "fneg exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if value == types.JDN {
		ctx.CurrentFrame.PushFrame(types.JDN)
		return nil
	}

	if reflect.TypeOf(value) == reflect.TypeOf(types.JDO) {
		ctx.CurrentFrame.PushFrame(types.JDU)
		return nil
	}

	if reflect.TypeOf(value) == reflect.TypeOf(types.JDU) {
		ctx.CurrentFrame.PushFrame(types.JDO)
		return nil
	}

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jfloat(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.CurrentFrame.PushFrame(types.Jfloat(value.(types.Jfloat) * -1))
	return nil
}

func (s I_fneg)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jfloat(9.123456789012345))
	f.PushFrame(types.Jfloat(9.123456789012345))
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
		操作				||		float 类型数据取负运算
======================================================================================
						||		fneg
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
		结构				||		fneg = 118(0x76)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
						||		value 必须为 float 类型数据，指令执行时，value 从操作数栈中出栈， 并且经过数值集合转换(§2.8.3)后得到值 value’
						||		接着对这个数进行算 术取负运算，结果转换为 float 类型值 result，最后 result 被压入到操作数栈中。
						||
						||		对于 float 类型数据，取负运算并不等同于与零做减法运算。如果 x 是+ 0.0，那么 0.0-x 等于+0.0，但是-x 则等于-0.0，后面这种一元减法运 算仅仅把数值的符号反转。
						||
		描述				||		下面是一些值得注意的场景:
						||			 如果操作数为 NaN，那运算结果也为 NaN(NaN 值是没有符号的)。
						||			 如果操作数是无穷大，那运算结果是与其符号相反的无穷大。
						||			 如果操作数是零，那运算结果是与其符号相反的零值。
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
						||
		注意				||
						||
						||
						||
======================================================================================
 */