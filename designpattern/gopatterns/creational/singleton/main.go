// Author: xufei
// Date: 2019/4/30

package main

import (
	"fmt"
	"learngo/designpattern/gopatterns/creational/singleton/singleton"
)

func main() {
	s1 := singleton.New()
	s1["this"] = "that"

	s2 := singleton.New()
	s2["that"] = "this"

	fmt.Println("This is", s2["this"])
}
