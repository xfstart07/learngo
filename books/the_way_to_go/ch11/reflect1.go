package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type: ", reflect.TypeOf(x)) // 通过反射接口返回类型
	v := reflect.ValueOf(x)                  // 通过反射接口返回值
	fmt.Println("value: ", v)

	fmt.Println("reflect type:", v.Type())    // 反射值的类型
	fmt.Println("reflect kind:", v.Kind())    // 反射值的类型种类
	fmt.Println("reflect value: ", v.Float()) // 反射值的具体类型值

	fmt.Println(v.Interface()) // 返回底层值
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
