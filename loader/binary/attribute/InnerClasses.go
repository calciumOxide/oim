package attribute

import "../../../utils"

/*
InnerClasses 属性是一个变长属性，位于 ClassFile(§4.1)结构的属性表。
本小结为了方便说明特别定义一个表示类或接口的 Class 格式为 C。
如果 C 的常量池中包含某个 CONSTANT_Class_info 成员，
且这个成员所表示的类或接口不属于任何一个包，那么 C 的 ClassFile 结构的属性表中就必须含有对应的 InnerClasses 属性
*/
type InnerClasses struct {
	ClassCount uint16
	Classes    []*Classes
}

type Classes struct {
	InnerClassIndex       uint16                //inner_class_info_index 项的值必须是一个对常量池的有效索引。常量池在该索引 处的项必须是 CONSTANT_Class_info(§4.4.1)结构，表示接口 C。当前元素的另外 3 项都用于描述 C 的信息
	OuterClassIndex       uint16                //如果 C 不是类或接口的成员(也就是 C 为顶层类或接口(JLS §7.6)、局部类(JLS § 14.3)或匿名类(JLS §15.9.5)),那么 outer_class_info_index 项的值为 0， 否则这个项的值必须是对常量池的一个有效索引，常量池在该索引处的项必须是 CONSTANT_Class_info(§4.4.1)结构，代表一个类或接口，C 为这个类或接口的 成员。
	InnerNameIndex        uint16                //如果 C 是匿名类(JLS §15.9.5)，inner_name_index 项的值则必须为 0。否则这 个项的值必须是对常量池的一个有效索引，常量池在该索引处的项必须 CONSTANT_Utf8_info(§4.4.7)结构，表示 C 的 Class 文件在对应的源文件中定 义的C的原始简单名称(Original Simple Name)
	InnerClassAccessFlags InnerClassAccessFlags //inner_class_access_flags 项的值是一个掩码标志，用于定义 Class 文件对应的 源文件中 C 的访问权和基本属性。用于编译器在无法访问源文件时可以恢复 C 的原始信 息
}

type InnerClassAccessFlags uint16

const (
	ACC_PUBLIC     InnerClassAccessFlags = 0x0001
	ACC_PRIVATE    InnerClassAccessFlags = 0x0002
	ACC_PROTECTED  InnerClassAccessFlags = 0x0004
	ACC_STATIC     InnerClassAccessFlags = 0x0008
	ACC_FINAL      InnerClassAccessFlags = 0x0010
	ACC_INTERFACE  InnerClassAccessFlags = 0x0200
	ACC_ABSTRACT   InnerClassAccessFlags = 0x0400
	ACC_SYNTHETIC  InnerClassAccessFlags = 0x1000
	ACC_ANNOTATION InnerClassAccessFlags = 0x2000
	ACC_ENUM       InnerClassAccessFlags = 0x4000
)

func AllocInnerClasses(b []byte) (*InnerClasses, int) {
	offset := 0
	v := InnerClasses{
		ClassCount: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.ClassCount; i++ {
		c, s := AllocClasses(b[offset : offset + 8])
		v.Classes = append(v.Classes, c)
		offset += s
	}
	return &v, offset
}

func AllocClasses(b []byte) (*Classes, int) {
	return &Classes{
		InnerClassIndex: utils.BigEndian2Little4U2(b[:2]),
		OuterClassIndex: utils.BigEndian2Little4U2(b[2 : 4]),
		InnerNameIndex: utils.BigEndian2Little4U2(b[4 : 6]),
		InnerClassAccessFlags: InnerClassAccessFlags(utils.BigEndian2Little4U2(b[6 : 8])),
	}, 8
}
