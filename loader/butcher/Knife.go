package butcher

import (
	"../../utils"
	"../clazz"
)

func Decoder(bytes []byte) (*clazz.ClassFile, bool) {
	var index = 0
	cf := *clazz.AllocClassFile()
	cf.Magic = utils.BigEndian2Little4U4((bytes)[index : index + 4]); index += 4
	cf.MinorVersion = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	cf.MajorVersion = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2

	cf.ConstantPoolCount = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	offset, _ := ConstantPool(bytes[index:], &cf)
	index = index + offset

	cf.AccessFlags = clazz.AccessFlags(utils.BigEndian2Little4U2((bytes)[index : index + 2])); index += 2
	cf.ThisClass = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	cf.SuperClass = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2

	cf.InterfacesCount = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	offset, _ = Interface(bytes[index:], &cf)
	index = index + offset

	cf.FieldsCount = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	offset, _ = Field(bytes[index:], &cf)
	index = index + offset

	cf.MethodsCount = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	offset, _ = Method(bytes[index:], &cf)
	index = index + offset

	cf.AttributesCount = utils.BigEndian2Little4U2((bytes)[index : index + 2]); index += 2
	offset, _ = Attribute(bytes[index:], &cf)
	index = index + offset

	return &cf, index == len(bytes)
}

func Attribute(b []byte, cf *clazz.ClassFile) (int, error) {
	offset := 0
	for i := uint16(0); i < cf.MethodsCount; i++ {
		m, s := clazz.AllocAttribute(b[offset:], cf)
		cf.Attributes = append(cf.Attributes, m)
		offset += s
	}
	return offset, nil
}

func Method(b []byte, cf *clazz.ClassFile) (int, error) {
	offset := 0
	for i := uint16(0); i < cf.MethodsCount; i++ {
		m, s := clazz.AllocMethod(b[offset:], cf)
		cf.Methods = append(cf.Methods, m)
		offset += s
	}
	return offset, nil
}

func Field(b []byte, cf *clazz.ClassFile) (int, error) {
	offset := 0
	for i := uint16(0); i<cf.FieldsCount; i++ {
		v, size := clazz.AllocField(b[offset : offset + 2], cf)
		cf.Fields = append(cf.Fields, v)
		offset += size
	}
	return offset, nil
}

func Interface(b []byte, cf *clazz.ClassFile) (int, error) {
	offset := 0
	for i := uint16(0); i<cf.InterfacesCount; i++ {
		cf.Interfaces = append(cf.Interfaces, utils.BigEndian2Little4U2(b[offset : offset + 2]))
		offset += 2
	}
	return offset, nil
}

func ConstantPool(b []byte, cf *clazz.ClassFile) (int, error) {
	offset := 0
	for i:=uint16(1); i<cf.ConstantPoolCount; i++ {
		cp, size := clazz.AllocConstantPool(clazz.TagCp(b[offset]), b[offset + 1:])
		cf.ConstantPool = append(cf.ConstantPool, cp)
		offset += 1 + size
	}
	return offset, nil
}

