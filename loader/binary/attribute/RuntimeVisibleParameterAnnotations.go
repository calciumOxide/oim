package attribute

import "../../../utils"

/*
RuntimeVisibleParameterAnnotations 属性是一个变长属性，位于 method_info(§4.6)结构的属性表中。
用于保存对应方法的参数的所有运行时可见 Java 语言注解。
每个method_info 结构最多只能包含一个 RuntimeVisibleParameterAnnotations 属性，
用于 保存当前方法的参数的所有可见的 Java 语言注解。Java 虚拟机必须保证这些注解可被反射的 API 使用。
*/
type RuntimeVisibleParameterAnnotations struct {
	ParameterCount      uint16
	ParameterAnnotation []*ParameterAnnotation
}

type ParameterAnnotation struct {
	AnnotationCount uint16
	Annotation      []*Annotation
}

/*
RuntimeInvisibleParameterAnnotations 属性和 RuntimeVisibleParameterAnnotations 属性类似，
区别是 RuntimeInvisibleParameterAnnotations 属性表示的注解不能被反射的 API 访问，
除非 Java 虚拟机通过特殊的实现相关的方式(譬如特定的命令行参数)收到才会(为反射的 API)使 用这些注解。
否则，Java 虚拟机将忽略 RuntimeInvisibleParameterAnnotations 属性。
RuntimeInvisibleParameterAnnotations 属性是一个变长属性，位于 method_info (§4.6)结构的属性表中。
用于保存对应方法的参数的所有运行时非可见的 Java 语言注解。
每 个 method_info 结构最多只能含有一个 RuntimeInvisibleParameterAnnotations 属性 用于保存当前 Java 语言中的程序元素的所有运行时非可见注解。
*/
type RuntimeInvisibleParameterAnnotations struct {
	ParameterCount      uint16
	ParameterAnnotation []*ParameterAnnotation
}

func AllocRuntimeVisibleParameterAnnotations(b []byte) (*RuntimeVisibleParameterAnnotations, int) {
	offset := 2
	v := RuntimeVisibleParameterAnnotations{
		ParameterCount: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.ParameterCount; i++ {
		a, s := AllocParameterAnnotation(b[offset:])
		v.ParameterAnnotation = append(v.ParameterAnnotation, a)
		offset += s
	}
	return &v, offset
}

func AllocRuntimeInvisibleParameterAnnotations(b []byte) (*RuntimeInvisibleParameterAnnotations, int) {
	offset := 2
	v := RuntimeInvisibleParameterAnnotations{
		ParameterCount: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.ParameterCount; i++ {
		a, s := AllocParameterAnnotation(b[offset:])
		v.ParameterAnnotation = append(v.ParameterAnnotation, a)
		offset += s
	}
	return &v, offset
}

func AllocParameterAnnotation(b []byte) (*ParameterAnnotation, int) {
	offset := 2
	v := ParameterAnnotation{
		AnnotationCount: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.AnnotationCount; i++ {
		a, s := AllocAnnotation(b[offset:])
		v.Annotation = append(v.Annotation, a)
		offset += s
	}
	return &v, offset
}
