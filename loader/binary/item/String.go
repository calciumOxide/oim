package item

import "../../../utils"

type String struct {
	StringIndex uint16 //对常量池的有效索引，常量池在该索引处的项必须是 CONSTANT_Utf8_info
}

func AllocStringItem(b []byte) (*String, int) {
	v := String{
		StringIndex: utils.BigEndian2Little4U2(b[:2]),
	}
	return &v, 2
}