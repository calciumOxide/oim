package runtime

import "../../types"
import "../../loader/clazz"
import (
	"../../loader/clazz/attribute"
	"reflect"
	"../variator"
)

type Context struct {
	PC                uint32
	Code              []uint8
	CodeStack         [][]uint8
	CurrentFrame      *Frame
	FrameStack        []*Frame
	CurrentAborigines *Aborigines
	AboriginesStack   []*Aborigines
	CurrentMethod     *clazz.Method
	MethodStack       []*clazz.Method
	Opoos             bool
	Clazz             *clazz.ClassFile
	ClazzStack        []*clazz.ClassFile
	ThreadId          int

	wide bool
}

type Frame struct {
	Layers []interface{}
	Depth uint32
}

type Aborigines struct {
	Layers []interface{}
	Count uint32
}

//type Layer struct {
//	Type  uint8
//	Value interface{}
//}


func (s *Context) Cinit(cf *clazz.ClassFile) bool {
	if !cf.ClassInit {
		clinit := cf.GetMethod("<clinit>", "()V")
		if clinit == nil {
			except, _ := variator.AllocExcept(variator.ClassNotFindException)
			s.Throw(except)
			return false
		}
		s.InvokeMethod(cf, clinit, nil)
		cf.ClassInit = true
		return true
	}
	return true
}

func (ctx *Context)InvokeMethod(class *clazz.ClassFile, method *clazz.Method, args []interface{}) *Context {
	a, _ := method.GetAttribute(clazz.CODE_ATTR)
	codes := a.AttributeItem.(*attribute.Codes)
	callee := &Context{
		PC: 0,
		Clazz: class,
		CurrentFrame: &Frame{
			//Layers: make([]interface{}, codes.MaxStack),
			Layers: []interface{}{},
			Depth:  0,
		},
		CurrentAborigines: &Aborigines{
			Layers: make([]interface{}, codes.MaxLocal),
			Count:  0,
		},
		CurrentMethod: method,
		Code: codes.Code,
	}
	if args != nil {
		length := len(args)
		for i := 0; i > length; i++ {
			callee.CurrentAborigines.SetAborigines(uint32(length -1 - i), args[i])
		}
	}
	ctx.PushContext(callee)
	return ctx
}

func (s *Context)PushContext(ctx *Context) error {
	
	s.CurrentFrame.PushFrame(s.PC)
	s.PC = ctx.PC
	
	s.FrameStack = append(s.FrameStack, s.CurrentFrame)
	s.CurrentFrame = ctx.CurrentFrame
	
	s.AboriginesStack = append(s.AboriginesStack, s.CurrentAborigines)
	s.CurrentAborigines = ctx.CurrentAborigines
	
	s.MethodStack = append(s.MethodStack, s.CurrentMethod)
	s.CurrentMethod = ctx.CurrentMethod

	s.CodeStack = append(s.CodeStack, s.Code)
	s.Code = ctx.Code

	s.ClazzStack = append(s.ClazzStack, s.Clazz)
	s.Clazz = ctx.Clazz

	return nil
}

func (s *Context)PopContext() (*Context, error) {

	f := s.CurrentFrame
	s.CurrentFrame = s.FrameStack[len(s.FrameStack) - 1]
	s.FrameStack = s.FrameStack[:len(s.FrameStack) - 1]

	pc, _ := s.CurrentFrame.PopFrame()
	s.PC = pc.(uint32)
	if s.Opoos {
		l, _ := f.PopFrame()
		s.CurrentFrame.PushFrame(l)
	}

	s.Code = s.CodeStack[len(s.CodeStack) - 1]
	s.CodeStack = s.CodeStack[:len(s.CodeStack) - 1]

	s.CurrentAborigines = s.AboriginesStack[len(s.AboriginesStack) - 1]
	s.AboriginesStack = s.AboriginesStack[:len(s.AboriginesStack) - 1]

	s.CurrentMethod = s.MethodStack[len(s.MethodStack) - 1]
	s.MethodStack = s.MethodStack[:len(s.MethodStack) - 1]

	s.Clazz = s.ClazzStack[len(s.ClazzStack)]
	s.ClazzStack = s.ClazzStack[:len(s.ClazzStack) - 1]
	return s, nil
}

func (s *Context)Throw(e interface{}) error {
	s.Opoos = true
	s.CurrentFrame.Clean()
	s.CurrentFrame.PushFrame(e)
	return nil
}

func (s *Context)Handle() error {
	if (!s.Opoos) {
		return nil
	}
	attr, _ := s.CurrentMethod.GetAttribute(clazz.CODE_ATTR)
	codes := attr.AttributeItem.(*attribute.Codes)
	if codes.ExceptTableLength > 0 {
		i := uint16(0)
		for ; i < codes.ExceptTableLength; i++ {
			et := codes.ExceptTable[i]
			l, _ := s.CurrentFrame.PeekFrame()
			if s.PC >= uint32(et.StartPc) && s.PC < uint32(et.EndPc) && IsCatched(l.(*types.Jreference), et) {
				s.PC = uint32(et.HandlerPc)
				s.CurrentFrame.PopFrame()
				s.Opoos = false
				return nil
			}
		}
	}
	s.PopContext()
	return nil
}

func (s *Context) PushWide() {
	s.wide = true
}

func (s *Context) PopWide() bool {
	b := s.wide
	s.wide = false
	return b
}

func IsCatched(r *types.Jreference, et *attribute.ExceptTable) bool {
	return r.Reference.(*types.Jobject).Class == et.CatchType
}


func (s *Frame)PushFrame(l interface{}) error {
	if s.Depth >= 1 && IsDoubleLong(s.Layers[s.Depth - 1]) {
		panic("set local variable double_long low byte.")
	}
	s.Layers = append(s.Layers, l)
	s.Depth += 1
	if IsDoubleLong(l) {
		s.Layers = append(s.Layers, "double_long")
		s.Depth += 1
	}
	return nil
}

func (s *Frame)PopFrame() (interface{}, error) {
	s.Depth -= 1
	l := s.Layers[s.Depth]
	s.Layers = s.Layers[0 : s.Depth]
	if s.Depth > 1 && IsDoubleLong(s.Layers[s.Depth - 1]) {
		s.Depth -= 1
		l = s.Layers[s.Depth]
		s.Layers = s.Layers[0 : s.Depth]
	}
	return l, nil
}

func (s *Frame)PeekFrame() (interface{}, error) {
	layer := s.Layers[s.Depth - 1]
	if s.Depth > 2 && IsDoubleLong(s.Layers[s.Depth - 2]) {
		layer = s.Layers[s.Depth - 2]
	}
	return layer, nil
}

func (s *Frame)Clean() (error) {
	s.Layers = s.Layers[0:0]
	s.Depth = 0
	return nil
}


func (s *Aborigines)GetAborigines(i uint32) (interface{}, error) {
	return s.Layers[i], nil
}

func (s *Aborigines)SetAborigines(i uint32, l interface{}) (error) {
	
	if i >= 1 && IsDoubleLong(s.Layers[i - 1].(string)) {
		panic("set local variable double_long low byte.")
	}
	if IsDoubleLong(l) {
		s.Layers[i + 1] = "double_long"
	}
	s.Layers[i] = l
	return nil
}

func IsDoubleLong(l interface{}) bool {
	to := reflect.TypeOf(l)
	return to == reflect.TypeOf(types.Jdouble(0)) ||
		to == reflect.TypeOf(types.JDO) ||
		to == reflect.TypeOf(types.JDU) ||
		to == reflect.TypeOf(types.JDN) ||
		to == reflect.TypeOf(types.Jlong(0))
}

/*
func (s *Frame)PushFrameX2(l interface{}) error {
	s.Layers = append(s.Layers, l)
	s.Layers = append(s.Layers, "double_long")
	s.Depth += 2
	return nil
}

func (s *Frame)PopFrameX2(reflect.Type) (interface{}, error) {
	s.Depth -= 2
	layer := s.Layers[s.Depth]
	s.Layers = s.Layers[0 : s.Depth]
	return layer, nil
}

func (s *Aborigines)GetAboriginesX2(i uint32) (interface{}, error) {
	if s.Layers[i + 1] != "double_long" {
		panic("get double_long err.")
	}
	return s.Layers[i], nil
}

func (s *Aborigines)SetAboriginesX2(i uint32, l interface{}) (error) {
	s.Layers[i] = l
	s.Layers[i + 1] = "double_long"
	return nil
}
*/

