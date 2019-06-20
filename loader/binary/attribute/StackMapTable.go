package attribute

import "../../../utils"

/**
StackMapTable 属性是一个变长属性，位于 Code(§4.7.3)属性的属性表中。这个属性 会在虚拟机类加载的类型阶段(§4.10.1)被使用
*/
type StackMapTable struct {
	FrameCount  uint16
	StackFrames []*StackFrame
}

func AllocStackMapTable(b []byte) (*StackMapTable, int) {
	v := StackMapTable{
		FrameCount: utils.BigEndian2Little4U2(b[:2]),
	}
	offset := 0
	for i := uint16(0); i < v.FrameCount; i++ {
		f, s := AllocStackFrame(b[offset:])
		v.StackFrames = append(v.StackFrames, f)
		offset += s
	}
	return &v, 2 + offset
}
