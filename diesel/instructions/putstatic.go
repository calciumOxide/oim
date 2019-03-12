package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	"../../loader/clazz/item"
	"../../loader/clazz"
)

type I_putstatic struct {
}

func init()  {
	INSTRUCTION_MAP[0xb3] = &I_putstatic{}
}

func (s I_putstatic)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "putstatic exce >>>>>>>>>\n")

	index := (uint16(ctx.Code[ctx.PC]) << 8) | uint16(ctx.Code[ctx.PC + 1])
	ctx.PC += 2

	value, _ := ctx.CurrentFrame.PopFrame()

	cp, _ := ctx.Clazz.GetConstant(index)
	if cp == nil || reflect.TypeOf(cp.Info) != reflect.TypeOf(&item.FieldRef{}) {
		except, _ := variator.AllocExcept(variator.ClassNotFindException)
		ctx.Throw(except)
		return nil
	}
	fieldNameIndex := cp.Info.(*item.FieldRef).NameAndTypeIndex
	classNameIndex := cp.Info.(*item.FieldRef).ClassIndex
	fieldName, _ := ctx.Clazz.GetConstant(fieldNameIndex)
	className, _ := ctx.Clazz.GetConstant(classNameIndex)

	if fieldName == nil || reflect.TypeOf(fieldName.Info) != reflect.TypeOf(&item.NameAndType{}) {
		except, _ := variator.AllocExcept(variator.ClassNotFindException)
		ctx.Throw(except)
		return nil
	}
	if className == nil || reflect.TypeOf(className.Info) != reflect.TypeOf(&item.Class{}) {
		except, _ := variator.AllocExcept(variator.ClassNotFindException)
		ctx.Throw(except)
		return nil
	}
	fnameIndex := fieldName.Info.(*item.NameAndType).NameIndex
	cnameIndex := className.Info.(*item.Class).NameIndex

	fname, _ := ctx.Clazz.GetConstant(fnameIndex)
	cname, _ := ctx.Clazz.GetConstant(cnameIndex)

	if fname == nil || reflect.TypeOf(fname.Info) != reflect.TypeOf(&item.Utf8{}) {
		except, _ := variator.AllocExcept(variator.ClassNotFindException)
		ctx.Throw(except)
		return nil
	}
	if cname == nil || reflect.TypeOf(cname.Info) != reflect.TypeOf(&item.Utf8{}) {
		except, _ := variator.AllocExcept(variator.ClassNotFindException)
		ctx.Throw(except)
		return nil
	}

	cla := clazz.GetClass(cname.Info.(*item.Utf8).Str)
	if !ctx.Cinit(cla) {
		return nil
	}

	field := cla.GetFiled(fname.Info.(*item.Utf8).Str, clazz.ACC_PRIVATE|clazz.ACC_PROTECTED|clazz.ACC_PUBLIC)
	field.Value = value

	except, _ := variator.AllocExcept(variator.FieldNotFindException)
	ctx.Throw(except)
	return nil
}

func (s I_putstatic)Test(octx *runtime.Context) *runtime.Context {
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
		Code: []byte{0x0, 0x0, 0x5},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		设置对象的静态字段值
======================================================================================
						||		putstatic
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
		结构				||		putstatic = 179(0xb3)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，
======================================================================================
						||		
		描述				||		无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运 行时常量池的索引值，构建方式为(indexbyte1 << 8)| indexbyte2，
该索引所指向的运行时常量池项应当是一个字段(§5.1)的符号引用，它包 含了字段的名称和描述符，
以及包含该字段的类或接口的符号引用。这个字段 的符号引用是已被解析过的(§5.4.3.2)。
在字段被成功解析之后，如果字段所在的类或者接口没有被初始化过(§5.5)， 那指令执行时将会触发其初始化过程。
被 putstatic 指令存储到字段中的 value 值的类型必须与字段的描述符相匹 配(§4.3.2)。
如果字段描述符的类型是 boolean、byte、char、short 或者 int，那么 value 必须为 int 类型。
如果字段描述符的类型是 float、 long 或者 double，那 value 的类型必须相应为 float、long 或者 double。
如果字段描述符的类型是 reference 类型，那 value 必须为一个可与之匹配 (JLS §5.2)的类型。
如果字段被声明为final的，那就只有在当前类的 类初始化方法(<clinit>)中设置当前类的 final 字段才是合法的。
指令执行时，value 从操作数栈中出栈，根据(§2.8.3)中定义的转换规则 转换为 value’，类的指定字段的值将被设置为 value’。


						||
======================================================================================
						||		
						||
						||		在字段的符号引用解析过程中，任何在§5.4.3.2 中描述过的异常都可能会 被抛出。
	   链接时异常			||		另外，如果已解析的字段是一个非静态(not static)字段，putstatic 指令将会抛出一个 IncompatibleClassChangeError 异常。
另外，如果字段声明为 final，那就只有在当前类的实例初始化方法 (<clinit>)中设置当前类的 final 字段才是合法的，否则将会抛出 IllegalAccessError 异常。
						||		
						||		
======================================================================================
						||
						||
						||
	   运行时异常			||		另外，如果 putstatic 指令触发了所涉及的类或接口的初始化，那 putstatic 指令就可能抛出在§5.5 中描述到的任何异常。
						||
						||
======================================================================================
						||
		注意				||
						||putstatic 指令只有在接口初始化时才能用来设置接口字段的值，接口字段 只会在接口初始化的时候初始化赋值一次(§5.5，JLS §9.3.1)。
						||
						||
======================================================================================
 */