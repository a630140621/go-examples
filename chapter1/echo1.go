package chapter1

import (
	"os"
)

// Echo1 使用 range
func Echo1() {
	// @description 使用
	var s, sep string
	for _, command := range os.Args {
		s += sep + command
		sep = " "
	}
}
