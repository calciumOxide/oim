package attribute

import "../../../utils"

type Annotation struct {
	TypeIndex uint16 //是对常量池的一个有效索引。常量池在该索引处的成 员必须是 CONSTANT_Utf8_info(§4.4.7)结构，表示一个字段描述符 这个字段描述符表示一个注解类型，和当前 annotation 结构表示的注解一致。
	ElementValuePairsCount uint16 //num_element_value_pairs 项的值给出了当前 annotation 结构表示的注 解的键值对(键值对的格式为:元素-值)的数量，即 element_value_pairs[] 数组成员数量。需要注意的是，在单独一个注解中可能含有数量最多为 65535 个键值对。
	ElementValuePairs []*ElementValuePairs //element_value_pairs[]数组的每一个成员的值对应当前 annotation 结 构表示的注解中的一个唯一的键值对。element_value_pairs 的成员包含如 下两个项。
}


func AllocAnnotation(b []byte) (*Annotation, int) {
	offset := 4
	v := Annotation {
		TypeIndex: utils.BigEndian2Little4U2(b[:2]),
		ElementValuePairsCount: utils.BigEndian2Little4U2(b[2 : 4]),
	}
	for i := uint16(0); i < v.ElementValuePairsCount; i++ {
		e, s := AllocAElementValuePairs(b[offset:])
		v.ElementValuePairs = append(v.ElementValuePairs, e)
		offset += s
	}
	return &v, offset
}

type ElementValuePairs struct {
	ElementNameIndex uint16 //element_name_index 项的值必须是对常量池的一个有效索引。常量池 在该索引处的成员必须是 CONSTANT_Utf8_info(§4.4.7)结构，表 示一个有效的字段描述符(§4.3.2)，这个字段描述符用于定义当前 element_value_pairs 的成员表示的注解的注解名。
	Value ElementValue
}


func AllocAElementValuePairs(b []byte) (*ElementValuePairs, int) {
	v := new(ElementValuePairs)
	v.ElementNameIndex = utils.BigEndian2Little4U2(b[:2])
	e, s := AllocElementValue(b[2:])
	v.Value = *e
	return v, 2 + s
}


/*
element_value 结构是一个可辨识联合体(Discriminated Union)1，用于表示“元素 -值”的键值对中的值。
它被用来描述所有注解(包括 RuntimeVisibleAnnotations、 RuntimeInvisibleAnnotations、RuntimeVisibleParameterAnnotations
和 untimeInvisibleParameterAnnotations)中涉及到的元素值。
*/
type ElementValue struct {
	Tag uint8 //tag 项表明了当前注解的元素-值对的类型。tag 值为字母"B"、"C"、"D"、"F"、"I"、 "J"、"S"和"Z"时表示的含义和章节 4.3.2 中表 4.2 所定义的一样。其余 tag 的预定 义值和对应解释在表 4.9 列出。 s -> String, e -> enum constant, c -> class, @ -> annotation type, [ -> array
	Value interface{}
	
}

func AllocElementValue(b []byte) (*ElementValue, int) {
	v := new(ElementValue)
	v.Tag = b[0]
	s := 0
	i := v.Value
	switch string(v.Tag) {
	case "e":
		s = 4
		s = 2
		i = &EnumElementValue {
			TypeNameIndex: utils.BigEndian2Little4U2(b[1:3]),
			ConstNameIndex: utils.BigEndian2Little4U2(b[3:5]),
		}
		break
	case "c":
		s = 2
		i = &ClassElementValue {
			ClassInfoIndex: utils.BigEndian2Little4U2(b[1:3]),
		}
		break
	case "@":
		a, size := AllocAnnotation(b[1:])
		i = &AnnotationElementValue{
			Annotation: *a,
		}
		s += size
		break
	case "[":
		s = 2
		a := ArrayElementValue {
			ValueCount: utils.BigEndian2Little4U2(b[1 : 3]),
		}
		for k := uint16(0); k < a.ValueCount; k++ {
			e, size := AllocElementValue(b[1+s:])
			a.ElementValue = append(a.ElementValue, e)
			s += size
		}
		i = &a
		break
	case "B":
	case "C":
	case "D":
	case "F":
	case "I":
	case "J":
	case "S":
	case "s":
		s = 2
		i = &ConstElementValue {
			ConstValueIndex: utils.BigEndian2Little4U2(b[1:3]),
		}
		break
	default:
		print("===================================================>>>> alloc element value")
		break
	}
	v.Value = i
	return v, s + 1
}

//func AllocElementValueInfo(b []byte) (*ElementValueInfo, int) {

//}

//type ElementValueInfo struct {
//
//}

type ConstElementValue struct {
	//empty ElementValueInfo
	ConstValueIndex uint16
}

type EnumElementValue struct {
	//empty ElementValueInfo
	TypeNameIndex uint16 //type_name_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的 成员必须是 CONSTANT_Utf8_info(§4.4.7)结构，表示一个有效的字段描述 符(§4.3.2)，这个字段描述符表示当前 element_value 结构所表示的枚举常量 类型的内部形式的二进制名称(§ 4.2.1)。
	ConstNameIndex uint16 //const_name_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的 成员必须是 CONSTANT_Utf8_info(§4.4.7)结构，表示一个有效的字段描述 符(§4.3.2)，这个字段描述符表示当前 element_value 结构所表示的枚举常量 类型的简单名称。
}

type ClassElementValue struct {
	//empty ElementValueInfo
	ClassInfoIndex uint16 //当 tag 项为"c"时，class_info_index 项才会被使用。class_info_index 项 的值必须是对常量池的一个有效索引。常量池在该索引处的成员必须是 CONSTANT_Utf8_info(§4.4.7)结构，表示返回描述符(§4.3.3)的类型， 这个类型由当前 element_value 结构所表示的类型决定(譬如:"V"表示 Void， "Ljava/lang/Object;"表示类 java.lang.Object 等)。
}

type AnnotationElementValue struct {
	Annotation Annotation //当 tag 项为"@"时，annotation_value 项才会被使用。这时 element_value 结构表示一个内部的注解(Nested Annotation)。
}
type ArrayElementValue struct {//当 tag 项为"["时，array_value 项才会被使用。
	ValueCount uint16 //num_values 项的值给定了当前 element_value 结构表示的数组类型的成员的数量。
	ElementValue []*ElementValue //values 的每个成员的值都给指明了当前 element_value 结构所表示的数组 类型的一个元素值
}

