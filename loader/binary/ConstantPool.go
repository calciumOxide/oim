package binary

import "./item"

type ConstantPool struct {
	Tag  TagCp
	Info interface{}
	Item interface{}
}

type TagCp uint8

const (
	CLASS               TagCp = 7
	FIELDREF            TagCp = 9
	METHODREF           TagCp = 10
	INTERFACE_METHODREF TagCp = 11
	STRING              TagCp = 8
	INTEGER             TagCp = 3
	FLOAT               TagCp = 4
	LONG                TagCp = 5
	DOUBLE              TagCp = 6
	NAME_AND_TYPE       TagCp = 12
	UTF8                TagCp = 1
	METHOD_HANDLE       TagCp = 15
	METHOD_TYPE         TagCp = 16
	INVOKE_DYNAMIC      TagCp = 18
)

func AllocConstantPool(tag TagCp, b []byte) (*ConstantPool, int) {
	cp := new(ConstantPool)
	cp.Tag = tag
	offset := 0
	switch tag {
	case CLASS:
		cp.Info, offset = item.AllocClassItem(b)
		break
	case FIELDREF:
		cp.Info, offset = item.AllocFieldRefItem(b)
		break
	case METHODREF:
		cp.Info, offset = item.AllocMethodRefItem(b)
		break
	case INTERFACE_METHODREF:
		cp.Info, offset = item.AllocInterfaceMethodRefItem(b)
		break
	case STRING:
		cp.Info, offset = item.AllocStringItem(b)
		break
	case INTEGER:
		cp.Info, offset = item.AllocIntegerItem(b)
		break
	case FLOAT:
		cp.Info, offset = item.AllocFloatItem(b)
		break
	case LONG:
		cp.Info, offset = item.AllocLongItem(b)
		break
	case DOUBLE:
		cp.Info, offset = item.AllocDoubleItem(b)
		break
	case NAME_AND_TYPE:
		cp.Info, offset = item.AllocNameAndTypeItem(b)
		break
	case UTF8:
		cp.Info, offset = item.AllocUtf8Item(b)
		break
	case METHOD_HANDLE:
		cp.Info, offset = item.AllocMethodHandleItem(b)
		break
	case METHOD_TYPE:
		cp.Info, offset = item.AllocMethodTypeItem(b)
		break
	case INVOKE_DYNAMIC:
		cp.Info, offset = item.AllocInvokeDynamicItem(b)
		break
	default:
		offset = 0
		break
	}
	return cp, offset
}
