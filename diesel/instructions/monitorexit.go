package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_monitorenter struct {
}

func init()  {
	INSTRUCTION_MAP[0xc2] = &I_monitorenter{}
}

func (s I_monitorenter)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "monitorenter exce >>>>>>>>>\n")

	obj, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(obj) != reflect.TypeOf(types.Jreference{}) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	if obj.(*types.Jreference).Reference == nil {
		except, _ := variator.AllocExcept(variator.NullPointerException)
		ctx.Throw(except)
		return nil
	}

	obj.(*types.Jreference).MonitorEnter(ctx.ThreadId)

	return nil
}

func (s I_monitorenter)Test(octx *runtime.Context) *runtime.Context {
	f := new(runtime.Frame)
	f.PushFrame(&types.Jarray{
		Reference: []types.Jbyte{1, 2, 3, 4},
	})
	f.PushFrame(types.Jlong(9))
	f.PushFrame(types.Jlong(9))
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
		操作				||		进入一个对象的 monitor
======================================================================================
						||		monitorenter
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
		结构				||		monitorenter = 194(0xc2)
======================================================================================
						||		...，objectref →
	   操作数栈			||------------------------------------------------------------
						||		„，
======================================================================================
						||
		描述				||		objectref 必须为 reference 类型数据。

任何对象都有一个 monitor 与之关联。当且仅当一个 monitor 被持有后，它将处于锁定状态。
线程执行到 monitorenter 指令时，将会尝试获取 objectref 所对应的 monitor 的所有权，那么:
如果 objectref 的 monitor 的进入计数器为 0，那线程可以成功进入 monitor，以及将计数器值设置为 1。
当前线程就是 monitor 的所有者。 如果当前线程已经拥有 objectref 的 monitor 的所有权，那它可以重入这 个 monitor，重入时需将进入计数器的值加 1。
如果其他线程已经拥有 objectref 的 monitor 的所有权，那当前线程将被 阻塞，直到 monitor 的进入计数器值变为 0 时，重新尝试获取 monitor 的 所有权。

						||
======================================================================================
						||		
	   运行时异常			||
						||	当 objectref 为 null 时，monitorenter 指令将抛出 NullPointerException 异常。
						||
======================================================================================
						||
		注意				||
						||一个 monitorenter 指令可能会与一个或多个 monitorexit 指令配合实现 Java 语言中 synchronized 同步语句块的语义(§3.14)。
但monitorenter 和 monitorexit 指令不会用来实现 synchronized 方法的 语义，尽管它们确实可以实现类似的语义。
当一个 synchronized 方法被调 用时，自动进入对应的 monitor，当方法返回时，自动退出 monitor，这些动作是 Java 虚拟机在调用和返回指令中隐式处理的。

对象与它的 monitor 之间的关联关系有很多种实现方式，这些内容已超出本 规范的范围之外，但可以稍作介绍。
monitor 即可以实现为与对象一同分配 和销毁，也可以在某条线程尝试获取对象所有权时动态生成，在没有任何线程 持有对象所有权时自动释放。

在 Java 语言里面，同步的概念除了包括 monitor 的进入和退出操作以外， 还包括有等待(Object.wait)和唤醒(Object.notifyAll 和 Object.notify)。
这些操作包含在 Java 虚拟机提供的标准包 java.lang 之中，而不是通过 Java 虚拟机的指令集来显式支持。

						||
						||
======================================================================================
 */