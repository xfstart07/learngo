package main

import (
	"fmt"
)

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func main() {
	var val Any
	val = 5
	fmt.Printf("val 的值 %v\n", val)

	son := new(Person)
	son.name = "Leon"
	son.age = 27
	val = son
	fmt.Printf("val 的值 %v\n", val)

	switch t := val.(type) {
	case int:
		fmt.Printf("int 类型 %T\n", t)
	case string:
		fmt.Printf("string 类型 %T\n", t)
	case bool:
		fmt.Printf("bool 类型 %T\n", t)
	case *Person:
		fmt.Printf("Person 指针类型 %T\n", t)
	default:
		fmt.Printf("未知类型%T\n", t)
	}
}
