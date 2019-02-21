package item

import "../../../utils"

type MethodType struct {
	DescriptorIndex uint16 //对常量池的有效索引，常量池在该索引处的项必须是 CONSTANT_Utf8_info(§4.4.7)结构，表示方法的描述符(§4.3.3)
}

func AllocMethodTypeItem(b []byte) (*MethodType, int) {
	return &MethodType {
		DescriptorIndex: utils.BigEndian2Little4U2(b[:2]),
	}, 2
}