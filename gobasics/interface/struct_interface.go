package main

import (
	"encoding/json"
	"fmt"
)

type IBook interface {
	getName() string
}

type Book struct {
	name string
}

func (b *Book) getName() string {
	return b.name
}

// func (b *Book) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(b)
// }

func main() {
	// var b Book
	// 做类型断言，值必须是一个接口值
	// b = &Book{name: "设计模式"}
	// fmt.Println(b.getName())

	// ib, ok := b.(IBook)
	// fmt.Println(ib, ok)

	var b IBook
	b = &Book{name: "设计模式"}
	fmt.Println(b.getName())

	ib, ok := b.(IBook)
	fmt.Println(ib, ok) // &{设计模式} true

	_, ok = b.(json.Marshaler)
	fmt.Println(ok) // false
}
