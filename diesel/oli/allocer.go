package oli

import (
	"../../loader/clazz/item"
	"../../loader/clazz"
	"../../types"
	"reflect"
)

func AllocJobject(cf *clazz.ClassFile) *types.Jobject {
	fields := map[string]interface{}{}
	for i := uint16(0); i < cf.FieldsCount; i++ {
		field := cf.Fields[i]
		nameCp, _ := cf.GetConstant(field.NameIndex)
		fields[nameCp.Info.(*item.Utf8).Str] = nil
	}
	return &types.Jobject{
		Class: cf,
		Fileds: fields,
	}
}

func AllocJarray(size types.Jint, typer interface{}, dimension uint32, dimensionArray []types.Jint) *types.Jarray {
	var et interface{}
	//if reflect.TypeOf(typer) == reflect.TypeOf("") {
		et = typer
	//} else
	if reflect.TypeOf(typer) == reflect.TypeOf(int(0)) {
		switch typer.(int) {
		case 4:
			et = "Z"
			break
		case 5:
			et = "C"
			break
		case 6:
			et = "F"
			break
		case 7:
			et = "D"
			break
		case 8:
			et = "B"
			break
		case 9:
			et = "S"
			break
		case 10:
			et = "I"
			break
		case 11:
			et = "J"
			break
		}
	}
	return &types.Jarray{
		ElementJype: et,
		Dimension: dimension,
		Reference: make([]interface{}, size),
	}
	//T_BOOLEAN		4
	//T_CHAR			5
	//T_FLOAT			6
	//T_DOUBLE		7
	//T_BYTE			8
	//T_SHORT			9
	//T_INT			10
	//T_LONG			11
}
