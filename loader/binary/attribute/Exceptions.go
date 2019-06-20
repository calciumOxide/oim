package attribute

import "../../../utils"

/*
Exceptions 属性是一个变长属性，它位于 method_info(§4.6)结构的属性表中。
Exceptions 属性指出了一个方法需要检查的可能抛出的异常。一个 method_info 结构中最多
只能有一个 Exceptions 属性
==== 方法签名的 throws
*/
type Exceptions struct {
	ExceptionCount uint16
	ExceptionIndex []uint16 //exception_index_table[]数组的每个成员的值都必须是对常量池的有效索引。常量 池在这些索引处的成员必须都是 CONSTANT_Class_info(§4.4.1)结构，表示这个 方法声明要抛出的异常的类的类型。
}

/*
一个方法如果要抛出异常，必须至少满足下列三个条件中的一个:
 要抛出的是 RuntimeException 或其子类的实例。
 要抛出的是 Error 或其子类的实例。
 要抛出的是在 exception_index_table[]数组中申明的异常类或其子类的实例。 这些要求没有在 Java 虚拟机中进行强制检查，它们只在编译时进行强制检查
*/

func AllocExceptions(b []byte) (*Exceptions, int) {
	offset := 2
	v := Exceptions{
		ExceptionCount: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.ExceptionCount; i++ {
		v.ExceptionIndex = append(v.ExceptionIndex, utils.BigEndian2Little4U2(b[offset:offset+2]))
		offset += 2
	}
	return &v, offset
}
