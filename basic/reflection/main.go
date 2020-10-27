// Author: xufei
// Date: 2019-08-19

package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Name string
	Age  int
	sex  string
}

func (u user) GetName() string {
	return u.Name
}

func main() {
	u := user{Name: "张三", Age: 10}
	tu := reflect.TypeOf(u)
	fmt.Println(tu)

	vu := reflect.ValueOf(&u)
	fmt.Println(vu)

	fmt.Println(vu.Type())

	fmt.Println(tu.Kind())

	for i := 0; i < tu.NumField(); i++ {
		fmt.Println(tu.Field(i).Name)
	}

	fmt.Println("NumMethod，大写导出的方法数", tu.NumMethod())
	for i := 0; i < tu.NumMethod(); i++ {
		fmt.Println(tu.Method(i).Name)
	}

	// 修改结构体字段值
	//vu := reflect.ValueOf(&u) 需要取地址才能行
	eu := vu.Elem()
	fmt.Println("元素是什么", eu)
	fmt.Println("私有字段能否修改:", eu.FieldByName("sex").CanSet())
	fmt.Println("公有字段能否修改:", eu.FieldByName("Name").CanSet())
	eu.FieldByName("Name").SetString("李四")
	fmt.Println("打印修改后值：", vu)
}

//main.user
//{张三 10}
//main.user
//struct
