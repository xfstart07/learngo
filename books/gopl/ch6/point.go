// Author: Xu Fei
// Date: 2018/8/21
package main

import "fmt"

type Point struct {
	X,Y float64
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p,q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i:= range path {
		// 相同调用是：path[i].Add(offset)
		path[i] = op(path[i], offset)
	}

}

func main() {
	path := Path{Point{X:1, Y:2}, Point{X:2, Y:2}}

	path.TranslateBy(Point{X: 1, Y:2}, true)
	path.TranslateBy(Point{X: 1, Y:2}, false)

	fmt.Println(path)
}