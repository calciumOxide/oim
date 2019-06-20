package item

import "../../../utils"

type DoubleItemBin struct {
	ValueH    uint32
	ValueL    uint32
	Value     float64
	Overflow  bool
	Underflow bool
	NaN       bool
}

func AllocDoubleItem(b []byte) (*DoubleItemBin, int) {
	v := new(DoubleItemBin)
	v.ValueH = utils.BigEndian2Little4U4(b[:4])
	v.ValueL = utils.BigEndian2Little4U4(b[4:8])
	l := (uint64(v.ValueH) << 32) + uint64(v.ValueL)
	if l == OVER_FLOW {
		v.Overflow = true
	} else if l == UNDER_FLOW {
		v.Underflow = true
	} else if (l >= NAN_MIN_1 && l <= NAN_MAX_1) || (l >= NAN_MIN_2 && l <= NAN_MAX_2) {
		v.NaN = true
	} else {
		s := int64(-1)
		e := (uint64)((l >> 52) & 0x7ff)
		m := (l & 0xfffffffffffff) << 1
		if (l >> 63) == 0 {
			s = 1
		}
		if e == 0 {
			m = (l & 0xfffffffffffff) | 0x10000000000000
		}
		c := uint64(1)
		u := e - 1075
		if u >= 0 {
			c <<= u
		} else {
			c >>= -u
		}
		v.Value = float64(s) * float64(m) * float64(u)
	}
	return v, 8
}

const (
	OVER_FLOW  uint64 = 0X7FF0000000000000
	UNDER_FLOW uint64 = 0XFFF0000000000000
	NAN_MIN_1  uint64 = 0X7FF0000000000001
	NAN_MAX_1  uint64 = 0X7FFFFFFFFFFFFFFF
	NAN_MIN_2  uint64 = 0XFFF0000000000001
	NAN_MAX_2  uint64 = 0XFFFFFFFFFFFFFFFF
)
