// Author: xufei
// Date: 2019/4/29

package main

import (
	hummer "learngo/books/gopattenzen/ch10/example2/hummer2"
)

func main() {
	h1 := &hummer.AHummer{
		Hummer: &hummer.H1Model{},
	}
	h1.Run()

	h2 := &hummer.AHummer{
		Hummer: &hummer.H2Model{},
	}
	h2.Run()
}
