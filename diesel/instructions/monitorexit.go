package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	)

type I_monitorexit struct {
}

func init()  {
	INSTRUCTION_MAP[0xc2] = &I_monitorexit{}
}

func (s I_monitorexit)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "monitorexit exce >>>>>>>>>\n")

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

	if !obj.(*types.Jreference).MonitorExit(ctx.ThreadId) {
		except, _ := variator.AllocExcept(variator.IllegalMonitorStateException)
		ctx.Throw(except)
		return nil
	}

	return nil
}

func (s I_monitorexit)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		退出一个对象的 monitor
======================================================================================
						||		monitorexit
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
		结构				||		monitorexit = 195(0xc3)
======================================================================================
						||		...，objectref →
	   操作数栈			||------------------------------------------------------------
						||		„，
======================================================================================
						||
		描述				||		objectref 必须为 reference 类型数据。

执行 monitorexit 指令的线程必须是 objectref 对应的 monitor 的所有者。

指令执行时，线程把 monitor 的进入计数器值减 1，如果减 1 后计数器值为 0，
那线程退出 monitor，不再是这个 monitor 的拥有者。其他被这个 monitor 阻塞的线程可以尝试去获取这个 monitor 的所有权。

						||
======================================================================================
						||		
	   运行时异常			||
						||	当 objectref 为 null 时，monitorexit 指令将抛出 NullPointerException 异常。
另外，如果执行 monitorexit 的线程原本并没有这个 monitor 的所有权， 那 monitorexit 指令将抛出 IllegalMonitorStateException.异常。
另外，如果 Java 虚拟机执行 monitorexit 时发现违反了§2.11.10 中第 二条规则，那 monitorexit 指令将抛出 IllegalMonitorStateException.异常。
						||
======================================================================================
						||
		注意				||
						||一个 monitorenter 指令可能会与一个或多个 monitorexit 指令配合实现 Java 语言中 synchronized 同步语句块的语义(§3.14)。
但monitorexit 和 monitorexit 指令不会用来实现 synchronized 方法的 语义，尽管它们确实可以实现类似的语义。

Java 虚拟机对在 synchronized 方法和 synchronized 同步语句块中抛出的异常有不同的处理方式:
在 synchronized 方法正常完成时，monitor 通过 Java 虚拟机的返回指令 退出。在 synchronized 方法非正常完成时，monitor 通过 Java 虚拟机的 athrow 指令退出。
当有异常从 synchronized 同步语句块抛出，将由 Java 虚拟机异常处理机 制(§3.14)来保证退出了之前在 synchronized 同步语句块开始时进入的 monitor。

						||
						||
======================================================================================
 */