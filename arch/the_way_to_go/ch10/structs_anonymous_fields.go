package main

import (
	"fmt"
)

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS // anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 1
	outer.in2 = 2

	fmt.Printf("outer %+v\n", outer)

	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Printf("outer2 %+v\n", outer2)

	// outer &{b:6 c:7.5 int:60 innerS:{in1:1 in2:2}}
	// outer2 {b:6 c:7.5 int:60 innerS:{in1:5 in2:10}}
}
