package main

import (
	"fmt"
	"reflect"
)

// 定义一个实现了 ServeHTTP 接口的类型, http.Handler 定义了接口
// type handlerFunc func(w http.ResponseWriter, r *http.Request) error

// 实现 Handler 的接口
// func (h handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

// func homeHandler(w http.ResponseWriter, r *http.Request) error {
// 	return executeTemplate(w, "home", 200, map[string]interface{}{
// 		"Section": "home",
// 	})
// }

func main() {

	// router := mux.NewRouter()
	// 将 homeHandler 转为 handlerFunc 类型 传递进去，在请求到达的时候，去执行 homeHandler
	// router.Handle("/", handlerFunc(homeHandler))

	type Int int

	a := 1

	b := Int(a)

	// 同类型的类型转换
	fmt.Println(reflect.TypeOf(a)) // int
	fmt.Println(reflect.TypeOf(b)) // main.Int

}
