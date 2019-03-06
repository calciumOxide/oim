package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
		"../../loader/clazz"
		"../../loader/clazz/item"
)

type I_getfield struct {
}

func init()  {
	INSTRUCTION_MAP[0xb4] = &I_getfield{}
}

func (s I_getfield)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "getfield exce >>>>>>>>>\n")

	index := (uint16(ctx.Code[ctx.PC]) << 8) | uint16(ctx.Code[ctx.PC + 1])
	ctx.PC += 2
	ref, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(ref) != reflect.TypeOf(types.Jreference{}) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	if ref.(types.Jreference).Reference == nil {
		except, _ := variator.AllocExcept(variator.NullPointerException)
		ctx.Throw(except)
		return nil
	}
	if reflect.TypeOf(ref.(types.Jreference).Reference) != reflect.TypeOf(types.Jobject{}) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}
	jobject := ref.(types.Jreference).Reference.(types.Jobject)
	cp, _ := ctx.Clazz.GetConstant(index)
	if cp == nil || reflect.TypeOf(cp.Info) != reflect.TypeOf(&item.FieldRef{}) {
		except, _ := variator.AllocExcept(variator.ClassNotFindException)
		ctx.Throw(except)
		return nil
	}
	objClassCp, _ := ctx.Clazz.GetConstant(cp.Info.(*item.FieldRef).ClassIndex)
	objClassNameCp, _ := ctx.Clazz.GetConstant(objClassCp.Info.(*item.Class).NameIndex)
	if jobject.Class.(*clazz.ClassFile) != clazz.GetClass(objClassNameCp.Info.(*item.Utf8).Str) {
		except, _ := variator.AllocExcept(variator.ClassNotFindException)
		ctx.Throw(except)
		return nil
	}
	fieldNameIndex := cp.Info.(*item.FieldRef).NameAndTypeIndex
	fieldName, _ := ctx.Clazz.GetConstant(fieldNameIndex)
	if cp == nil || reflect.TypeOf(fieldName.Info) != reflect.TypeOf(&item.NameAndType{}) {
		except, _ := variator.AllocExcept(variator.ClassNotFindException)
		ctx.Throw(except)
		return nil
	}
	nameIndex := fieldName.Info.(*item.NameAndType).NameIndex
	name, _ := ctx.Clazz.GetConstant(nameIndex)
	if cp == nil || reflect.TypeOf(name.Info) != reflect.TypeOf(&item.Utf8{}) {
		except, _ := variator.AllocExcept(variator.ClassNotFindException)
		ctx.Throw(except)
		return nil
	}
	ctx.CurrentFrame.PushFrame(jobject.Fileds[name.Info.(*item.Utf8).Str])
	return nil
}

func (s I_getfield)Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	fields := make(map[string]interface{})
	fields["app"] = "bb"
	f.PushFrame(types.Jreference{
		Reference: types.Jobject{
			Class: clazz.GetClass("com/oxide/A"),
			Fileds: fields,
		},
	})
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x0, 0x0, 0x3},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		获取对象的字段值
======================================================================================
						||		getfield
						||------------------------------------------------------------
						||		indexbyte1
						||------------------------------------------------------------
						||		indexbyte2
		格式				||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		getfield = 180(0xb4)
======================================================================================
						||		...，objectref →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||		
						||		objectref 必须是一个 reference 类型的数据，在指令执行时，objectref 将从操作数栈中出栈。
						||		无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运行时常量池的索引值，
						||		构建方式为(indexbyte1 << 8) | indexbyte2，
						||		该索引所指向的运行时常量池项应当是一个字段(§5.1) 的符号引用，
		描述				||		它包含了字段的名称和描述符，以及包含该字段的类的符号引用。
						||		这个字段的符号引用是已被解析过的(§5.4.3.2)。指令执行后，
						||		被 objectref 所引用的对象中该字段的值将会被取出，并推入到操作数栈顶。
						||		objectref 所引用的对象不能是数组类型，如果取值的字段是 protected 的(§4.6)，
						||		如果取值的字段是 protected 的(§4.6)，并且这个字段是当前类的父类成员，
						||		并且这个字段没有在同一个 运行时包(§5.3)中定义过，那 objectref 所指向的对象的类型必须为当 前类或者当前类的子类。
						||
======================================================================================
						||		
						||
						||		在字段的符号引用解析过程中，任何在§5.4.3.2 中描述过的异常都可能会 被抛出。
	   链接时异常			||		另外，如果已解析的字段是一个静态(static)字段，getfield 指令将会 抛出一个 IncompatibleClassChangeError 异常
						||		
						||		
						||		
======================================================================================
						||
						||
						||
	   运行时异常			||		如果 objectref 为 null，getfield 指令将抛出一个 NullPointerException.异常。
						||
						||
						||
======================================================================================
						||
						||
						||
		注意				||		不可以使用 getfield 指令来访问数组对象的 length 字段，如果要访问这个字段，应当使用 arraylength 指令。
						||
						||
						||
======================================================================================
 */