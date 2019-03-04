package clazz

import "./item"
import "../../utils"
import "../butcher/rope"

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
	if className == "java/lang/Object" {
		return nil
	}
	cf := CLASS_MAP[className]
	if cf != nil {
		return cf
	}
	path := "/Users/calciumoxide/coder/znkf/arsenal/console/src/main/java/"
	bytes, _ := rope.ReadClass(path + className + ".class")
	cf, _ = Decoder(bytes)
	CLASS_MAP[className] = cf
	return cf
}

func (s *ClassFile) GetMethod(name, descriptor string) *Method {

	for i := uint16(0); i < s.MethodsCount; i++ {
		method := s.Methods[i]
		mNameCp, _ := s.GetConstant(method.NameIndex)
		mDescCp, _ := s.GetConstant(method.DescriptorIndex)
		if !method.IsStatic() && mNameCp.Info.(*item.Utf8).Str == name && mDescCp.Info.(*item.Utf8).Str == descriptor {
			return method
		}
	}
	if s.SuperClass == 0 {
		return nil
	}
	pool, _ := s.GetConstant(s.SuperClass)
	superNameCp, _ := s.GetConstant(pool.Info.(*item.Class).NameIndex)
	superClass := GetClass(superNameCp.Info.(*item.Utf8).Str)
	return superClass.GetMethod(name, descriptor)
}


func Decoder(bytes []byte) (*ClassFile, bool) {
	var index = 0
	cf := *AllocClassFile()
	cf.Magic = utils.BigEndian2Little4U4((bytes)[index : index + 4]); index += 4
	cf.MinorVersion = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	cf.MajorVersion = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2

	cf.ConstantPoolCount = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	offset, _ := constantPoolInflat(bytes[index:], &cf)
	index = index + offset

	cf.AccessFlags = AccessFlags(utils.BigEndian2Little4U2((bytes)[index : index + 2])); index += 2
	cf.ThisClass = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	cf.SuperClass = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2

	cf.InterfacesCount = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	offset, _ = interfaceInflat(bytes[index:], &cf)
	index = index + offset

	cf.FieldsCount = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	offset, _ = fieldInflat(bytes[index:], &cf)
	index = index + offset

	cf.MethodsCount = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	offset, _ = methodInflat(bytes[index:], &cf)
	index = index + offset

	cf.AttributesCount = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	offset, _ = attributeInflat(bytes[index:], &cf)
	index = index + offset

	return &cf, index == len(bytes)
}

func attributeInflat(b []byte, cf *ClassFile) (int, error) {
	offset := 0
	for i := uint16(0); i < cf.AttributesCount; i++ {
		m, s := AllocAttribute(b[offset:], cf)
		cf.Attributes = append(cf.Attributes, m)
		offset += s
	}
	return offset, nil
}

func methodInflat(b []byte, cf *ClassFile) (int, error) {
	offset := 0
	for i := uint16(0); i < cf.MethodsCount; i++ {
		m, s := AllocMethod(b[offset:], cf)
		cf.Methods = append(cf.Methods, m)
		offset += s
	}
	return offset, nil
}

func fieldInflat(b []byte, cf *ClassFile) (int, error) {
	offset := 0
	for i := uint16(0); i<cf.FieldsCount; i++ {
		v, size := AllocField(b[offset : offset + 2], cf)
		cf.Fields = append(cf.Fields, v)
		offset += size
	}
	return offset, nil
}

func interfaceInflat(b []byte, cf *ClassFile) (int, error) {
	offset := 0
	for i := uint16(0); i<cf.InterfacesCount; i++ {
		cf.Interfaces = append(cf.Interfaces, utils.BigEndian2Little4U2(b[offset : offset + 2]))
		offset += 2
	}
	return offset, nil
}

func constantPoolInflat(b []byte, cf *ClassFile) (int, error) {
	offset := 0
	for i:=uint16(1); i<cf.ConstantPoolCount; i++ {
		cp, size := AllocConstantPool(TagCp(b[offset]), b[offset + 1:])
		cf.ConstantPool = append(cf.ConstantPool, cp)
		offset += 1 + size
	}
	return offset, nil
}
