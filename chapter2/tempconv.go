package chapter2

import "fmt"

// C 摄氏度
type C float64
// F 华氏度
type F float64

const (
	// ABSOLUTEZEROC 绝对零度
	ABSOLUTEZEROC C = -273.15
	// FREEZINGC 冰点
	FREEZINGC     C = 0
	// BOILINGC 沸点
	BOILINGC      C = 100
)

func (c C) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f F) String() string {
	return fmt.Sprintf("%g°C", f)
}
