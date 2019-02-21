package attribute

import "../../../utils"

type ConstantValue struct {
	ConstantValueIndex uint16 //对常量池的有效索引。常量池在该索引处 的项给出该属性表示的常量值
}

func AllocConstantValue(b []byte) (*ConstantValue, int) {
	return &ConstantValue {
		ConstantValueIndex: utils.BigEndian2Little4U2(b[:2]),
	}, 2
}