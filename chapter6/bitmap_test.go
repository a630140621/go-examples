package chapter6

import (
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	intSet := &IntSet{}

	ok := intSet.Has(0)
	if ok {
		t.Fatalf("intSet do not have 0 element")
	}

	intSet.Add(0)
	ok = intSet.Has(0)
	if !ok {
		t.Fatalf("intSet now have 0 element")
	}

	intSet.Add(2)
	intSet.Add(32)
	intSet.Add(33)
	intSet.Add(64)
	intSet.Add(129)
	fmt.Println(intSet)
}