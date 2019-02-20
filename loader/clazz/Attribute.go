package clazz

import "../../utils"
import "./item"
import "./attribute"

type Attribute struct {
	NameIndex       uint16
	AttributeLength uint32
	AttributeItem   interface{}
	AttributeName   AttributeName
}

type AttributeName string

const (
	CONSTANT_VALUE_ATTR                          = "ConstantValue"
	CODE_ATTR                                    = "Code"
	STACK_MAP_TABLE_ATTR                         = "StackMapTable"
	EXCEPTIONS_ATTR                              = "Exceptions"
	INNER_CLASSES_ATTR                           = "InnerClasses"
	ENCLOSING_METHOD_ATTR                        = "EnclosingMethod"
	SYNTHETIC_ATTR                               = "Synthetic"
	SIGNATURE_ATTR                               = "Signature"
	SOURCE_FILE_ATTR                             = "SourceFile"
	SOURCE_DEBUG_EXTENSION_ATTR                  = "SourceDebugExtension"
	LINE_NUMBER_TABLE_ATTR                       = "LineNumberTable"
	LOCAL_VARIABLE_TABLE_ATTR                    = "LocalVariableTable"
	LOCAL_VARIABLE_TYPE_TABLE_ATTR               = "LocalVariableTypeTable"
	DEPRECATED_ATTR                              = "Deprecated"
	RUNTIME_VISIBLE_ANNOTATIONS_ATTR             = "RuntimeVisibleAnnotations"
	RUNTIME_INVISIBLE_ANNOTATIONS_ATTR           = "RuntimeInvisibleAnnotations"
	RUNTIME_VISIBLE_PARAMETER_ANNOTATIONS_ATTR   = "RuntimeVisibleParameterAnnotations"
	RUNTIME_INVISIBLE_PARAMETER_ANNOTATIONS_ATTR = "RuntimeInvisibleParameterAnnotations"
	ANNOTATION_DEFAULT_ATTR                      = "AnnotationDefault"
	BOOTSTRAP_METHODS_ATTR                       = "BootstrapMethods"
)

func AllocAttribute(b []byte, cf *ClassFile) (*Attribute, int) {
	v := Attribute{
		NameIndex:       utils.BigEndian2Little4U2(b[:2]),
		AttributeLength: utils.BigEndian2Little4U4(b[2:6]),
	}
	b = b[6: v.AttributeLength + 6]
	a := v.AttributeItem
	constant, _ := cf.GetConstant(v.NameIndex)
	v.AttributeName = AttributeName(constant.Info.(*item.Utf8).Str)
	switch v.AttributeName {
	case CONSTANT_VALUE_ATTR:
		a, _ = attribute.AllocConstantValue(b)
		break
	case CODE_ATTR:
		a, _ = AllocCodes(b, cf)
		break
	case STACK_MAP_TABLE_ATTR:
		a, _ = attribute.AllocStackMapTable(b)
		break
	case EXCEPTIONS_ATTR:
		a, _ = attribute.AllocExceptions(b)
		break
	case INNER_CLASSES_ATTR:
		a, _ = attribute.AllocInnerClasses(b)
		break
	case ENCLOSING_METHOD_ATTR:
		a, _ = attribute.AllocEnclosingMethod(b)
		break
	case SYNTHETIC_ATTR:
		a, _ = attribute.AllocSynthetic(b)
		break
	case SIGNATURE_ATTR:
		a, _ = attribute.AllocSignature(b)
		break
	case SOURCE_FILE_ATTR:
		a, _ = attribute.AllocSourceFile(b)
		break
	case SOURCE_DEBUG_EXTENSION_ATTR:
		a = &attribute.SourceDebugExtension{
			DebugExtesion: b,
		}
		break
	case LINE_NUMBER_TABLE_ATTR:
		a, _ = attribute.AllocLineNumberTable(b)
		break
	case LOCAL_VARIABLE_TABLE_ATTR:
		a, _ = attribute.AllocLocalVariableTable(b)
		break
	case LOCAL_VARIABLE_TYPE_TABLE_ATTR:
		a, _ = attribute.AllocLocalVariableTypeTable(b)
		break
	case DEPRECATED_ATTR:
		a = nil
		break
	case RUNTIME_VISIBLE_ANNOTATIONS_ATTR:
		a, _ = attribute.AllocRuntimeVisibleAnnotations(b)
		break
	case RUNTIME_INVISIBLE_ANNOTATIONS_ATTR:
		a, _ = attribute.AllocRuntimeInvisibleAnnotations(b)
		break
	case RUNTIME_VISIBLE_PARAMETER_ANNOTATIONS_ATTR:
		a, _ = attribute.AllocRuntimeVisibleParameterAnnotations(b)
		break
	case RUNTIME_INVISIBLE_PARAMETER_ANNOTATIONS_ATTR:
		a, _ = attribute.AllocRuntimeInvisibleParameterAnnotations(b)
		break
	case ANNOTATION_DEFAULT_ATTR:
		a, _ = attribute.AllocAnnotationDefault(b)
		break
	case BOOTSTRAP_METHODS_ATTR:
		a, _ = attribute.AllocBootstrapMethods(b)
		break
	default:
		print("============================================================>>>>>>")
	}
	v.AttributeItem = a
	return &v, int(v.AttributeLength) + 6
}

func AllocCodes(b []byte, cf *ClassFile) (*attribute.Codes, int) {
	offset := 8
	v := attribute.Codes{
		MaxStack:   utils.BigEndian2Little4U2(b[:2]),
		MaxLocal:   utils.BigEndian2Little4U2(b[2:4]),
		CodeLength: utils.BigEndian2Little4U4(b[4:8]),
	}
	offset += int(v.CodeLength)
	v.Code = b[8 : offset]
	v.ExceptTableLength = utils.BigEndian2Little4U2(b[offset : offset + 2])
	offset += 2
	for i := uint16(0); i < v.ExceptTableLength; i++ {
		t, size := attribute.AllocExceptTable(b[offset:])
		v.ExceptTable = append(v.ExceptTable, t)
		offset += size
	}
	v.AttributeCount = utils.BigEndian2Little4U2(b[offset : offset + 2])
	offset += 2
	for i := uint16(0); i < v.AttributeCount; i++ {
		a, size := AllocAttribute(b[offset:], cf)
		v.Attributes = append(v.Attributes, a)
		offset += size
	}
	return &v, offset
}