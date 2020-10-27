package main

import (
	"fmt"
)

type TowInts struct {
	a int
	b int
}

func main() {
	two1 := new(TowInts) // &TowInts{}
	two1.a = 12
	two1.b = 10

	fmt.Println("和: ", two1.AddThem())
	fmt.Println("参数和: ", two1.AddToParam(20))

	two2 := TowInts{3, 4}
	fmt.Println("和", two2.AddThem())
}

func (tn *TowInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TowInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
