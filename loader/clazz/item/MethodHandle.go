package item

import "../../../utils"

type MethodHandle struct {
	ReferenceKind ReferenceKind //reference_kind 项的值必须在 1 至 9 之间(包括 1 和 9)，它决定了方法句柄的类型。方法句柄类型的值表示方法句柄的字节码行为(Bytecode Behavior §5.4.3.5)
	ReferenceIndex uint16
}

func AllocMethodHandleItem(b []byte) (*MethodHandle, int) {
	return &MethodHandle{
		ReferenceKind: ReferenceKind(b[0]),
		ReferenceIndex: utils.BigEndian2Little4U2(b[1:3]),
	}, 3
}

type ReferenceKind uint8

const (
	//如果 reference_kind 项的值为 1(REF_getField)、2(REF_getStatic)、3 (REF_putField)或 4(REF_putStatic)，
	// 那么常量池在 reference_index 索引处的项必须是 CONSTANT_Fieldref_info(§4.4.2)结构，表示由一个字 段创建的方法句柄
	GET_FIELD ReferenceKind = iota
	GET_STATIC
	PUT_FIELD
	PUT_STATIC

	//如果 reference_kind 项的值是 5(REF_invokeVirtual)、6 (REF_invokeStatic)、7(REF_invokeSpecial)或 8 (REF_newInvokeSpecial)，
	// 那么常量池在 reference_index 索引处的项必须 是 CONSTANT_Methodref_info(§4.4.2)结构，表示由类的方法或构造函数 创建的方法句柄
	INVOKE_VIRTUAL
	INVOKE_STATIC
	INVOKE_SPECIAL
	NEW_INVOKE_SPECIAL

	//如果 reference_kind 项的值是 9(REF_invokeInterface)，
	// 那么常量池在 reference_index 索引处的项必须是 CONSTANT_InterfaceMethodref_info (§4.4.2)结构，表示由接口方法创建的方法句柄
	INVOKE_INTERFACE

)

