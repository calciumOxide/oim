package item

import "../../../utils"

type Long struct {
	ValueH uint32
	ValueL uint32
	Value int64
}

func AllocLongItem(b []byte) (*Long, int) {
	v := Long {
		ValueH: utils.BigEndian2Little4U4(b[:4]),
		ValueL: utils.BigEndian2Little4U4(b[4:8]),
	}
	v.Value = (int64(v.ValueH) << 32) + int64(v.ValueL)
	return &v, 8
}