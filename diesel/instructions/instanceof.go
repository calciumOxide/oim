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

type I_instanceof struct {
}

func init()  {
	INSTRUCTION_MAP[0xc1] = &I_instanceof{}
}

func (s I_instanceof)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "instanceof exce >>>>>>>>>\n")

	i := uint16(ctx.Code[ctx.PC]) << 8 | uint16(ctx.Code[ctx.PC + 1])
	obj, _ := ctx.CurrentFrame.PopFrame()
	ot := reflect.TypeOf(&types.Jreference{})

	cp, _ := ctx.Clazz.GetConstant(i)

	if obj != nil && reflect.TypeOf(obj) != ot && reflect.TypeOf(obj) != reflect.TypeOf(&types.Jarray{}) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	if obj == nil || (reflect.TypeOf(obj) == ot && obj.(*types.Jreference).Reference == nil) {
		ctx.CurrentFrame.PushFrame(types.Jint(0))
		return nil
	}
	name, _ := ctx.Clazz.GetConstant(cp.Info.(*item.Class).NameIndex)
	typeClass := clazz.GetClass(name.Info.(*item.Utf8).Str)
	if !ctx.Cinit(typeClass) {
		return nil
	}

	o := &types.Jreference{}
	if reflect.TypeOf(obj) == reflect.TypeOf(&types.Jarray{}) {
		o.ElementType = obj.(*types.Jarray).ElementJype
	} else {
		o = obj.(*types.Jreference)
	}
	if reflect.TypeOf(o.ElementType) == reflect.TypeOf(&clazz.ClassFile{}) {
		if o.ElementType.(*clazz.ClassFile).AccessFlags == clazz.ACC_INTERFACE {
			if typeClass.AccessFlags&clazz.ACC_INTERFACE == clazz.ACC_INTERFACE {
				cn, _ := typeClass.GetConstant(typeClass.ThisClass)
				tn, _ := o.ElementType.(*clazz.ClassFile).GetConstant(o.ElementType.(*clazz.ClassFile).ThisClass)
				sc := clazz.GetClass(tn.Info.(*item.Utf8).Str)
				if !ctx.Cinit(sc) {
					return nil
				}

				for ; ; {
					if cn.Info.(*item.Utf8).Str == tn.Info.(*item.Utf8).Str {
						ctx.CurrentFrame.PushFrame(types.Jint(1))
						return nil
					}
					if sc == nil || sc.SuperClass == 0 {
						break
					}
					tn, _ = sc.GetConstant(sc.ThisClass)
					sn, _ := sc.GetConstant(sc.SuperClass)
					sc = clazz.GetClass(sn.Info.(*item.Utf8).Str)
					if !ctx.Cinit(sc) {
						return nil
					}
				}
			} else {
				if typeClass.SuperClass == 0 {
					ctx.CurrentFrame.PushFrame(types.Jint(1))
					return nil
				}
			}
		} else {
			if typeClass.AccessFlags&clazz.ACC_INTERFACE == clazz.ACC_INTERFACE {
				for ; ; {
					is := o.ElementType.(*clazz.ClassFile).Interfaces
					for ii := 0; ii < len(is); ii++ {
						in, _ := o.ElementType.(*clazz.ClassFile).GetConstant(is[ii])
						cn, _ := typeClass.GetConstant(typeClass.ThisClass)
						if in.Info.(*item.Utf8).Str == cn.Info.(*item.Utf8).Str {
							ctx.CurrentFrame.PushFrame(types.Jint(1))
							return nil
						}
						inf := clazz.GetClass(in.Info.(*item.Utf8).Str)
						if !ctx.Cinit(inf) {
							return nil
						}

						for ; ; {
							if inf != nil && inf.SuperClass != 0 {
								in, _ = inf.GetConstant(inf.ThisClass)
								if in.Info.(*item.Utf8).Str == cn.Info.(*item.Utf8).Str {
									ctx.CurrentFrame.PushFrame(types.Jint(1))
									return nil
								}
								super, _ := inf.GetConstant(inf.SuperClass)
								inf = clazz.GetClass(super.Info.(*item.Utf8).Str)
								if !ctx.Cinit(inf) {
									return nil
								}

							}

						}
					}
					if o == nil || o.ElementType.(*clazz.ClassFile).SuperClass == 0 {
						break
					}
					sn, _ := o.ElementType.(*clazz.ClassFile).GetConstant(o.ElementType.(*clazz.ClassFile).SuperClass)
					is = clazz.GetClass(sn.Info.(*item.Utf8).Str).Interfaces
					if !ctx.Cinit(is) {
						return nil
					}

				}
			} else {
				cn, _ := typeClass.GetConstant(typeClass.ThisClass)
				tn, _ := o.ElementType.(*clazz.ClassFile).GetConstant(o.ElementType.(*clazz.ClassFile).ThisClass)
				sc := clazz.GetClass(tn.Info.(*item.Utf8).Str)
				if !ctx.Cinit(sc) {
					return nil
				}

				for ; ; {
					if cn.Info.(*item.Utf8).Str == tn.Info.(*item.Utf8).Str {
						ctx.CurrentFrame.PushFrame(types.Jint(1))
						return nil
					}
					if sc == nil || sc.SuperClass == 0 {
						break
					}
					tn, _ = sc.GetConstant(sc.ThisClass)
					sn, _ := sc.GetConstant(sc.SuperClass)
					sc = clazz.GetClass(sn.Info.(*item.Utf8).Str)
					if !ctx.Cinit(sc) {
						return nil
					}

				}
			}
		}
	} else {

	}

	ctx.CurrentFrame.PushFrame(types.Jint(0))
	return nil
}

func (s I_instanceof)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jfloat(9.123456789012345))
	f.PushFrame(types.Jint(9))
	a := new(runtime.Aborigines)
	a.Layers = append(a.Layers, &[]uint32{1234})
	return &runtime.Context{
		Code: []byte{0x0},
		CurrentFrame: f,
		CurrentAborigines: a,
	}
}
/**
======================================================================================
		操作				||		判断对象是否指定的类型
======================================================================================
						||		instanceof
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
		结构				||		instanceof = 193(0xc1)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		...，result
======================================================================================
						||		
						||		objectref 必须是一个 reference 类型的数据，在指令执行时，objectref 将从操作数栈中出栈。
						||		无符号数 indexbyte1 和 indexbyte2 用于构建一个当前类(§2.6)的运行时常量池的索引值，
		描述				||		构建方式为(indexbyte1 << 8) | indexbyte2，该索引所指向的运行时常量池项应当是一个类、接口或者数 组类型的符号引用。
						||		如果 objectref 为 null 的话，那 instanceof 指令将会把 int 值 0 推入 到操作数栈栈顶。
						||		另外，如果参数指定的类、接口或者数组类型会被虚拟机解析(§5.4.3.1)。 如果 objectref 可以转换为这个类、接口或者数组类型，那 instanceof 指令将会把 int 值 1 推入到操作数栈栈顶;否则，推入栈顶的就是 int 值 0。
						||		以下规则可以用来确定一个非空的 objectref 是否可以转换为指定的已解析 类型:
						||		   假设 S 是 objectref 所指向的对象的类型，T 是进行比较的已解析的 类、接口或者数组类型，
						||		   checkcast 指令根据这些规则来判断转换是否成立:
						||
						||		 如果 S 是类类型(Class Type)，那么:
						||			 如果T也是类类型，那S必须与T是同一个类类型，或者S是T所 代表的类型的子类。
						||			 如果 T 是接口类型，那 S 必须实现了 T 的接口。
						||		 如果 S 是接口类型(Interface Type)，那么:
						||			 如果 T 是类类型，那么 T 只能是 Object。
						||			 如果T是接口类型，那么T与S应当是相同的接口，或者T是S的父接口。
						||		 如果 S 是数组类型(Array Type)，假设为 SC[]的形式，这个数组的组件类型为 SC，那么:
						||			 如果 T 是类类型，那么 T 只能是 Object。
						||			 如果 T 是数组类型，假设为 TC[]的形式，这个数组的组件类型为 TC，那么下面两条规则之一必须成立:
						||				 TC 和 SC 是同一个原始类型。
						||				 TC 和 SC 都是 reference 类型，并且 SC 能与 TC 类型相匹配(以此处描述的规则来判断是否互相匹配)。
						||		？？如果T是接口类型，那T必须是数组所实现的接口之一(JLS3 §4.10.3)。
						||
						||
						||
======================================================================================
						||		
						||
						||
	   链接时异常			||		在类、接口或者数组的符号解析阶段，任何在§5.4.3.1 章节中描述的异常 都可能被抛出。
						||		
						||		
						||		
======================================================================================
						||
						||
						||
	   运行时异常			||		？？？如果 objectref 不能转换成参数指定的类、接口或者数组类型，checkcast 指令将抛出 ClassCastException 异常
						||
						||
						||
======================================================================================
						||
						||
						||
		注意				||		instanceof 指令与 checkcast 指令非常类似，它们之间的区别是如何处理 null 值的情况、测试类型转换的结果反馈方式(checkcast 是抛异常，而 instanceof 是返回一个比较结果)以及指令执行后对操作数栈的影响。
						||
						||
						||
======================================================================================
 */