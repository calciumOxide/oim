package clazz

import "./item"

type ClassFile struct {
	Magic uint32 //魔数
	MinorVersion uint16 //副版本
	MajorVersion uint16 //主版本
	ConstantPoolCount uint16 //常量池大小
	ConstantPool []*ConstantPool //常量池
	AccessFlags AccessFlags //访问标识
	ThisClass uint16 //类索引，this_class 的值必须是对 constant_pool 表中项目的一个有效索引值。 constant_pool 表在这个索引处的项必须为 CONSTANT_Class_info 类型常量
	SuperClass uint16 //父类索引，对于类来说，super_class 的值必须为 0 或者是对 constant_pool 表中 项目的一个有效索引值
	InterfacesCount uint16 //接口计数器，interfaces_count 的值表示当前类或接口的直接父接口数量
	Interfaces []uint16 //接口表，interfaces[]数组中的每个成员的值必须是一个对 constant_pool 表中项 目的一个有效索引值
	FieldsCount uint16 //字段计数器，fields_count 的值表示当前 Class 文件 fields[]数组的成员个数
	Fields []*Field //字段表，fields[]数组中的每个成员都必须是一个 fields_info 结构,描述当前类或接口 声明的所有字段，但不包括从父类或父接口继承的部分
	MethodsCount uint16
	Methods []*Method //方法表，methods[]数组中的每个成员都必须是一个 method_info 结构
	AttributesCount uint16
	Attributes []*Attribute //属性表，attributes 表的每个项的值必须是 attribute_info 结构
}

func AllocClassFile() *ClassFile {
	return new(ClassFile)
}

func (s *ClassFile) GetConstant(index uint16) (*ConstantPool, error) {
	return s.ConstantPool[index - 1], nil
}

var CLASS_MAP = make(map[string] *ClassFile)

func GetClass(className string) *ClassFile {
	return CLASS_MAP[className]
}


func (s *ClassFile) GetMethod(andType *item.NameAndType) *Method {
	for i := uint16(0); i < s.MethodsCount; i++ {
		//method := s.Methods[i]
		//if !method.IsStatic() && method.GetName() == andType.GetName() && method.GetDescriptor() == andType.GetDescriptor() {
		//	return method
		//}
	}
	if s.SuperClass == 0 {
		return nil
	}
	pool, _ := s.GetConstant(s.SuperClass)
	superClass := GetClass(pool.Info.(*item.Class).GetName())
	return superClass.GetMethod(andType)
}