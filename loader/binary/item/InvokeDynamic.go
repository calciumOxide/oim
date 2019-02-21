package item

import "../../../utils"

/*CONSTANT_InvokeDynamic_info 用于表示 invokedynamic 指令所使用到的引导方法 (Bootstrap Method)、
引导方法使用到动态调用名称(Dynamic Invocation Name)、
参 数和请求返回类型、以及可以选择性的附加被称为静态参数(Static Arguments)的常量序列*/

type InvokeDynamic struct {
	BootstrapMethodsAttributeIndex uint16 //对当前 Class 文件中引导方法表(§ 4.7.21)的 bootstrap_methods[]数组的有效索引
	NameAndTypeIndex uint16 //对当前常量池的有效索引，常量池在该索引处的 项必须是 CONSTANT_NameAndType_info(§4.4.6)结构，表示方法名和方法描述 符(§4.3.3)。
}

func AllocInvokeDynamicItem(b []byte) (*InvokeDynamic, int) {
	return &InvokeDynamic {
		BootstrapMethodsAttributeIndex: utils.BigEndian2Little4U2(b[:2]),
		NameAndTypeIndex: utils.BigEndian2Little4U2(b[2:4]),
	}, 4
}

