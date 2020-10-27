// Author: xufei
// Date: 2019/4/29

package main

import (
	hummer "learngo/books/gopattenzen/ch10/example3/hummer3"
)

func main() {
	h11 := &hummer.H1Model{}
	h11.SetAlarm(true)
	h1 := &hummer.AHummer{
		Hummer: h11,
	}
	h1.Run()

	h2 := &hummer.AHummer{
		Hummer: &hummer.H2Model{},
	}
	h2.Run()

	h3 := &hummer.AHummer{
		Hummer: &hummer.H3Model{},
	}
	h3.Run()
}
