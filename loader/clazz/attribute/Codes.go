package attribute

import "../../../utils"

type Codes struct {
	MaxStack          uint16
	MaxLocal          uint16
	CodeLength        uint32
	Code              []uint8
	ExceptTableLength uint16
	ExceptTable       []*ExceptTable
	AttributeCount    uint16
	Attributes        []interface{}
	//Attributes        []*clazz.Attribute   包依赖
	//本规范中定义的、可以出现在 Code 属性的属性表中的成员只能是 LineNumberTable (§4.7.12)，LocalVariableTable(§4.7.13)，LocalVariableTypeTable(§4.7.14)和 StackMapTable(§4.7.4)属性。
}

type ExceptTable struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16 //如果 catch_type 项的值不为 0，那么它必须是对常量池的一个有效索引，常量池 在该索引处的项必须是 CONSTANT_Class_info(§4.4.1)
}
//包依赖
//func AllocCodes(b []byte, cf *clazz.ClassFile) (*Codes, int) {
//	offset := 8
//	v := Codes{
//		MaxStack:   utils.BigEndian2Little4U2(b[:2]),
//		MaxLocal:   utils.BigEndian2Little4U2(b[2:4]),
//		CodeLength: utils.BigEndian2Little4U4(b[4:8]),
//	}
//	offset += int(v.CodeLength)
//	v.Code = b[8:v.CodeLength]
//	v.ExceptTableLength = utils.BigEndian2Little4U2(b[v.CodeLength:2])
//	for i := uint16(0); i < v.ExceptTableLength; i++ {
//		t, size := AllocExceptTable(b[offset:])
//		v.ExceptTable = append(v.ExceptTable, t)
//		offset += size
//	}
//	v.AttributeCount = utils.BigEndian2Little4U2(b[offset : offset + 2])
//	for i := uint16(0); i < v.AttributeCount; i++ {
//		a, size := clazz.AllocAttribute(b[offset:], cf)
//		v.Attributes = append(v.Attributes, a)
//		offset += size
//	}
//	return &v, offset
//}

func AllocExceptTable(b []byte) (*ExceptTable, int) {
	t := ExceptTable{
		StartPc:   utils.BigEndian2Little4U2(b[:2]),
		EndPc:     utils.BigEndian2Little4U2(b[2:4]),
		HandlerPc: utils.BigEndian2Little4U2(b[4:6]),
		CatchType: utils.BigEndian2Little4U2(b[6:8]),
	}
	return &t, 8
}