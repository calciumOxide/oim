package item

import "../../../utils"

type Class struct {
	NameIndex uint16 //对常量池的一个有效索引。常量池在该索引处的项必须是 CONSTANT_Utf8_info

	//ClassFile *binary.ClassFile
	//Name string
}

//func (s *Class) GetName() string {
//	//if len(s.Name) == 0 {
//	//	pool, _ := s.ClassFile.GetConstant(s.NameIndex)
//	//	s.Name = pool.Info.(*Utf8).Str
//	//}
//	return s.Name
//}

func AllocClassItem(b []byte) (*Class, int) {
	v := new(Class)
	v.NameIndex = utils.BigEndian2Little4U2(b[:2])
	return v, 2
}