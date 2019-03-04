package main

import (
	"fmt"
			"../diesel"
	"../loader/clazz"
		)


func main(){

	clazzName := "com/oxide/A"
	cf := clazz.GetClass(clazzName)
	fmt.Printf("%X, %d, %d, %d\n", cf.Magic, cf.MajorVersion, cf.MinorVersion, cf.ConstantPoolCount)
	diesel.SteamCylinder()
	return
}
