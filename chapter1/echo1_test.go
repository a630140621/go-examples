package chapter1

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	os.Args = generateStringSlice(100)
	for i := 0; i < b.N; i++ {
		Echo1()
	}
}

func TestEcho1(t *testing.T) {
	os.Args = []string{"a", "b", "c"}
	Echo1()
}

func generateStringSlice(len int) []string {
	// 生成指定长度的 []string
	ret := make([]string, len)
	for i := 0; i < len; i++ {
		ret = append(ret, strconv.Itoa(rand.Intn(10)))
	}

	return ret
}