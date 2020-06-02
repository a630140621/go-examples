package bitmap

import (
	"fmt"
	"strings"
)

type bitmap struct {
	bytes []uint32
}

func New() *bitmap {
	return &bitmap{}
}

func (b *bitmap) Add(k int) {
	byteIndex := k / 32
	bitIndex := k % 32
	for len(b.bytes) < byteIndex+1 {
		b.bytes = append(b.bytes, 0)
	}

	b.bytes[byteIndex] |= (1 << bitIndex)
}

func (b *bitmap) Exists(k int) bool {
	byteIndex := k / 32
	bitIndex := k % 32
	if len(b.bytes) < byteIndex+1 {
		return false
	}

	return (b.bytes[byteIndex] & (1 << bitIndex)) != 0
}

func (b *bitmap) String() string {
	ret := []string{}
	for _, bt := range b.bytes {
		ret = append(ret, fmt.Sprintf("%032b", bt))
	}
	return strings.Join(ret, "")
}
