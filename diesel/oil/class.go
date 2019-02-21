package oil

import (
	"../../loader"
	 "../../loader/binary"
	 "../../loader/binary/item"
	"regexp"
)

type Class struct {

	ClassName string
	ClassFile *binary.ClassFile

	SuperClass *Class
	Interfaces []*Class
	ClassLoader *Class

	IsFinal bool
	IsStatic bool
	IsPublic bool
	IsProtected bool
	IsPrivate bool
	IsDefault bool
	IsAbstract bool
	IsInterface bool

	Fields []*Field
	Methods []*Method
}

type Field struct {

	ClassType *Class

	IsFinal bool
	IsStatic bool
	IsPublic bool
	IsProtected bool
	IsPrivate bool
	IsDefault bool
	IsVolatile bool
	IsEnum bool
	IsTransient bool

	Name string
	Descriptor string

	BField *binary.Field

}

type Method struct {

	IsFinal bool
	IsStatic bool
	IsPublic bool
	IsProtected bool
	IsPrivate bool
	IsDefault bool
	IsAbstract bool
	IsVarargs bool
	IsNative bool
	IsSynchronized bool

	Name string
	Descriptor string
	AttributesCount uint16

	BMethod *binary.Method
}

var CLASS_MAP = make(map[string] *Class)

func GetClass(className string) *Class {
	class := CLASS_MAP[className]
	if class == nil {
		classMap := make(map[string] *Class)
		cf := loader.Loader(className)
		if cf == nil {
			panic("class [" + className + "] loader err.")
		}
		class = &Class{
			ClassName: className,
			ClassFile: cf,
		}
		class.inflation()
		class.ClassName = className
		classMap[className] = class
		if class.SuperClass != nil {
			if CLASS_MAP[class.SuperClass.ClassName] == nil {
				GetClass(class.SuperClass.ClassName)
			}
		}
		for k, v := range classMap {
			CLASS_MAP[k] = v
		}
	}
	return class
}

func (s *Class)inflation() {
	cf := s.ClassFile

	superClassBin, _ := cf.GetConstant(cf.SuperClass)
	superClassNameBin, _ := cf.GetConstant(superClassBin.Info.(*item.Class).NameIndex)
	s.SuperClass = GetClass(superClassNameBin.Info.(*item.Utf8).Str)

	if cf.InterfacesCount > 0 {
		s.Interfaces = make([]*Class, cf.InterfacesCount)
		for i := 0; i < len(cf.Interfaces); i++ {
			iClassBin, _ := cf.GetConstant(cf.Interfaces[i])
			iClassNameBin, _ := cf.GetConstant(iClassBin.Info.(*item.Class).NameIndex)
			iClass := GetClass(iClassNameBin.Info.(*item.Utf8).Str)
			s.Interfaces = append(s.Interfaces, iClass)
		}
	}

	if cf.FieldsCount > 0 {
		s.Fields = make([]*Field, cf.FieldsCount)
		for i := uint16(0); i < cf.FieldsCount; i++ {
			fieldBin := cf.Fields[i]
			fieldNameBin, _ := cf.GetConstant(fieldBin.NameIndex)
			fieldDescBin, _ := cf.GetConstant(fieldBin.DescriptorIndex)

			field := &Field{
				BField:     fieldBin,
				IsFinal: fieldBin.AccessFlags & binary.ACC_FINAL == 1,
				IsStatic: fieldBin.AccessFlags & binary.ACC_STATIC == 1,
				IsPublic: fieldBin.AccessFlags & binary.ACC_PUBLIC == 1,
				IsProtected: fieldBin.AccessFlags & binary.ACC_PROTECTED == 1,
				IsPrivate: fieldBin.AccessFlags & binary.ACC_PRIVATE == 1,
				IsVolatile: fieldBin.AccessFlags & binary.ACC_VOLATILE == 1,
				IsEnum: fieldBin.AccessFlags & binary.ACC_ENUM == 1,
				IsTransient: fieldBin.AccessFlags & binary.ACC_TRANSIENT == 1,
				Name:       fieldNameBin.Info.(*item.Utf8).Str,
				Descriptor: fieldDescBin.Info.(*item.Utf8).Str,
			}
			reg := regexp.MustCompile(`^\[*L$`)
			if reg.MatchString(field.Descriptor) {
				fClassName := reg.ReplaceAllString(field.Descriptor, "")
				fClass := GetClass(fClassName[:len(fClassName)-1])
				field.ClassType = fClass
			}
			s.Fields = append(s.Fields, field)
		}
	}
	
	if cf.MethodsCount > 0 {
		s.Methods = make([]*Method, cf.MethodsCount)
		for i := uint16(0); i < cf.MethodsCount; i++ {
			methodBin := cf.Methods[i]
			methodNameBin, _ := cf.GetConstant(methodBin.NameIndex)
			methodDescBin, _ := cf.GetConstant(methodBin.DescriptorIndex)
			
			method := &Method{
				BMethod: methodBin,
				IsFinal: methodBin.AccessFlags & binary.ACC_FINAL == 1,
				IsStatic: methodBin.AccessFlags & binary.ACC_STATIC == 1,
				IsPublic: methodBin.AccessFlags & binary.ACC_PUBLIC == 1,
				IsProtected: methodBin.AccessFlags & binary.ACC_PROTECTED == 1,
				IsPrivate: methodBin.AccessFlags & binary.ACC_PRIVATE == 1,
				IsAbstract: methodBin.AccessFlags & binary.ACC_ABSTRACT == 1,
				IsVarargs: methodBin.AccessFlags & binary.ACC_VARARGS == 1,
				IsNative: methodBin.AccessFlags & binary.ACC_NATIVE == 1,
				IsSynchronized: methodBin.AccessFlags & binary.ACC_SYNCHRONIZED == 1,
				Name:       methodNameBin.Info.(*item.Utf8).Str,
				Descriptor: methodDescBin.Info.(*item.Utf8).Str,
			}

			s.Methods = append(s.Methods, method)
		}
	}
}


func (s *Class) GetMethod(name string, descriptor string) *Method {
	for i := 0; i < len(s.Methods); i++ {
		method := s.Methods[i]
		if !method.IsStatic && method.Name == name && method.Descriptor == descriptor {
			return method
		}
	}
	if s.SuperClass == nil {
		return nil
	}
	return s.SuperClass.GetMethod(name, descriptor)
}