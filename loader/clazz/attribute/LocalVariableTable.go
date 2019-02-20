 package attribute

import "../../../utils"

type LocalVariableTable struct {
	LocalVariableTableLength uint16
	LocalVariableTableInfo []*LineNumberTableInfo
}

type LocalVariableTableInfo struct {
	StartPc uint16
	Length uint16
	NameIndex uint16 //name_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的成员必 须是 CONSTANT_Utf8_info(§4.4.7)结构，表示一个局部变量的有效的非全 限定名(§4.2.2)
	DescriptorIndex uint16 //descriptor_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的成员必须是 CONSTANT_Utf8_info(§4.4.7)结构，表示源程序中局部变量类型的字段描述符(§4.3.2)
	Index uint16 //index 为此局部变量在当前栈帧的局部变量表中的索引。如果在 index 索引处的局 部变量是 long 或 double 型，则占用 index 和 index+1 两个索引。
}
/*
startPc uint16
length uint16
所有给定的局部变量的索引都在范围[start_pc, start_pc+length)中，
即从 start_pc(包括自身值)至 start_pc+length(不包括自身值)。
start_pc 的 值必须是一个对当前 Code 属性的 code[]数组的有效索引，
code[]数组在这个索 引处必须是一条指令操作码。start_pc+length 要么是当前 Code 属性的 code[] 数组的有效索引，
code[]数组在该索引处必须是一条指令的操作码，要么是刚超过 code[]数组长度的最小索引值。
*/

func AllocLocalVariableTable(b []byte) (*LocalVariableTable, int) {
	offset := 2
	v := LocalVariableTable{
		LocalVariableTableLength: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.LocalVariableTableLength; i++ {
		AllocLocalVariableTableInfo(b[offset : offset + 10])
		offset += 10
	}
	return &v, offset
}

func AllocLocalVariableTableInfo(b[]byte) (*LocalVariableTableInfo, int) {
	return &LocalVariableTableInfo {
		StartPc: utils.BigEndian2Little4U2(b[:2]),
		Length: utils.BigEndian2Little4U2(b[2 : 4]),
		NameIndex: utils.BigEndian2Little4U2(b[4 : 6]),
		DescriptorIndex: utils.BigEndian2Little4U2(b[6 : 8]),
		Index: utils.BigEndian2Little4U2(b[8 : 10]),
	}, 10
}