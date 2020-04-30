package chapter2

// F2C 华氏度转摄氏度
func F2C(f F) C {
	return C((f - 32) * 5 / 9)
}

// C2F 摄氏度转华氏度
func C2F(c C) F {
	return F(c * 9 / 5 + 32)
}