package item

import "../../../utils"

type FieldRefItemBin struct {
	ClassIndex uint16 //对常量池的有效索引，常量池在该索引处的项必须是 CONSTANT_Class_info(§4.4.1)结构，表示一个类或接口，当前字段或方法是这 个类或接口的成员 mothodRef必须是类，interfaceMethodRef必须是接口
	NameAndTypeIndex uint16 //name_and_type_index 项的值必须是对常量池的有效索引，常量池在该索引处的项必是 CONSTANT_NameAndType_info(§4.4.6)结构，它表示当前字段或方法的名 字和描述符。
}

func AllocFieldRefItem(b []byte) (*FieldRefItemBin, int) {
	v := FieldRefItemBin {
		ClassIndex: utils.BigEndian2Little4U2(b[:2]),
		NameAndTypeIndex: utils.BigEndian2Little4U2(b[2:4]),
	}
	return &v, 4
}