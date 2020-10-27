// Author: xufei
// Date: 2018/11/27

package main

import "fmt"

type User struct {
	Name string
}

// point 和 值类型 都可以访问
func (u *User) GetName() string {
	return u.Name
}

func main() {
	u := User{
		Name: "kk",
	}
	fmt.Println(u.GetName())
}
