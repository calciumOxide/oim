package clazz

import "../../utils"

type Field struct {
	AccessFlags     AccessFlags
	NameIndex       uint16 //name_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的项必须是 CONSTANT_Utf8_info(§4.4.7)结构，表示一个有效的字段的非全限定名(§ 4.2.2)
	DescriptorIndex uint16 //descriptor_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的项必 须是 CONSTANT_Utf8_info(§4.4.7)结构，表示一个有效的字段的描述符(§ 4.3.2)
	AttributesCount uint16 //当前字段的附加属性(§4.7)的数量
	Attributes      []*Attribute
	Value			interface{}
	/*
	本规范所定义的 field_info 结构中，attributes 表可出现的成员有:
	ConstantValue(§4.7.2), Synthetic(§4.7.8), Signature(§4.7.9), Deprecated(§4.7.15),
	RuntimeVisibleAnnotations(§4.7.16) 和 RuntimeInvisibleAnnotations(§4.7.17)
	*/
}

func AllocField(b []byte, cf *ClassFile) (*Field, int) {
	offset := 0
	v := new(Field)
	v.AccessFlags = AccessFlags(utils.BigEndian2Little4U2(b[:offset+2]))
	offset += 2
	v.NameIndex = utils.BigEndian2Little4U2(b[offset : offset+2])
	offset += 2
	v.DescriptorIndex = utils.BigEndian2Little4U2(b[offset : offset+2])
	offset += 2
	v.AttributesCount = utils.BigEndian2Little4U2(b[offset : offset+2])
	offset += 2
	for i:=uint16(0); i<v.AttributesCount; i++ {
		attr, size := AllocAttribute(b[offset:], cf)
		v.Attributes = append(v.Attributes, attr)
		offset += size
	}
	return v, offset
}
