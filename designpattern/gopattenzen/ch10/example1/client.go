// Author: xufei
// Date: 2019/4/29

package main

import "learngo/books/gopattenzen/ch10/example1/hummer"

func main() {
	h1 := &hummer.H1Model{}
	h1.Run()

	h2 := &hummer.H2Model{}
	h2.Run()
}
