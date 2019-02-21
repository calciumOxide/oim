package instructions

import (
	"../runtime"
	"../../utils"
	"../../loader/binary"
	"reflect"
	"../oil/types"
	"../variator"
)

type I_freturn struct {
}

func init()  {
	INSTRUCTION_MAP[0xae] = &I_freturn{}
}

func (s I_freturn)Stroke(ctx *runtime.Context) error {
	utils.Log(1, "freturn exce >>>>>>>>>\n")

	value, _ := ctx.CurrentFrame.PopFrame()

	if reflect.TypeOf(value) != reflect.TypeOf(types.JFN) && reflect.TypeOf(value) != reflect.TypeOf(types.Jfloat(0)) &&
		reflect.TypeOf(value) != reflect.TypeOf(types.JFO) && reflect.TypeOf(value) != reflect.TypeOf(types.JFU){
		except, _ := variator.AllocExcept(variator.ClassCastException)
		ctx.Throw(except)
		return nil
	}

	ctx.PopContext()
	ctx.CurrentFrame.PushFrame(value)

	return nil
}

func (s I_freturn)Test() *runtime.Context {
	f2 := new(runtime.Frame)
	f2.PushFrame(uint32(2))
	f2.PushFrame(types.Jfloat(222))
	f2.PushFrame(nil)
	f2.PushFrame(types.Jfloat(222))
	f1 := new(runtime.Frame)
	f1.PushFrame(uint32(1))
	f1.PushFrame(uint32(1))

	a2 := new(runtime.Aborigines)
	a2.Layers = append(a2.Layers, []uint32{2})
	a1 := new(runtime.Aborigines)
	a1.Layers = append(a1.Layers, []uint32{1})

	m2 := new(binary.Method)
	m2.NameIndex = 2
	m2.DescriptorIndex = 2
	m1 := new(binary.Method)
	m1.NameIndex = 1

	return &runtime.Context{
		PC: 2,

		Code: []byte{0x2},
		CodeStack: [][]byte{{0x1}},

		CurrentFrame: f2,
		FrameStack: []*runtime.Frame{f1},

		CurrentMethod: m2,
		MethodStack: []*binary.Method{m1},

		CurrentAborigines: a2,
		AboriginesStack: []*runtime.Aborigines{a1},
	}
}
/**
======================================================================================
		操作				||		结束方法，并返回一个 float 类型数据
======================================================================================
						||		freturn
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
		结构				||		freturn = 174(0xae)
======================================================================================
						||		...，value →
	   操作数栈			||------------------------------------------------------------
						||		[empty]
======================================================================================
						||		
						||
						||		当前方法的返回值必须为 float 类型，value 也必须是一个 float 类型 的数据。
						||		如果当前方法是一个同步(声明为 synchronized)方法，那在方法调用时 进入或者重入的管程应当被正确更新状态或退出，就像当前线程执行了 monitorexit 指令一样。
		描述				||		如果执行过程当中没有异常被抛出的话，那 value 将从当前栈帧(§2.6)中出栈，并且经过数值集合转换(§2.8.3) 后得到值 value’，并压入到调用者栈帧的操作数栈中，在当前栈帧操作数栈中所有其他的值都将会被丢弃掉。
						||		指令执行后，解释器会恢复调用者的栈帧，并且把程序控制权交回到调用者。
						||
						||
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
						||		如果虚拟机实现没有严格执行在§2.11.10 中规定的结构化锁定规则，导致 当前方法是一个同步方法，
						||		但当前线程在调用方法时没有成功持有(Enter)或重入(Reentered)相应的管程，那 freturn 指令将会抛出 IllegalMonitorStateException 异常。
	   运行时异常			||		这是可能出现的，譬如一个同步 方法只包含了对方法同步对象的 monitorexit 指令，但是未包含配对的 monitorenter 指令。
						||
						||		另外，如果虚拟机实现严格执行了§2.11.10 中规定的结构化锁定规则，但 当前方法调用时，
						||		其中的第一条规则被违反的话，freturn 指令也会抛出 IllegalMonitorStateException 异常。
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