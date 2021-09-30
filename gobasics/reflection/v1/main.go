// Author: xufei
// Date: 2020/11/25

package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

func main() {
	u := user{name: "张三", age: 10}
	tu := reflect.TypeOf(u)
	fmt.Println(tu)

	vu := reflect.ValueOf(u)
	for i := 0; i < vu.NumField(); i++ {
		fmt.Printf("fieldName: %v, Value: %v\n", vu.Field(i).Type(), vu.Field(i))
	}
}

//main.user
//fieldName: string, Value: 张三
//fieldName: int, Value: 10
