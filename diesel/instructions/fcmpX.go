package instructions

import (
	"../runtime"
	"../../utils"
	"../oil/types"
	"reflect"
	"../variator"
)

type I_fcmpX struct {
}

func init()  {
	INSTRUCTION_MAP[0x96] = &I_fcmpX{}
	INSTRUCTION_MAP[0x95] = &I_fcmpX{}
}

func (s I_fcmpX)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "fcmpX exce >>>>>>>>>\n")

	op := ctx.Code[ctx.PC - 1]
	value2, _ := ctx.CurrentFrame.PopFrame()
	value1, _ := ctx.CurrentFrame.PopFrame()

	if value1 == types.JDN || value2 == types.JDN {
		if op == 0x96 {
			ctx.CurrentFrame.PushFrame(types.Jint(-1))
		} else {
			ctx.CurrentFrame.PushFrame(types.Jint(1))
		}
		return nil
	}

	if reflect.TypeOf(value1) == reflect.TypeOf(types.JDO) {
		value1 = types.Jfloat(types.JDO)
	}
	if reflect.TypeOf(value2) == reflect.TypeOf(types.JDO) {
		value2 = types.Jfloat(types.JDO)
	}
	if reflect.TypeOf(value1) == reflect.TypeOf(types.JDU) {
		value1 = types.Jfloat(types.JDO)
	}
	if reflect.TypeOf(value2) == reflect.TypeOf(types.JDU) {
		value2 = types.Jfloat(types.JDO)
	}

	if reflect.TypeOf(value1) != reflect.TypeOf(types.Jfloat(0)) || reflect.TypeOf(value2) != reflect.TypeOf(types.Jfloat(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	if value1.(types.Jfloat) > value2.(types.Jfloat) {
		ctx.CurrentFrame.PushFrame(types.Jint(1))
	} else if value1.(types.Jfloat) < value2.(types.Jfloat) {
		ctx.CurrentFrame.PushFrame(types.Jint(-1))
	} else {
		ctx.CurrentFrame.PushFrame(types.Jint(0))
	}
	return nil
}

func (s I_fcmpX)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jfloat(9.123456789012345))
	f.PushFrame(types.Jfloat(9.123456789012343))
	//f.PushFrame(nil)
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x96},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		比较 2 个 float 类型数据的大小
======================================================================================
						||		dcmp<op>
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
						||		fcmpg = 150(0x96)
		结构				||------------------------------------------------------------
						||		fcmpl = 149(0x95)
======================================================================================
						||		...，value1，value2 →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||
						||		value1 和 value2 都必须为 float 类型数据，指令执行时，value1 和 value2 从操作数栈中出栈，
						||		并且经过数值集合转换(§2.8.3)后得到值value1’和 value2’，接着对这 2 个值进行浮点比较操作:
						||			 如果 value1’大于 value2’的话，int 值 1 将压入到操作数栈中。
		描述				||			 另外，如果 value1’与 value2’相等的话，int 值 0 将压入到操作数栈中。
						||			 另外，如果 value1’小于 value2’相等的话，int 值-1 将压入到操作数栈中。
						||			 另外，如果 value1’和 value2’之中最少有一个为 NaN，那 dcmpg 指令将 int 值 1 压入到操作数栈中，而 dcmpl 指令则把 int 值-1 压入到操作数栈中。
						||
						||		浮点比较操作将根据IEEE 754规范定义进行，除了NaN之外的所有数值都 是有序的，正无穷大大于所有有限值，正数零和负数零则被看作是相等的。
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
						||		dcmpg 和 dcmpl 指令之间的差别仅仅是当比价参数中出现 NaN 值时的处理方 式而已。
						||		NaN 值是没有顺序的，因此只要参数中出现一个或者两个都为 NaN值时，比较操作就会失败。
		注意				||		但对于 dcmpg 和 dcmpl 指令来说，无论何种情况 导致比较失败，都会有一个确定的返回值压入到操作数栈中。
						||		读者可以参见§3.5“更多控制例子”获取更多的信息。
						||
======================================================================================
 */