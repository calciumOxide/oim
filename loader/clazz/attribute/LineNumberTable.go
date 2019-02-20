package attribute

import "../../../utils"

/*
LineNumberTable 属性是可选变长属性，位于 Code(§4.7.3)结构的属性表。
它被调试 器用于确定源文件中行号表示的内容在 Java 虚拟机的 code[]数组中对应的部分。
在 Code 属性 的属性表中，LineNumberTable 属性可以按照任意顺序出现，此外，
多个 LineNumberTable 属性可以共同表示一个行号在源文件中表示的内容，
即 LineNumberTable 属性不需要与源文件 的行一一对应。
*/
type LineNumberTable struct {
	LineNumberTableLength uint16 //line_number_table[]数组的成员个 数。
	LineNumberTableInfo []*LineNumberTableInfo
}

type LineNumberTableInfo struct {
	StartPc uint16 //是 code[]数组的一个索引，code[]数组在该索引处的字符 表示源文件中新的行的起点。start_pc 项的值必须小于当前 LineNumberTable 属性所在的 Code 属性的 code_length 项的值。
	LineNumber uint16 //line_number 项的值必须与源文件的行数相匹配。
}

func AllocLineNumberTable(b []byte) (*LineNumberTable, int) {
	offset := 2
	v := LineNumberTable{
		LineNumberTableLength: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.LineNumberTableLength; i++ {
		a, s := AllocLineNumberTableInfo(b[offset:])
		v.LineNumberTableInfo = append(v.LineNumberTableInfo, a)
		offset += s
	}
	return &v, offset
}

func AllocLineNumberTableInfo(b []byte) (*LineNumberTableInfo, int) {
	return &LineNumberTableInfo{
		StartPc: utils.BigEndian2Little4U2(b[:2]),
		LineNumber: utils.BigEndian2Little4U2(b[2 : 4]),
	}, 4
}