package item

import "../../../utils"

type NameAndType struct {
	NameIndex       uint16 //对常量池的有效索引，常量池在该索引处的项必须是 CONSTANT_Utf8_info(§4.4.7)结构，这个结构要么表示特殊的方法名<init>(§ 2.9)，要么表示一个有效的字段或方法的非限定名(Unqualified Name)
	DescriptorIndex uint16 //对常量池的有效索引，常量池在该索引处的项必须是 CONSTANT_Utf8_info(§4.4.7)结构，这个结构表示一个有效的字段描述符(§ 4.3.2)或方法描述符(§4.3.3)
	Name            string
	Descriptor      string

	//ClassFile		clazz.ClassFile
}

func AllocNameAndTypeItem(b []byte) (*NameAndType, int) {
	return &NameAndType {
		NameIndex: utils.BigEndian2Little4U2(b[:2]),
		DescriptorIndex: utils.BigEndian2Little4U2(b[2:4]),
	}, 4
}

//func (s *NameAndType) GetName() string {
//	if len(s.Name) == 0 {
//		pool, _ := s.ClassFile.GetConstant(s.NameIndex)
//		s.Name = pool.Info.(*Utf8).Str
//	}
//	return s.Name
//}
//
//func (s *NameAndType) GetDescriptor() string {
//	if len(s.Descriptor) == 0 {
//		pool, _ := s.ClassFile.GetConstant(s.DescriptorIndex)
//		s.Descriptor = pool.Info.(*Utf8).Str
//	}
//	return s.Descriptor
//}

