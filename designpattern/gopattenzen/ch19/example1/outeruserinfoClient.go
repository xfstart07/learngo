// Author: xufei
// Date: 2019-05-15

package main

import (
	"fmt"
	"learngo/books/gopattenzen/ch19/example1/outeruser"
	"learngo/books/gopattenzen/ch19/example1/userinfo"
)

func main() {
	outerUser := outeruser.NewOuterUser()
	outerUserInfo := outeruser.NewOuterUserInfo(*outerUser)

	print(outerUserInfo)
}

func print(info userinfo.IUserInfo) {
	fmt.Println(info.GetHomeAddress())
}
