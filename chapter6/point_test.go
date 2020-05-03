package chapter6

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	p := Point{1, 1}
	q := Point{2, 2}
	fmt.Printf("%f", p.Distance(q)) // 1.4142135623730951
}

func TestType(t *testing.T) {
	p := Point{1, 1}
	q := Point{2, 2}
	fmt.Println(Point.Distance(p, q)) // 1.4142135623730951
}

func TestExpression(t *testing.T) {
	p := Point{1, 1}
	distanceP := p.Distance
	q := Point{0, 0}
	fmt.Println(distanceP(q)) // 1.4142135623730951
}
