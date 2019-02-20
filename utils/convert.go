package utils

func BigEndian2Little4U4(b []byte) uint32 {
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func BigEndian2Little4U2(b []byte) uint16 {
	return uint16(b[1]) | uint16(b[0])<<8
}
