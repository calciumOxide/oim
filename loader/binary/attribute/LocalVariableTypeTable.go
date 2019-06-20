package attribute

import "../../../utils"

/*
LocalVariableTypeTable 属性是可选变长属性，位于 Code(§4.7.3)的属性表。
它 被用于给调试器确定方法在执行中局部变量的信息，在 Code 属性的属性表中，
LocalVariableTable 属性可以按照任意顺序出现。
Code 属性中的每个局部变量最多只能有一 个 LocalVariableTable 属性。

LocalVariableTypeTable 属性和 LocalVariableTable 属性并不相同，
LocalVariableTypeTable 提供签名信息而不是描述符信息。这仅仅对泛型类型有意义。
泛型 类型的属性会同时出现在 LocalVariableTable 属性和 LocalVariableTypeTable 属性中，
其他的属性仅出现在 LocalVariableTable 属性表中。
*/
type LocalVariableTypeTable struct {
	LocalVariableTypeTableLength uint16
	LocalVariableTypeTableInfo   []*LocalVariableTypeTableInfo
}

/*
local_variable_type_table[]数组的每一个成员表示一个局部变量的值在 code[] 数组中的偏移量范围。它同时也是用于从当前帧的局部变量表找出所需的局部变量的索 引。
*/
type LocalVariableTypeTableInfo struct {
	StartPc        uint16
	Length         uint16
	NameIndex      uint16 //name_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的成员必 须是 CONSTANT_Utf8_info(§4.4.7)结构，表示一个局部变量的有效的非全 限定名(§4.2.2)。
	SignatureIndex uint16 //signature_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的 成员必须是 CONSTANT_Utf8_info(§4.4.7)结构，表示给源程序中局部变量 类型的字段签名(§4.3.4)。
	Index          uint16 //index 为此局部变量在当前栈帧的局部变量表中的索引。如果在 index 索引处的局 部变量是 long 或 double 型，则占用 index 和 index+1 两个索引。
}

/*
start_pc, length
所有给定的局部变量的索引都在范围[start_pc, start_pc+length)中，即从 start_pc(包括自身)至 start_pc+length(不包括自身)。
start_pc 的值必 须是一个对当前 Code 属性的 code[]数组的有效索引，code[]数组在这个索引处 必须是一条指令操作码。
start_pc+length 要么是当前 Code 属性的 code[]数组 的有效索引，code[]数组在该索引处必须是一条指令的操作码，要么是刚超过 code[]数组长度的最小索引值。
*/

func AllocLocalVariableTypeTable(b []byte) (*LocalVariableTypeTable, int) {
	offset := 2
	v := LocalVariableTypeTable{
		LocalVariableTypeTableLength: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.LocalVariableTypeTableLength; i++ {
		AllocLocalVariableTypeTableInfo(b[offset : offset+10])
		offset += 10
	}
	return &v, offset
}

func AllocLocalVariableTypeTableInfo(b []byte) (*LocalVariableTypeTableInfo, int) {
	return &LocalVariableTypeTableInfo{
		StartPc:        utils.BigEndian2Little4U2(b[:2]),
		Length:         utils.BigEndian2Little4U2(b[2:4]),
		NameIndex:      utils.BigEndian2Little4U2(b[4:6]),
		SignatureIndex: utils.BigEndian2Little4U2(b[6:8]),
		Index:          utils.BigEndian2Little4U2(b[8:10]),
	}, 10
}
