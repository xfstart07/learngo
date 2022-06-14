package main

import "fmt"

type Benz struct {
}

func (b Benz) run() {
	fmt.Println("奔驰汽车开始运行...")
}

type Driver struct{}

func (d Driver) drive(benz Benz) {
	benz.run()
}

func main() {
	zhangsan := Driver{}
	benz := Benz{}

	zhangsan.drive(benz)
}

/*
紧耦合例子
*/
