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

	constantPool *constantPool
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
			constantPool: &constantPool{

			},
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
	superClassNameBin, _ := cf.GetConstant(superClassBin.Info.(*item.ClassItemBin).NameIndex)
	s.SuperClass = GetClass(superClassNameBin.Info.(*item.Utf8ItemBin).Str)

	if cf.InterfacesCount > 0 {
		s.Interfaces = make([]*Class, cf.InterfacesCount)
		for i := 0; i < len(cf.Interfaces); i++ {
			iClassBin, _ := cf.GetConstant(cf.Interfaces[i])
			iClassNameBin, _ := cf.GetConstant(iClassBin.Info.(*item.ClassItemBin).NameIndex)
			iClass := GetClass(iClassNameBin.Info.(*item.Utf8ItemBin).Str)
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
				Name:       fieldNameBin.Info.(*item.Utf8ItemBin).Str,
				Descriptor: fieldDescBin.Info.(*item.Utf8ItemBin).Str,
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
				Name:       methodNameBin.Info.(*item.Utf8ItemBin).Str,
				Descriptor: methodDescBin.Info.(*item.Utf8ItemBin).Str,
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


type constantPool struct {
	cp 						   []*binary.ConstantPool
	ClassItemPool              map[uint16] *ClassItem
	DoubleItemPool             map[uint16] *DoubleItem
	FieldRefItemPool           map[uint16] *FieldRefItem
	FloatItemPool              map[uint16] *FloatItem
	IntegerItemPool            map[uint16] *IntegerItem
	InterfaceMethodRefItemPool map[uint16] *InterfaceMethodRefItem
	InvokeDynamicItemPool      map[uint16] *InvokeDynamicItem
	LongItemPool               map[uint16] *LongItem
	MethodHandleItemPool       map[uint16] *MethodHandleItem
	MethodRefItemPool          map[uint16] *MethodRefItem
	MethodTypeItemPool         map[uint16] *MethodTypeItem
	NameAndTypeItemPool        map[uint16] *NameAndTypeItem
	StringItemPool             map[uint16] *StringItem
	Utf8ItemPool               map[uint16] *Utf8Item
}

func (s *constantPool)inflateConstantPool() {
	for i := uint16(0); i< uint16(len(s.cp)); i++ {

	}

}

func (s *constantPool)GetUtf8Item(i uint16) *Utf8Item {
	v := s.Utf8ItemPool[i]
	if v == nil {
		s.inflateConstantPoolIndex(i)
	}
	return s.Utf8ItemPool[i]
}

func (s *constantPool)GetClassItem(i uint16) *ClassItem {
	v := s.Utf8ItemPool[i]
	if v == nil {
		s.inflateConstantPoolIndex(i)
	}
	return s.ClassItemPool[i]
}

func (s *constantPool)GetNameAndTypeItem(i uint16) *NameAndTypeItem {
	v := s.NameAndTypeItemPool[i]
	if v == nil {
		s.inflateConstantPoolIndex(i)
	}
	return s.NameAndTypeItemPool[i]
}

func (s *constantPool)GetMethodRefItem(i uint16) *MethodRefItem {
	v := s.MethodRefItemPool[i]
	if v == nil {
		s.inflateConstantPoolIndex(i)
	}
	return s.MethodRefItemPool[i]
}

func (s *constantPool)inflateConstantPoolIndex(i uint16) {

		it := s.cp[i]
		tag := it.Tag
		switch tag {
			case binary.CLASS:
				bin := it.Info.(*item.ClassItemBin)
				classItem := &ClassItem{
					classItemBin: bin,
				}
				classItem.Name = s.GetUtf8Item(bin.NameIndex).Str
				s.ClassItemPool[i] = classItem
				break
			case binary.FIELDREF:
				bin := it.Info.(*item.FieldRefItemBin)
				classItem := s.GetClassItem(bin.ClassIndex)
				nameAndTypeItem := s.GetNameAndTypeItem(bin.NameAndTypeIndex)
				this := &FieldRefItem{
					fieldRefItemBin: bin,
					ClassName: classItem.Name,
					Name: nameAndTypeItem.Name,
					Descriptor: nameAndTypeItem.Descriptor,
				}
				s.FieldRefItemPool[i] = this
				break
			case binary.METHODREF:
				bin := it.Info.(*item.MethodRefItemBin)
				classItem := s.GetClassItem(bin.ClassIndex)
				nameAndTypeItem := s.GetNameAndTypeItem(bin.NameAndTypeIndex)
				this := &MethodRefItem{
					methodRefItemBin: bin,
					ClassName: classItem.Name,
					Name: nameAndTypeItem.Name,
					Descriptor: nameAndTypeItem.Descriptor,
				}
				s.MethodRefItemPool[i] = this
				break
			case binary.INTERFACE_METHODREF:
				bin := it.Info.(*item.InterfaceMethodRefItemBin)
				classItem := s.GetClassItem(bin.ClassIndex)
				nameAndTypeItem := s.GetNameAndTypeItem(bin.NameAndTypeIndex)
				this := &InterfaceMethodRefItem{
					interfaceMethodRefItemBin: bin,
					ClassName: classItem.Name,
					Name: nameAndTypeItem.Name,
					Descriptor: nameAndTypeItem.Descriptor,
				}
				s.InterfaceMethodRefItemPool[i] = this
				break
			case binary.STRING:
				bin := it.Info.(*item.StringItemBin)
				this := &StringItem{
					stringItemBin: bin,
					String: s.GetUtf8Item(bin.StringIndex).Str,
				}
				s.StringItemPool[i] = this
				break
			case binary.INTEGER:
				bin := it.Info.(*item.IntegerItemBin)
				this := &IntegerItem{
					integerItemBin: bin,
					Value: bin.Value,
				}
				s.IntegerItemPool[i] = this
				break
			case binary.FLOAT:
				bin := it.Info.(*item.FloatItemBin)
				this := &FloatItem{
					floatItemBin: bin,
					Value: bin.Value,
					Overflow: bin.Overflow,
					Underflow: bin.Underflow,
					NaN: bin.NaN,
				}
				s.FloatItemPool[i] = this
				break
			case binary.LONG:
				bin := it.Info.(*item.LongItemBin)
				this := &LongItem{
					longItemBin: bin,
					Value: bin.Value,
				}
				s.LongItemPool[i] = this
				break
			case binary.DOUBLE:
				bin := it.Info.(*item.DoubleItemBin)
				this := &DoubleItem{
					doubleItemBin: bin,
					Value: bin.Value,
					Overflow: bin.Overflow,
					Underflow: bin.Underflow,
					NaN: bin.NaN,
				}
				s.DoubleItemPool[i] = this
				break
			case binary.NAME_AND_TYPE:
				bin := it.Info.(*item.NameAndTypeItemBin)
				this := &NameAndTypeItem{
					nameAndTypeItemBin: bin,
					Name: s.GetUtf8Item(bin.NameIndex).Str,
					Descriptor: s.GetUtf8Item(bin.DescriptorIndex).Str,
				}
				s.NameAndTypeItemPool[i] = this
				break
			case binary.UTF8:
				bin := it.Info.(*item.Utf8ItemBin)
				this := &Utf8Item{
					utf8ItemBin: bin,
					Str: bin.Str,
				}
				s.Utf8ItemPool[i] = this
				break
			case binary.METHOD_HANDLE:
				bin := it.Info.(*item.MethodHandleItemBin)
				this := &MethodHandleItem {
					methodHandleItemBin: bin,
					MethodRefItem: s.GetMethodRefItem(bin.ReferenceIndex),
				}
				s.MethodHandleItemPool[i] = this
				break
			case binary.METHOD_TYPE:
				bin := it.Info.(*item.MethodTypeItemBin)
				this := &MethodTypeItem {
					methodTypeItemBin: bin,
					Descriptor: s.GetUtf8Item(bin.DescriptorIndex).Str,
				}
				s.MethodTypeItemPool[i] = this
				break
			case binary.INVOKE_DYNAMIC:
				bin := it.Info.(*item.InvokeDynamicItemBin)
				nameAndTypeItem := s.GetNameAndTypeItem(bin.NameAndTypeIndex)
				this := &InvokeDynamicItem {
					invokeDynamicItemBin: bin,
					//TODO
					BootstrapMethodsAttributeIndex: bin.BootstrapMethodsAttributeIndex,
					Name: nameAndTypeItem.Name,
					Descriptor: nameAndTypeItem.Descriptor,
				}
				s.InvokeDynamicItemPool[i] = this
				break
		default:
			panic("class cp err >>>>>>>")
			break
		}

}

type ClassItem struct {
	classItemBin *item.ClassItemBin
	Name string
}

type DoubleItem struct {
	doubleItemBin *item.DoubleItemBin
	Value     float64
	Overflow  bool
	Underflow bool
	NaN       bool
}

type FieldRefItem struct {
	fieldRefItemBin *item.FieldRefItemBin
	ClassName       string
	Name       		string
	Descriptor      string
}

type FloatItem struct {
	floatItemBin *item.FloatItemBin
	Value float32
	Overflow  bool
	Underflow bool
	NaN       bool
}

type IntegerItem struct {
	integerItemBin *item.IntegerItemBin
	Value uint32 //按照 Big-Endian 的顺序存储
}


type InterfaceMethodRefItem struct {
	interfaceMethodRefItemBin *item.InterfaceMethodRefItemBin
	ClassName       string
	Name       		string
	Descriptor      string
}

type InvokeDynamicItem struct {
	invokeDynamicItemBin *item.InvokeDynamicItemBin
	BootstrapMethodsAttributeIndex uint16 //对当前 Class 文件中引导方法表(§ 4.7.21)的 bootstrap_methods[]数组的有效索引
	Name            string
	Descriptor      string
}

type LongItem struct {
	longItemBin *item.LongItemBin
	Value int64
}

type MethodHandleItem struct {
	methodHandleItemBin *item.MethodHandleItemBin
	ReferenceKind item.ReferenceKind //reference_kind 项的值必须在 1 至 9 之间(包括 1 和 9)，它决定了方法句柄的类型。方法句柄类型的值表示方法句柄的字节码行为(Bytecode Behavior §5.4.3.5)
	MethodRefItem *MethodRefItem
}

type MethodRefItem struct {
	methodRefItemBin *item.MethodRefItemBin
	ClassName       string
	Name       		string
	Descriptor      string
}

type MethodTypeItem struct {
	methodTypeItemBin *item.MethodTypeItemBin
	Descriptor string
}

type NameAndTypeItem struct {
	nameAndTypeItemBin *item.NameAndTypeItemBin
	Name            string
	Descriptor      string

	//ClassFile		binary.ClassFile
}

type StringItem struct {
	stringItemBin *item.StringItemBin
	String string
}

type Utf8Item struct {
	utf8ItemBin *item.Utf8ItemBin
	Str string
}

