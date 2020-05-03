package chapter6

import (
	"bytes"
	"fmt"
)

// intSet.Add(2)
// intSet.Has(2) // true
// intSet.Has(3) // false

// 此处不使用结构体类型是专门的
type IntSet []uint

func (s *IntSet) Add(x int) {
	div, mod := x/32, uint(x%32)
	for len(*s) <= div {
		*s = append(*s, 0)
	}
	(*s)[div] |= 1 << mod // 将指定位置置为 1
}

func (s *IntSet) Has(x int) (ok bool) {
	div, mod := x/32, uint(x%32)
	if len(*s) <= div {
		return
	}

	return (*s)[div]&(1<<mod) != 0 // 指定位置是否为 1
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range *s {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 32*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
