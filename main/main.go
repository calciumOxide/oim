package main

import (
	"../diesel"
	"../loader/butcher"
	"../loader/butcher/rope"
	"fmt"
)

func main() {

	//clazzName := ""
	bytes, _ := rope.ReadClass("")
	cf, _ := butcher.Decoder(bytes)
	//binary.CLASS_MAP[clazzName] = cf
	fmt.Printf("%X, %d, %d, %d\n", cf.Magic, cf.MajorVersion, cf.MinorVersion, cf.ConstantPoolCount)
	diesel.SteamCylinder()
	return
}
