package loader

import (
	"./binary"
	"../loader/butcher"
	"../loader/butcher/rope"
)

func Loader(name string) *binary.ClassFile {
	bytes, _ := rope.ReadClass(name)
	cf, _ := butcher.Decoder(bytes)
	return cf
}
