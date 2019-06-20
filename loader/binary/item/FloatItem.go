package item

import "../../../utils"

type FloatItemBin struct {
	Bytes     uint32 //按照 Big-Endian 的顺序存储
	Value     float32
	Overflow  bool
	Underflow bool
	NaN       bool
}

func AllocFloatItem(b []byte) (*FloatItemBin, int) {
	v := FloatItemBin{
		Bytes: utils.BigEndian2Little4U4(b[:4]),
	}
	if v.Bytes == OVER_FLOW_FLOAT {
		v.Overflow = true
	} else if v.Bytes == UNDER_FLOW_FLOAT {
		v.Underflow = true
	} else if (v.Bytes >= NAN_MIN_1_FLOAT && v.Bytes <= NAN_MAX_1_FLOAT) || (v.Bytes >= NAN_MIN_2_FLOAT && v.Bytes <= NAN_MAX_2_FLOAT) {
		v.NaN = true
	} else {
		s := int64(-1)
		e := (uint64)((v.Bytes >> 23) & 0xff)
		m := (v.Bytes & 0x7fffff) << 1
		if (v.Bytes >> 31) == 0 {
			s = 1
		}
		if e == 0 {
			m = (v.Bytes & 0x7fffff) | 0x800000
		}
		c := uint64(1)
		u := e - 1075
		if u >= 0 {
			c <<= u
		} else {
			c >>= -u
		}
		v.Value = float32(s) * float32(m) * float32(u)
	}
	return &v, 4
}

const (
	OVER_FLOW_FLOAT  uint32 = 0x7f800000
	UNDER_FLOW_FLOAT uint32 = 0xff800000
	NAN_MIN_1_FLOAT  uint32 = 0x7f800001
	NAN_MAX_1_FLOAT  uint32 = 0x7fffffff
	NAN_MIN_2_FLOAT  uint32 = 0xff800001
	NAN_MAX_2_FLOAT  uint32 = 0xffffffff
)
