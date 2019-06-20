package attribute

import "../../../utils"

/*
RuntimeInvisibleAnnotations 属性和 RuntimeVisibleAnnotations 属性相似，
不 同的是 RuntimeVisibleAnnotations 表示的注解不能被反射 API 访问，
除非 Java 虚拟机通 过特殊的实现相关的方式(譬如特定的命令行参数)收到才会(为反射的 API)使用这些注解。
否则，Java 虚拟机将忽略 RuntimeVisibleAnnotations 属性。

RuntimeInvisibleAnnotations 属性是一个变长属性，
位于 ClassFile(§4.1), field_info(§4.5)或 method_info(§4.6)结构的属性表中。用于保存 Java 语言中的 类、字段或方法的运行时的非可见注解。
每个 ClassFile，field_info 和 method_info 结构 最多只能含有一个 RuntimeInvisibleAnnotations 属性，
它为当前的程序元素保存所有的运 行时非可见的 Java 语言注解。
*/
type RuntimeInvisibleAnnotations struct {
	AnnotationsCount uint16
	Annotations      []*Annotation
}

func AllocRuntimeInvisibleAnnotations(b []byte) (*RuntimeInvisibleAnnotations, int) {
	offset := 2
	v := RuntimeInvisibleAnnotations{
		AnnotationsCount: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.AnnotationsCount; i++ {
		a, s := AllocAnnotation(b[offset:])
		v.Annotations = append(v.Annotations, a)
		offset += s
	}
	return &v, offset
}
