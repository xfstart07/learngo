// Author: xufei
// Date: 2019-08-19

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) Print(prfix string) {
	fmt.Printf("%s:Name is %s,Age is %d", prfix, u.Name, u.Age)
}

func main() {
	u := User{Name: "张三", Age: 10}
	vu := reflect.ValueOf(u)

	mPrint := vu.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(mPrint.Call(args))
}
