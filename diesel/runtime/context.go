package runtime

import (
	"reflect"
	"../oil/types"
	"../../loader/binary"
	"../../loader/binary/attribute"
)

type Context struct {
	PC                uint32
	Code              []uint8
	CodeStack         [][]uint8
	CurrentFrame      *Frame
	FrameStack        []*Frame
	CurrentAborigines *Aborigines
	AboriginesStack   []*Aborigines
	CurrentMethod     *binary.Method
	MethodStack       []*binary.Method
	Opoos             bool
	Clazz *binary.ClassFile
	ClazzStack []*binary.ClassFile
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

func (ctx *Context)InvokeMethod(method *binary.Method, args []interface{}) *Context {
	callee := &Context{
		PC: 0,
		CurrentFrame: &Frame{
			Layers: []interface{}{},
			Depth:  0,
		},
		CurrentAborigines: &Aborigines{
			Layers: []interface{}{},
			Count:  0,
		},
		CurrentMethod: method,
	}
	if args != nil {
		for i := 0; i < len(args); i++ {
			callee.CurrentFrame.PushFrame(args[i])
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
	a, _ := ctx.CurrentMethod.GetAttribute(binary.CODE_ATTR)
	ctx.Code = a.AttributeItem.(*attribute.Codes).Code

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
	attr, _ := s.CurrentMethod.GetAttribute(binary.CODE_ATTR)
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

func IsCatched(r *types.Jreference, et *attribute.ExceptTable) bool {
	return r.Reference.(*types.Jobject).ClassTypeIndex == et.CatchType
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

