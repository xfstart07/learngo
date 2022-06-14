package main

import (
	"fmt"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("类型：%T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("类型 %T\n", u)
	} else {
		fmt.Println("未知类型")
	}

}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius
}
