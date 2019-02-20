package variator

import "../lube"
import "../../types"

func AllocExcept(et ExceptType) (interface{}, error) {
	 return EXCEPT_MAP[et].(func()(*types.Jreference, error))()
}

type ExceptType uint16

var EXCEPT_MAP = make(map[ExceptType]interface{})

const (
	NullPointerException ExceptType = iota
	ArrayIndexOutOfBoundsException
	ClassCastException
	ClassNotFindException
	FieldNotFindException
	ArithmeticException
	MethodNotFindException
	AbstractMethodError
	InstructionException
)
func except(){}
func init() {

	EXCEPT_MAP[AbstractMethodError] = except


	EXCEPT_MAP[NullPointerException] = func()(*types.Jreference, error) {
		jobject, _ := lube.AllocObject()
		return &types.Jreference{
			Reference: jobject,
		}, nil
	}
	EXCEPT_MAP[ArrayIndexOutOfBoundsException] = func()(*types.Jreference, error) {
		jobject, _ := lube.AllocObject()
		return &types.Jreference{
			Reference: jobject,
		}, nil
	}
	EXCEPT_MAP[ClassCastException] = func()(*types.Jreference, error) {
		jobject, _ := lube.AllocObject()
		return &types.Jreference{
			Reference: jobject,
		}, nil
	}
	EXCEPT_MAP[ClassNotFindException] = func()(*types.Jreference, error) {
		jobject, _ := lube.AllocObject()
		return &types.Jreference{
			Reference: jobject,
		}, nil
	}

	EXCEPT_MAP[FieldNotFindException] = func()(*types.Jreference, error) {
		jobject, _ := lube.AllocObject()
		return &types.Jreference{
			Reference: jobject,
		}, nil
	}

	EXCEPT_MAP[ArithmeticException] = func()(*types.Jreference, error) {
		jobject, _ := lube.AllocObject()
		return &types.Jreference{
			Reference: jobject,
		}, nil
	}

	EXCEPT_MAP[MethodNotFindException] = func()(*types.Jreference, error) {
		jobject, _ := lube.AllocObject()
		return &types.Jreference{
			Reference: jobject,
		}, nil
	}



	EXCEPT_MAP[InstructionException] = func()(*types.Jreference, error) {
		jobject, _ := lube.AllocObject()
		return &types.Jreference{
			Reference: jobject,
		}, nil
	}
}