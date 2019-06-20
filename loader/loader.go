package loader

import (
	"../loader/butcher"
	"../loader/butcher/rope"
	"./binary"
)

func Loader(name string) *binary.ClassFile {
	bytes, _ := rope.ReadClass(name)
	cf, _ := butcher.Decoder(bytes)
	return cf
}
