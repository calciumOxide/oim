package attribute

import "../../../utils"

/*
Signature 属性是可选的定长属性，位于 ClassFile(§4.1)，field_info(§4.5) 或 method_info(§4.6)结构的属性表中。
在 Java 语言中，任何类、接口、初始化方法或成 员的泛型签名如果包含了类型变量(Type Variables)或参数化类型(Parameterized Types)，
则 Signature 属性会为它记录泛型签名信息。
*/
type Signature struct {
	SignatureIndex uint16
}
/*
signature_index 项的值必须是一个对常量池的有效索引。
常量池在该索引处的项必 须是 CONSTANT_Utf8_info(§4.4.7)结构，
表示类签名或方法类型签名或字段类 型签名:如果当前的 Signature 属性是 ClassFile 结构的属性，
则这个结构表示类签 名，如果当前的 Signature 属性是 method_info 结构的属性，
则这个结构表示方法类 型签名，如果当前 Signature 属性是 field_info 结构的属性，则这个结构表示字段 类型签名。
*/

func AllocSignature(b []byte) (*Signature, int) {
	return &Signature{
		SignatureIndex: utils.BigEndian2Little4U2(b[:2]),
	}, 2
}