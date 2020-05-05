package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")
var dog = &Dog{}

type Dog struct {
	name string
}

func (d *Dog) String() string {
	return fmt.Sprintf("my name is %s", d.name)
}

func (d *Dog) Set(name string) error {
	d.name = name
	return nil
}

func main() {
	flag.Var(dog, "name", "enter dog name")
	flag.Parse()
	fmt.Printf("Sleep for %v...", period)
	fmt.Println()
	fmt.Println(dog)
	time.Sleep(*period)
	fmt.Println()
}
