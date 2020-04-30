package chapter2

import "fmt"

var column = 5
var row = 5
var table [][]int

// 初始化一个 5行5列的 table
func init() {
	fmt.Println("first init")
	for i := 0; i < column; i++ {
		table = append(table, []int{})
		for j := 0; j < row; j++ {
			table[i] = append(table[i], i*j)
		}
	}
}

func init() {
	fmt.Println("second init")
}

var table2 [][]int = func() (table [][]int) {
	fmt.Println("init table2 with anonymous func")
	for i := 0; i < column; i++ {
		table = append(table, []int{})
		for j := 0; j < row; j++ {
			table[i] = append(table[i], i*j)
		}
	}
	return
}()

func printTable() {
	fmt.Println(table)
	fmt.Println(table2)
}
