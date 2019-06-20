package item

import "../../../utils"

type Utf8ItemBin struct {
	Length uint16
	Bytes  []uint8 //字符串值的 byte 数组，bytes[]数组中每个成员的 byte 值都不会是 0， 也不在 0xf0 至 0xff 范围内
	Str    string
}

func AllocUtf8Item(b []byte) (*Utf8ItemBin, int) {
	v := new(Utf8ItemBin)
	v.Length = utils.BigEndian2Little4U2(b[:2])
	v.Bytes = b[2 : 2+v.Length]
	for i := uint16(0); i < v.Length; {
		if v.Bytes[i]&0x80 == 0x00 {
			v.Str += string(int64(v.Bytes[i]))
			i += 1
		} else if v.Bytes[i]&0xE0 == 0xC0 {
			v.Str += string((int64(v.Bytes[i]&0x1f) << 6) + int64(v.Bytes[i+1]&0x3f))
			i += 2
		} else if v.Bytes[i]&0xFF == 0xED {
			v.Str = string(0x10000 + (int64(v.Bytes[i+1]&0x0f) << 16) + (int64(v.Bytes[i+2]&0x3f) << 10) + (int64(v.Bytes[i+4]&0x0f) << 6) + int64(v.Bytes[i+5]&0x3f))
			i += 3
		} else if v.Bytes[i]&0xF0 == 0xE0 {
			v.Str = string((int64(v.Bytes[i]&0xf) << 12) + (int64(v.Bytes[i+1]&0x3f) << 6) + int64(v.Bytes[i+2]&0x3f))
			i += 6
		} else {
			print("======================================================>>>>>>>>>>>>>>>>>>>>>>>>>>")
			break
		}
	}
	return v, 2 + int(v.Length)
}
