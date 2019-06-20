package utils

import "fmt"

func Log(level int, format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
