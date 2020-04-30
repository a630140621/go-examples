package chapter2

import "fmt"

var column = 5
var row = 5
var table [][]int

// 初始化一个 5行5列的 table
func init()  {
	fmt.Println("first init")
	for i := 0; i < column; i++ {
		table = append(table, []int{})
		for j := 0; j < row; j++ {
			table[i] = append(table[i], i * j)
		}
	}
}

func init() {
	fmt.Println("second init")
}

func printTable()  {
	fmt.Println(table)
}