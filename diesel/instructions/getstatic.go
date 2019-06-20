package instructions

import (
	"../../loader/binary"
	"../../loader/binary/item"
	"../../utils"
	"../oil/types"
	"../runtime"
	"../variator"
	"reflect"
)

type I_getstatic struct {
}

func init() {
	INSTRUCTION_MAP[0xb2] = &I_getstatic{}
}

func (s I_getstatic) Stroke(ctx *runtime.Context) error {
	utils.Log(1, "getstatic exce >>>>>>>>>\n")

	index := (uint16(ctx.Code[ctx.PC]) << 8) | uint16(ctx.Code[ctx.PC+1])
	ctx.PC += 2

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

	cla := binary.GetClass(cname.Info.(*item.Utf8).Str)
	if cla.Fields != nil && len(cla.Fields) > 0 {
		i := 0
		for i = 0; i < len(cla.Fields); i++ {
			field := cla.Fields[i]
			n := field.NameIndex
			f, _ := cla.GetConstant(n)
			if fname.Info.(*item.Utf8).Str == f.Info.(*item.Utf8).Str {
				ctx.CurrentFrame.PushFrame(field.Value)
				return nil
			}
		}
	}
	except, _ := variator.AllocExcept(variator.FieldNotFindException)
	ctx.Throw(except)
	return nil
}

func (s I_getstatic) Test() *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	fields := make(map[string]interface{})
	fields["app"] = "bb"
	f.PushFrame(types.Jreference{
		Reference: types.Jobject{
			ClassTypeIndex: 4,
			Fileds:         fields,
		},
	})
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code:              []byte{0x0, 0x0, 0x5},
		CurrentFrame:      f,
		CurrentAborigines: a,
	}
}

/**
======================================================================================
		操作				||		获取对象的静态字段值
======================================================================================
						||		getstatic
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
		结构				||		getstatic = 178(0xb2)
======================================================================================
						||		...， →
	   操作数栈			||------------------------------------------------------------
						||		...，value
======================================================================================
						||
						||		无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运 行时常量池的索引值，
						||		构建方式为(indexbyte1 << 8)| indexbyte2，该索引所指向的运行时常量池项应当是一个字段(§5.1)的符号引用，
						||		它包 含了字段的名称和描述符，以及包含该字段的类或接口的符号引用。
		描述				||		这个字段 的符号引用是已被解析过的(§5.4.3.2)。
						||		在字段被成功解析之后，如果字段所在的类或者接口没有被初始化过(§5.5)， 那指令执行时将会触发其初始化过程。
						||		参数所指定的类或接口的该字段的值将会被取出，并推入到操作数栈顶。
						||
======================================================================================
						||
						||
						||		在字段的符号引用解析过程中，任何在§5.4.3.2 中描述过的异常都可能会 被抛出。
	   链接时异常			||		另外，如果已解析的字段是一个非静态(not static)字段，getstatic 指令将会抛出一个 IncompatibleClassChangeError 异常
						||
						||
						||
======================================================================================
						||
						||
						||
	   运行时异常			||		如果 getstatic 指令触发了所涉及的类或接口的初始化，那 getstatic 指 令就可能抛出在§5.5 中描述到的任何异常。
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
						||
======================================================================================
*/
