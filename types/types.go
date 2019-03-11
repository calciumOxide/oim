package types

import "sync"

/**
字符		类型				含义
B		byte			有符号字节型数
C		char			Unicode	字符，UTF-16	编码
D		double			双精度浮点数
F		float			单精度浮点数
I		int				整型数
J		long			长整数
S		short			有符号短整数
Z		boolean			布尔值	true/false
L		Classname;		reference		一个名为<Classname>的实例
[		reference		一个一维数组
*/

type Jbyte int8

type Jchar uint8

type Jdouble  float64
type jdoubleO float64
var  JDO = jdoubleO(1.7976931348623157e+308)
type jdoubleU float64
var  JDU = jdoubleU(-1.7976931348623157e+308)
type jdoubleN float64
var  JDN = jdoubleN(1.4E-45)

type Jfloat  float32
type jfloatO float32
var  JFO = jfloatO(3.4028235e+38)
type jfloatU float32
var  JFU = jfloatU(-3.4028235e+38)
type jfloatN float32
var  JFN = jfloatN(1.4e-45)

type Jint int32

type Jlong int64

type Jshort int16

type Jboolean bool

type Jaddress uint32

type Jreference struct {
	ElementType interface{}
	Reference   interface{}
	threadId int
	monitorCount int
	rwLock sync.RWMutex
}

func (s *Jreference) MonitorEnter(threadId int) bool {
	wait:
	s.rwLock.Lock()
	if s.threadId != 0 && s.threadId != threadId {
		s.rwLock.Unlock()
		goto wait
	}
	s.monitorCount++
	s.threadId = threadId
	s.rwLock.Unlock()
	return true
}

func (s *Jreference) MonitorExit(threadId int) bool {
	if s.threadId != threadId || s.monitorCount == 0 {
		return false
	}
	s.rwLock.Lock()
	s.monitorCount--
	if s.monitorCount == 0 {
		s.threadId = 0
	}
	s.rwLock.Unlock()
	return true
}

type Jarray struct {
	Dimension   uint32
	ElementJype interface{}
	Reference   interface{}
}

type Jobject struct {
	Class interface{}
	Fileds map[string]interface{}
}
