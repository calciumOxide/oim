package attribute

import "../../../utils"

type StackFrame struct {
	FrameType FrameType
	FrameItem interface{}
}

type FrameType uint8

const (
	ZERO_FRAME                              FrameType = 0
	SAME_FRAME                              FrameType = 63
	SAME_LOCALS_1_STACK_ITEM_FRAME          FrameType = 127
	EMPTY_FRAME_1                           FrameType = 246
	SAME_LOCALS_1_STACK_ITEM_FRAME_EXTENDED FrameType = 247
	CHOP_FRAME                              FrameType = 250
	SAME_FRAME_EXTENDED                     FrameType = 251
	APPEND_FRAME                            FrameType = 254
	FULL_FRAME
)

func AllocStackFrame(b []byte) (*StackFrame, int) {
	offset := 1
	v := StackFrame{
		FrameType: FrameType(b[0]),
	}
	t := v.FrameType
	i := v.FrameItem
	if t < ZERO_FRAME {
		print("===================================================>>>>>> stack fram under")
	} else if t <= SAME_FRAME {
		i = nil
	} else if t <= SAME_LOCALS_1_STACK_ITEM_FRAME {
		ve, size := AllocVerificationTypeInfo(b[offset:])
		i = &SameLocalsAnd1stackItemFrame{
			Stack: *ve,
		}
		offset += size
	} else if t <= EMPTY_FRAME_1 {
		print("===================================================>>>>>> stack fram empty")
	} else if t <= SAME_LOCALS_1_STACK_ITEM_FRAME_EXTENDED {
		ve, size := AllocVerificationTypeInfo(b[offset+2:])
		i = &SameLocalsAnd1stackItemFrameExtended{
			OffsetDelta: utils.BigEndian2Little4U2(b[offset : offset+2]),
			Stack:       *ve,
		}
		offset += 2 + size
	} else if t <= CHOP_FRAME {
		i = &ChopFrame{
			OffsetDelta: utils.BigEndian2Little4U2(b[offset : offset+2]),
		}
		offset += 2
	} else if t <= SAME_FRAME_EXTENDED {
		i = &SameFrameExtended{
			OffsetDelta: utils.BigEndian2Little4U2(b[offset : offset+2]),
		}
	} else if t <= APPEND_FRAME {
		a := AppendFrame{
			OffsetDelta: utils.BigEndian2Little4U2(b[offset : offset+2]),
		}
		i = &a
		offset += 2
		for k := uint8(0); k < uint8(t)-uint8(251); k++ {
			ve, size := AllocVerificationTypeInfo(b[offset:])
			a.Locals = append(a.Locals, ve)
			offset += size
		}
	} else if t <= FULL_FRAME {

	} else {
		print("===================================================>>>>>> stack frame over")
	}
	v.FrameItem = i
	return &v, 1
}

/*
帧类型 same_frame 的类型标记(frame_type)的取值范围是 0 至 63，
如果类型标记所 确定的帧类型是 same_frame 类型，则明当前帧拥有和前一个栈映射帧完全相同的 locals[]数 组，
并且对应的 stack 项的成员个数为 0。当前帧的 offset_delta 值就使用 frame_type 项 的值来表示
*/
type SameFrame struct {
}

/*
帧类型 same_locals_1_stack_item_frame 的类型标记的取值范围是 64 至 127。
如果 类型标记所确定的帧类型是 same_locals_1_stack_item_frame 类型，
则说明当前帧拥有和 前一个栈映射帧完全相同的 locals[]数组，同时对应的 stack[]数组的成员个数为 1。
当前帧 的 offset_delta 值为 frame_type-64。
并且有一个 verification_type_info 项跟随在 此帧类型之后，用于表示那一个 stack 项的成员
*/
type SameLocalsAnd1stackItemFrame struct {
	Stack VerificationTypeInfo
}

/*
帧类型 same_locals_1_stack_item_frame_extended 由值为 247 的类型标记表示，
它表明当前帧拥有和前一个栈映射帧完全相同的 locals[]数组，
同时对应的 stack[]数组的成 员个数为 1。当前帧的 offset_delta 的值需要由 offset_delta 项明确指定。
有一个 stack[]数组的成员跟随在 offset_delta 项之后
*/
type SameLocalsAnd1stackItemFrameExtended struct {
	OffsetDelta uint16
	Stack       VerificationTypeInfo
}

/*
帧类型 chop_frame 的类型标记的取值范围是 248 至 250。
如果类型标记所确定的帧类型是 为 chop_frame，则说明对应的操作数栈为空，
并且拥有和前一个栈映射帧相同的 locals[]数 组，
不过其中的第 k 个之后的 locals 项是不存在的。k 的值由 251-frame_type 确定。
*/
type ChopFrame struct {
	OffsetDelta uint16
}

/*
帧类型 same_frame_extended 由值为 251 的类型标记表示。
如果类型标记所确定的帧类 型是 same_frame_extended 类型，
则说明当前帧有拥有和前一个栈映射帧的完全相同的 locals[]数组，
同时对应的 stack[]数组的成员数量为 0
*/
type SameFrameExtended struct {
	OffsetDelta uint16
}

/*
帧类型 append_frame 的类型标记的取值范围是 252 至 254。
如果类型标记所确定的帧类 型为 append_frame，
则说明对应操作数栈为空，并且包含和前一个栈映射帧相同的 locals[] 数组，
不过还额外附加 k 个的 locals 项。k 值为 frame_type-251
*/
type AppendFrame struct {
	OffsetDelta uint16
	Locals      []*VerificationTypeInfo
}

/*************************************************************************************************************************/
type VerificationTypeInfo struct {
	Tag                  uint8
	VerificationTypeItem interface{}
}

const (
	TOP_VARIABLE = iota
	INTEGER_VARIABLE
	FLOAT_VARIABLE
	LONG_VARIABLE
	DOUBLE_VARIABLE
	NULL_VARIABLE
	UNINITIALIZEDTHIS_VARIABLE
	OBJECT_VARIABLE
	UNINITIALIZED_VARIABLE
)

func AllocVerificationTypeInfo(b []byte) (*VerificationTypeInfo, int) {
	offset := 1
	t := b[0]
	v := VerificationTypeInfo{
		Tag: t,
	}
	i := v.VerificationTypeItem
	if TOP_VARIABLE == t {
		i = nil
	} else if INTEGER_VARIABLE == t {
		i = nil
	} else if FLOAT_VARIABLE == t {
		i = nil
	} else if LONG_VARIABLE == t {
		i = nil
	} else if DOUBLE_VARIABLE == t {
		i = nil
	} else if NULL_VARIABLE == t {
		i = nil
	} else if UNINITIALIZEDTHIS_VARIABLE == t {
		i = nil
	} else if OBJECT_VARIABLE == t {
		i = ObjectVariableInfo{
			ClassIndex: utils.BigEndian2Little4U2(b[offset : offset+2]),
		}
		offset += 2
	} else if UNINITIALIZED_VARIABLE == t {
		i = UninitializedVariableInfo{
			Offset: utils.BigEndian2Little4U2(b[offset : offset+2]),
		}
		offset += 2
	} else {
		print("===================================================>>>>>> verification type no case")
	}
	v.VerificationTypeItem = i
	return &v, offset
}

//Top_variable_info 类型说明这个局部变量拥有验证类型 top(ᴛ)
type TopVariableInfo struct {
}

//Integer_variable_info 类型说明这个局部变量包含验证类型 int
type IntegerVariableInfo struct {
}

//Float_variable_info 类型说明局部变量包含验证类型 float
type FloatVariableInfo struct {
}

/*
Long_variable_info 类型说明存储单元包含验证类型 long，如果存储单元是局部变量， 则要求:
 不能是最大索引值的局部变量。
 按顺序计数的下一个局部变量包含验证类型 ᴛ 如果单元存储是操作数栈成员，则要求:
 当前的存储单元不能在栈顶。
 靠近栈顶方向的下一个存储单元包含验证类型 ᴛ。 Long_variable_info 结构在局部变量表或操作数栈中占用 2 个存储单元。
*/
type LongVariableInfo struct {
}

/*
Double_variable_info 类型说明存储单元包含验证类型 double。如果存储单元是局部 变量，则要求:
 不能是最大索引值的局部变量。
 按顺序计数的下一个局部变量包含验证类型 ᴛ 如果单元存储是操作数栈成员，则要求:
 当前的存储单元不能在栈顶
 靠近栈顶方向的下一个存储单元包含验证类型 ᴛ
*/
type DoubleVariableInfo struct {
}

//Null_variable_info 类型说明存储单元包含验证类型 null。
type NullVariableInfo struct {
}

//UninitializedThis_variable_info 类型说明存储单元包含验证类型 uninitializedThis
type UninitializedThis_variable_info struct {
}

//Object_variable_info 类型说明存储单元包含某个 Class 的实例。由常量池在 cpool_index 给出的索引处的 CONSTANT_CLASS_Info(§4.4.1)结构表示
type ObjectVariableInfo struct {
	ClassIndex uint16
}

//uninitialized(offset)。offset 项给出了一个偏移量，表示在包含此 StackMapTable 属 性的 Code 属性中，new 指令创建的对象所存储的位置
type UninitializedVariableInfo struct {
	Offset uint16
}
