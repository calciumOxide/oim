package instructions

import (
	"../runtime"
	"../../utils"
	"../../types"
	"reflect"
	"../variator"
	"../../loader/clazz/item"
)

type I_lreturn struct {
}

func init()  {
	INSTRUCTION_MAP[0xad] = &I_lreturn{}
}

func (s I_lreturn)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "lreturn exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.Jlong(0)) {
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	methodDescCp, _ := ctx.Clazz.GetConstant(ctx.CurrentMethod.DescriptorIndex)
	returnType := ctx.CurrentMethod.GetReturnType(methodDescCp.Info.(*item.Utf8).Str)
	ctx.CurrentMethod.CheckReturnType(returnType, value)
	ctx.PopContext()
	ctx.CurrentFrame.PushFrame(value)
	return nil
}

func (s I_lreturn)Test(octx *runtime.Context) *runtime.Context {
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
		操作				||		结束方法，并返回一个 long 类型数据
======================================================================================
						||		lreturn
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
		结构				||		lreturn = 173(0xad)
======================================================================================
						||		...，value1 →
	   操作数栈			||------------------------------------------------------------
						||		[empty]
======================================================================================
						||
		描述				||		当前方法的返回值必须为 long 类型，value 必须是一个 long 类型的数据。
如果当前方法是一个同步(声明为 synchronized)方法，那在方法调用时
进入或者重入的管程应当被正确更新状态或退出，就像当前线程执行了 monitorexit 指令一样。
如果执行过程当中没有异常被抛出的话，那 value 将从当前栈帧(§2.6)中出栈，然后压入到调用者栈帧的操作数栈中，在当 前栈帧操作数栈中所有其他的值都将会被丢弃掉。
指令执行后，解释器会恢复调用者的栈帧，并且把程序控制权交回到调用者。

						||
======================================================================================
						||		
						||
	   运行时异常			||		如果虚拟机实现没有严格执行在§2.11.10 中规定的结构化锁定规则，导致 当前方法是一个同步方法，但当前线程在调用方法时没有成功持有(Enter)
或重入(Reentered)相应的管程，那 lreturn 指令将会抛出 IllegalMonitorStateException 异常。这是可能出现的，譬如一个同步 方法只包含了对方法同步对象的 monitorexit 指令，但是未包含配对的 monitorenter 指令。
另外，如果虚拟机实现严格执行了§2.11.10 中规定的结构化锁定规则，但 当前方法调用时，其中的第一条规则被违反的话，lreturn 指令也会抛出 IllegalMonitorStateException 异常。


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