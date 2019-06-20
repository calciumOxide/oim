package attribute

import "../../../utils"

/*
RuntimeVisibleAnnotations 属性是可变长属性，位于 ClassFile(§4.1)， field_info(§4.5)或 method_info(§4.6)结构的属性表中。
RuntimeVisibleAnnotations 用于保存 Java 语言中的类、字段或方法的运行时的可见注解 (Annotations)。
每个 ClassFile，field_info 和 method_info 结构最多只能含有一个 RuntimeVisibleAnnotations属性为当前的程序元素保存所有的运行时可见的 Java 语言注解。
Java 虚拟机必须支持这些注解可被反射的 API 使用它们。
*/
type RuntimeVisibleAnnotations struct {
	AnnotationCount uint16
	Annotations     []*Annotation
}

func AllocRuntimeVisibleAnnotations(b []byte) (*RuntimeVisibleAnnotations, int) {
	offset := 2
	v := RuntimeVisibleAnnotations{
		AnnotationCount: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.AnnotationCount; i++ {
		a, s := AllocAnnotation(b[offset:])
		v.Annotations = append(v.Annotations, a)
		offset += s
	}
	return &v, offset
}
