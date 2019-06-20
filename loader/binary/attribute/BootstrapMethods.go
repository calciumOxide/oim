package attribute

import "../../../utils"

/*
BootstrapMethods 属性是一个变长属性，位于 ClassFile(§4.1)结构的属性表中。 用于保存 invokedynamic 指令引用的引导方法限定符。
如果某个 ClassFile 结构的常量池中有至少一个 CONSTANT_InvokeDynamic_info(§ 4.4.10)项，
那么这个 ClassFile 结构的属性表中必须有一个明确的 BootstrapMethods 属 性。ClassFile 结构的属性表中最多只能有一个 BootstrapMethods 属性。
*/
type BootstrapMethods struct {
	BootstrapMethodCount uint16
	BootstrapMethods     []*BootstrapMethodInfo //bootstrap_methods[]数组的每个成员包含一个指向 CONSTANT_MethodHandle 结 构的索引值，它代表了一个引导方法。还包含了这个引导方法静态参数的序列(可能为空)。
}

type BootstrapMethodInfo struct {
	BootstrapMethodIndex uint16 //bootstrap_method_ref 项的值必须是一个对常量池的有效索引。常量池在该索 引处的值必须是一个 CONSTANT_MethodHandle_info 结构。
	//注意:此 CONSTANT_MethodHandle_info 结构的 reference_kind 项应为值 6 (REF_invokeStatic)或 8(REF_newInvokeSpecial)(§5.4.3.5)，否则 在 invokedynamic 指令解析调用点限定符时，引导方法会执行失败。
	BootstrapMethodArgumentCount uint16
	BootstrapMethodArguments     []uint16
	//bootstrap_arguments[]数组的每个成员必须是一个对常量池的有效索引。常量 池在该索引出必须是下列结构之一:
	//CONSTANT_String_info,CONSTANT_Class_info、 CONSTANT_Integer_info,CONSTANT_Long_info、
	// CONSTANT_Float_info,CONSTANT_Double_info、CONSTANT_MethodHandle_info 或 CONSTANT_MethodType_info。
}

func AllocBootstrapMethods(b []byte) (*BootstrapMethods, int) {
	offset := 2
	v := BootstrapMethods{
		BootstrapMethodCount: utils.BigEndian2Little4U2(b[:2]),
	}
	for i := uint16(0); i < v.BootstrapMethodCount; i++ {
		b, s := AllocBootstrapMethodInfo(b[offset:])
		v.BootstrapMethods = append(v.BootstrapMethods, b)
		offset += s
	}
	return &v, offset
}

func AllocBootstrapMethodInfo(b []byte) (*BootstrapMethodInfo, int) {
	offset := 4
	v := BootstrapMethodInfo{
		BootstrapMethodIndex:         utils.BigEndian2Little4U2(b[:2]),
		BootstrapMethodArgumentCount: utils.BigEndian2Little4U2(b[2:4]),
	}
	for i := uint16(0); i < v.BootstrapMethodArgumentCount; i++ {
		v.BootstrapMethodArguments = append(v.BootstrapMethodArguments, utils.BigEndian2Little4U2(b[offset:offset+2]))
		offset += 2
	}
	return &v, offset
}
