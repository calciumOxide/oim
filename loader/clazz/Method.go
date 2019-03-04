package clazz

import "../../utils"
import "./item"

type Method struct {
	AccessFlags AccessFlags
	NameIndex uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes []*Attribute //attributes 表的每一个成员的值必须是 attribute(§4.7)结构，一个方法可以有 任意个与之相关的属性
	/*
	本规范所定义的 method_info 结构中，属性表可出现的成员有:
		Code(§4.7.3)， Exceptions(§4.7.5)，Synthetic(§4.7.8)，Signature(§4.7.9)， Deprecated(§4.7.15)，
	untimeVisibleAnnotations(§4.7.16)， RuntimeInvisibleAnnotations(§4.7.17)， RuntimeVisibleParameterAnnotations(§4.7.18)，
	RuntimeInvisibleParameterAnnotations(§4.7.19)和 AnnotationDefault(§4.7.20)结构。
	*/
	ClassFile *ClassFile
	name string
	descriptor string
}

func AllocMethod(b []byte, cf *ClassFile) (*Method, int) {
	offset := 8
	v := Method {
		AccessFlags: AccessFlags(utils.BigEndian2Little4U2(b[:2])),
		NameIndex: utils.BigEndian2Little4U2(b[2 : 4]),
		DescriptorIndex: utils.BigEndian2Little4U2(b[4 : 6]),
		AttributesCount: utils.BigEndian2Little4U2(b[6 : 8]),
	}
	for i := uint16(0); i < v.AttributesCount; i++ {
		a, s := AllocAttribute(b[offset:], cf)
		v.Attributes = append(v.Attributes, a)
		offset += s
	}
	return &v, offset
}

func (s *Method) GetAttribute(t AttributeName) (*Attribute, error) {
	for _, a := range s.Attributes {
		if a.AttributeName == t {
			return a, nil
		}
	}
	return nil, nil
}

func (s *Method) GetName() string {
	if len(s.name) == 0{
		pool, _ := s.ClassFile.GetConstant(s.NameIndex)
		s.name = pool.Info.(*item.Utf8).Str
	}
	return s.name
}

func (s *Method) GetDescriptor() string {
	if len(s.descriptor) == 0 {
		pool, _ := s.ClassFile.GetConstant(s.DescriptorIndex)
		s.descriptor = pool.Info.(*item.Utf8).Str
	}
	return s.descriptor
}

func (s *Method) IsStatic() bool {
	return s.AccessFlags & ACC_STATIC > 0
}

func (s *Method) IsAbstarct() bool {
	return s.AccessFlags & ACC_ABSTRACT > 0
}

func (s *Method) IsPrivate() bool {
	return s.AccessFlags & ACC_PRIVATE > 0
}


