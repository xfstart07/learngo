package main

import (
	"fmt"
)

type Employee struct {
	salary float32
}

func main() {
	e := Employee{salary: 10.1}

	fmt.Println(e.getRaise())
}

// 类型和作用在它上面定义的方法必须在同一个包里定义
// cannot define new methods on non-local type int
// func (e int) getRaise() float32 {
func (e Employee) getRaise() float32 {
	return e.salary * 1.1
}
