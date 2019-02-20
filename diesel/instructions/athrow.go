package instructions

import (
	"../../types"
	"../runtime"
	"../../utils"
	"../../loader/clazz"
	"../../loader/clazz/attribute"
		)

type I_athrow struct {
}

func init()  {
	INSTRUCTION_MAP[0xbf] = &I_athrow{}
}

func (s I_athrow)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "athrow exce >>>>>>>>>\n")

	ref, _ := ctx.CurrentFrame.PopFrame()
	ctx.Throw(ref)

	return nil
}

func (s I_athrow)Test() *runtime.Context {
	f2 := new(runtime.Frame)
	f2.PushFrame(uint32(2))
	f2.PushFrame(&types.Jreference{
		Reference: &types.Jobject{
			ClassTypeIndex: 333,
		},
	})
	f1 := new(runtime.Frame)
	f1.PushFrame(uint32(1))
	f1.PushFrame(uint32(1))

	a2 := new(runtime.Aborigines)
	a2.Layers = append(a2.Layers, []uint32{2})
	a1 := new(runtime.Aborigines)
	a1.Layers = append(a1.Layers, []uint32{1})

	m2 := new(clazz.Method)
	m2.NameIndex = 2
	m2.DescriptorIndex = 2
	m2.Attributes = []*clazz.Attribute{
		{
			AttributeName: clazz.CODE_ATTR,
			AttributeItem: &attribute.Codes{
				ExceptTableLength: 1,
				ExceptTable: []*attribute.ExceptTable{
					{
						StartPc: 1,
						EndPc: 4,
						HandlerPc: 666,
						CatchType: 333,
					},
				},
			},
		},
	}
	m1 := new(clazz.Method)
	m1.NameIndex = 1

	return &runtime.Context{
		PC: 2,

		Code: []byte{0x2},
		CodeStack: [][]byte{{0x1}},

		CurrentFrame: f2,
		FrameStack: []*runtime.Frame{f1},

		CurrentMethod: m2,
		MethodStack: []*clazz.Method{m1},

		CurrentAborigines: a2,
		AboriginesStack: []*runtime.Aborigines{a1},
	}
}
/**
======================================================================================
		操作				||		抛出一个异常实例(exception 或者 error)
======================================================================================
						||		athrow
						||------------------------------------------------------------
						||
						||------------------------------------------------------------
						||
		格式				||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
						||------------------------------------------------------------
						||		
======================================================================================
		结构				||		athrow = 191(0xbf)
======================================================================================
						||		...，objectref →
	   操作数栈			||------------------------------------------------------------
						||		objectref
======================================================================================
						||		
						||		objectref 必须为一个 reference 类型的数据，它指向一个 Throwable 或其子类的对象实例。
						||		在指令执行时，objectref 首先从操作数栈中出栈，然后通过§2.10 中描述的算法搜索当前方法(§2.6)中与 objectref 的类 型相匹配的第一个异常处理器。
						||
						||		如果找到了适合 objectref 的异常处理器，这个异常处理器将包含一个用于 处理此异常的代码位置。
						||		PC 寄存器的值就会被重设为异常处理器指定的那个 位置上，整个当前栈帧的操作数栈都会被清空，objectref 重新压入到操作 数栈中，然后程序继续执行。
		描述				||
						||		如果在当前栈帧中没有找到适合的异常处理器，那么栈帧就要从操作数栈中出 栈，如果当前栈帧对应的方法是一个同步方法，
						||		那在方法调用时持有或重入的 管程就应当释放(对于重入来说是计数减 1)，就像执行了 monitorexit 一 样。
						||		最后，这个栈帧的调用者被恢复。如果此栈帧仍然没有找到合适的异常处 理器，那它也会继续退出，objectref 也会不断重新抛出，
						||		假设已经没有任 何的栈帧可以退出，那当前线程将被结束掉。
						||
======================================================================================
						||		
						||
						||
	   链接时异常			||
						||		
						||		
						||		
======================================================================================
						||
						||
						||		如果 objectref 为 null，athrow 指令将会抛出 NullPointerException 来代替 objectref 所代表的异常。
						||
						||		另外，如果虚拟机实现没有严格执行在§2.11.10 中规定的结构化锁定规则， 导致当前方法是一个同步方法，但当前线程在调用方法时没有成功持有或重入
		运行时异常		||		相应的管程，那 athrow 指令将会抛出 IllegalMonitorStateException 异常。这是可能出现的，
						||		譬如一个同步方法只包含了对方法同步对象的 monitorexit 指令，但是未包含配对的 monitorenter 指令。
						||		另外，如果虚拟机实现严格执行了§2.11.10 中规定的结构化锁定规则，但 当前方法调用时，其中的第一条规则被违反的话，
						||		athrow 指令也会抛出 IllegalMonitorStateException 异常。
						||
						||
======================================================================================
						||
						||		athrow 指令的操作数栈图(本表中“操作数栈”行的图)可能会产生一些误 解:如果当前方法中某个异常处理器被匹配到，
		注意				||		athrow 指令将抛弃掉操作数 栈上所有的值，然后重新将被抛出的异常对象入栈，但是如果在当前方法中没 有找到适合的异常处理器，
						||		即异常被抛到方法调用链其他地方时，被清空的和 objectref 入栈的操作数栈是真正处理异常的那个方法的操作数栈，
						||		而从最 初抛出异常的那个方法一直到最终处理异常的那个方法(不含)之间的栈帧全 部都会被丢弃掉。
						||
======================================================================================
 */