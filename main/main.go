package main

import (
	"fmt"
	"../loader/butcher"
	"../loader/butcher/rope"
	"../diesel"
	"../loader/clazz"
		)


func main(){

	clazzName := ""
	bytes, _ := rope.ReadClass("")
	cf, _ := butcher.Decoder(bytes)
	clazz.CLASS_MAP[clazzName] = cf
	fmt.Printf("%X, %d, %d, %d\n", cf.Magic, cf.MajorVersion, cf.MinorVersion, cf.ConstantPoolCount)
	diesel.SteamCylinder()
	return
}
