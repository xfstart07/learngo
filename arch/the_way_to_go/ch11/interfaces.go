package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
	// Sum() float32  // 添加这个接口方法，会报下面的错
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sql := new(Square)
	sql.side = 5

	var area Shaper
	// ERROR cannot use sql (type *Square) as type Shaper in assignment:
	// *Square does not implement Shaper (missing Sum method)
	area = sql

	fmt.Println("area: ", area.Area())
}
