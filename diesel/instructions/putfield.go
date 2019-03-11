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

type I_putfield struct {
}

func init()  {
	INSTRUCTION_MAP[0xb4] = &I_putfield{}
}

func (s I_putfield)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "putfield exce >>>>>>>>>\n")

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

func (s I_putfield)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		设置对象字段
======================================================================================
						||		putfield
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
		结构				||		putfield = 181(0xb5)
======================================================================================
						||		...，objectref，value →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||
		描述				||		无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运 行时常量池的索引值，构建方式为(indexbyte1 << 8)| indexbyte2，
该索引所指向的运行时常量池项应当是一个字段(§5.1)的符号引用，它包 含了字段的名称和描述符，以及包含该字段的类的符号引用。
objectref 所 引用的对象不能是数组类型，如果取值的字段是 protected 的(§4.6)，
并 且这个字段是当前类的父类成员，并且这个字段没有在同一个运行时包 (§5.3)中定义过，那 objectref 所指向的对象的类型必须为当前类或者 当前类的子类。

这个字段的符号引用是已被解析过的(§5.4.3.2)。被 putfield 指令存储 到字段中的 value 值的类型必须与字段的描述符相匹配(§4.3.2)。
如果字 段描述符的类型是 boolean、byte、char、short 或者 int，那么 value 必须为 int 类型。
如果字段描述符的类型是 float、long 或者 double，那 value 的类型必须相应为 float、long 或者 double。
如果字段描述符的类 型是 reference 类型，那 value 必须为一个可与之匹配(JLS §5.2)的 类型。
如果字段被声明为 final 的，那就只有在当前类的实例初始化方法 (<init>)中设置当前类的 final 字段才是合法的。

指令执行时，value 和 objectref 从操作数栈中出栈，objectref 必须为reference 类型数据，value 将根据(§2.8.3)中定义的转换规则转换为 value’，objectref 的指定字段的值将被设置为 value’。



======================================================================================
						||		
						||
	   链接时异常			||		在字段的符号引用解析过程中，任何在§5.4.3.2 中描述过的异常都可能会 被抛出。
另外，如果已解析的字段是一个静态(static)字段，getfield 指令将会 抛出一个 IncompatibleClassChangeError 异常。
另外，如果字段声明为 final，那就只有在当前类的实例初始化方法(<init>) 中设置当前类的 final 字段才是合法的，否则将会抛出 IllegalAccessError 异常。
						||
						||		
======================================================================================
						||
						||
						||
	   运行时异常			||		另外，如果 objectref 为 null，putfield 指令将抛出一个 NullPointerException.异常。

						||
						||
						||
======================================================================================
						||
						||
						||
		注意				||
						||
						||
======================================================================================
 */