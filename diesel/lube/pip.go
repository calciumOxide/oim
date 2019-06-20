package lube

import "../oil/types"

func AllocArray(ls ...uint32) (interface{}, error) {
	return allocArray(ls)
}

func allocArray(ls []uint32) (interface{}, error) {
	if ls == nil || len(ls) == 0 {
		return nil, nil
	}
	var a []interface{}
	for i := uint32(0); i < ls[0]; i++ {
		ar, _ := allocArray(ls[1:])
		a = append(a, ar)
	}
	return a, nil
}

func AllocObject() (*types.Jobject, error) {
	return new(types.Jobject), nil
}
