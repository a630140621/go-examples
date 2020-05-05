// 排序
package main

import (
	"fmt"
	"sort"
)

type Product struct {
	no    int // 商品编号
	price float64 // 商品价格
	name  string // 商品名称
}

func (p *Product) String() string {
	return fmt.Sprintf("no. %d price %f name %v,", p.no, p.price, p.name)
}


type byNo []*Product

func (b byNo) Len() int {
	return len(b)
}

func (b byNo) Less(i, j int) bool {
	return b[i].no < b[j].no
}

func (b byNo) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	products := []*Product{
		&Product{
			no: 1,
			price: 5,
			name: "apple",
		},
		&Product{
			no: 3,
			price: 2,
			name: "bannane",
		},
		&Product{
			no: 2,
			price: 1,
			name: "pea",
		},
	}

	fmt.Println(products)
	sort.Sort(byNo(products))
	fmt.Println(products)
	sort.Sort(sort.Reverse(byNo(products)))
	fmt.Println(products)
}
