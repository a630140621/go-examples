package chapter1

import (
	"os"
	"strings"
)

// Echo2 使用 string.Join 
func Echo2() {
	strings.Join(os.Args, " ")
}
