package attribute

import "../../../utils"

/*
EnclosingMethod 属性是可选的定长属性，位于 ClassFile(§4.1)结构的属性表。
当 且仅当 Class 为局部类或者匿名类时，才能具有 EnclosingMethod 属性。
一个类最多只能有一 个 EnclosingMethod 属性。
*/

type EnclosingMethod struct {
	ClassIndex uint16 //class_index 项的值必须是一个对常量池的有效索引。常量池在该索引出的项必须是 CONSTANT_Class_info(§4.4.1)结构，表示包含当前类声明的最内层类。
	MethodIndex uint16 //如果当前类不是在某个方法或初始化方法中直接包含(Enclosed)，那么 method_index 值为 0，否则 method_index 项的值必须是对常量池的一个有效索引，
					   // 常量池在该索引处的成员必须是 CONSTANT_NameAndType_info(§4.4.6)结构， 表示由 class_index 属性引用的类的对应方法的方法名和方法类型。
					   // Java 编译器有责 任在语法上保证通过 method_index 确定的方法是语法上最接近那个包含 EnclosingMethod属性的类的方法(Closest Lexically Enclosing Method)。
}

func AllocEnclosingMethod(b []byte) (*EnclosingMethod, int) {
	return &EnclosingMethod {
		ClassIndex: utils.BigEndian2Little4U2(b[:2]),
		MethodIndex: utils.BigEndian2Little4U2(b[2 : 4]),
	}, 4
}