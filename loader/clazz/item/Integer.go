package item

import "../../../utils"

type Integer struct {
	Value uint32 //按照 Big-Endian 的顺序存储
}

func AllocIntegerItem(b []byte) (*Integer, int) {
	v := Integer{
		Value: utils.BigEndian2Little4U4(b[:4]),
	}
	return &v, 4
}
