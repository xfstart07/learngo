// Author: xufei
// Date: 2020/10/21

package main

import "fmt"

type User struct {
	Name string `json:"name"`
}

// struct 可以作为 map key 使用
var hash = map[interface{}]bool{
	User{}: true,
}

func main() {
	got, ok := hash[User{}]
	if !ok {
		fmt.Println("不存在")
	}
	fmt.Println(got)
}
