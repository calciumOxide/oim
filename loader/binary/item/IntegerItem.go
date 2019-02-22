package item

import "../../../utils"

type IntegerItemBin struct {
	Value uint32 //按照 Big-Endian 的顺序存储
}

func AllocIntegerItem(b []byte) (*IntegerItemBin, int) {
	v := IntegerItemBin{
		Value: utils.BigEndian2Little4U4(b[:4]),
	}
	return &v, 4
}
