package chapter1

import (
	"os"
	"testing"
)

func BenchmarkEcho2(b *testing.B) {
	os.Args = generateStringSlice(100)
	for i := 0; i < b.N; i++ {
		Echo2()
	}
}

func TestTEcho2(t *testing.T)  {
	os.Args = []string{"a", "b", "c"}
	Echo2()
}