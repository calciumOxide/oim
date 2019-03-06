package clazz

import "../../utils"
import (
	"./item"
	"reflect"
	"../../types"
)

type Method struct {
	AccessFlags AccessFlags
	NameIndex uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes []*Attribute //attributes 表的每一个成员的值必须是 attribute(§4.7)结构，一个方法可以有 任意个与之相关的属性
	/*
	本规范所定义的 method_info 结构中，属性表可出现的成员有:
		Code(§4.7.3)， Exceptions(§4.7.5)，Synthetic(§4.7.8)，Signature(§4.7.9)， Deprecated(§4.7.15)，
	untimeVisibleAnnotations(§4.7.16)， RuntimeInvisibleAnnotations(§4.7.17)， RuntimeVisibleParameterAnnotations(§4.7.18)，
	RuntimeInvisibleParameterAnnotations(§4.7.19)和 AnnotationDefault(§4.7.20)结构。
	*/
	ClassFile *ClassFile
	name string
	descriptor string
}

func AllocMethod(b []byte, cf *ClassFile) (*Method, int) {
	offset := 8
	v := Method {
		AccessFlags: AccessFlags(utils.BigEndian2Little4U2(b[:2])),
		NameIndex: utils.BigEndian2Little4U2(b[2 : 4]),
		DescriptorIndex: utils.BigEndian2Little4U2(b[4 : 6]),
		AttributesCount: utils.BigEndian2Little4U2(b[6 : 8]),
	}
	for i := uint16(0); i < v.AttributesCount; i++ {
		a, s := AllocAttribute(b[offset:], cf)
		v.Attributes = append(v.Attributes, a)
		offset += s
	}
	return &v, offset
}

func (s *Method) GetAttribute(t AttributeName) (*Attribute, error) {
	for _, a := range s.Attributes {
		if a.AttributeName == t {
			return a, nil
		}
	}
	return nil, nil
}

func (s *Method) GetName() string {
	if len(s.name) == 0{
		pool, _ := s.ClassFile.GetConstant(s.NameIndex)
		s.name = pool.Info.(*item.Utf8).Str
	}
	return s.name
}

func (s *Method) GetDescriptor() string {
	if len(s.descriptor) == 0 {
		pool, _ := s.ClassFile.GetConstant(s.DescriptorIndex)
		s.descriptor = pool.Info.(*item.Utf8).Str
	}
	return s.descriptor
}

func (s *Method) IsStatic() bool {
	return s.AccessFlags & ACC_STATIC > 0
}

func (s *Method) IsAbstarct() bool {
	return s.AccessFlags & ACC_ABSTRACT > 0
}

func (s *Method) IsPrivate() bool {
	return s.AccessFlags & ACC_PRIVATE > 0
}

func (s *Method) IsProtected() bool {
	return s.AccessFlags & ACC_PROTECTED > 0
}

func (s *Method) GetParmsType(str string) []string {
	desc := []byte(str)
	parmsType := []string{}
	if len(desc) > 0 {
		if desc[0] != '(' {
			print("decode method params err.>>>>>>>>>>>>>>>>>>>>>>>")
		}
		outer:
		for i := 1; i < len(desc); {
			types := ""
			inner:
			switch desc[i] {
			case '[':
				types += string(desc[i])
				i++
				goto inner
			case 'L':
				if len(types) == 0 || types[len(types) - 1] == '[' {
					types += string(desc[i])
				}
				i++
				if desc[i] == ';' {
					i++
					break
				}
				types += string(desc[i])
				desc[i] = 'L'
				goto inner
			case ')':
				i = len(desc)
				break outer
			default:
				types = string(desc[i])
				i++
				break
			}
			parmsType = append(parmsType, types)
		}
	}
	return parmsType
}
func (s *Method) CheckParams(parmsType []string, args []interface{}) bool {
	if len(parmsType) != len(args) {
		return false
	}
	length := len(args)
	for i := 0; i< length; i++ {
		types := parmsType[i]
		arg := args[length - 1 - i]
		if !checkParams(types, arg) {
			return false
		}
	}
	return true
}

func checkParams(t string, arg interface{}) bool {
	of := reflect.TypeOf(arg)
	switch t[0] {
	case '[':
		return of == reflect.TypeOf(&types.Jarray{}) &&
		checkParams(t[1:], arg)
		break
	case 'L':
		return of == reflect.TypeOf(&types.Jreference{}) && 
			(arg.(*types.Jreference).ElementType.(*ClassFile) == GetClass(t[1:]))
		break
	case 'B':
		return of == reflect.TypeOf(types.Jbyte(0))
	case 'C':
		return of == reflect.TypeOf(types.Jchar(0))
	case 'D':
		return of == reflect.TypeOf(types.Jdouble(0))
	case 'F':
		return of == reflect.TypeOf(types.Jfloat(0))
	case 'I':
		return of == reflect.TypeOf(types.Jint(0))
	case 'J':
		return of == reflect.TypeOf(types.Jlong(0))
	case 'S':
		return of == reflect.TypeOf(types.Jshort(0))
	case 'Z':
		return of == reflect.TypeOf(types.Jboolean(true))
	default:
		return false
		break
	}
	return false
}


