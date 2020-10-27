package main

import "fmt"

type ICar interface {
	run()
}

type Benz struct {
}

type BMW struct {
}

func (b Benz) run() {
	fmt.Println("奔驰汽车开始运行...")
}

func (b BMW) run() {
	fmt.Println("宝马汽车开始运行...")
}

type IDriver interface {
	drive()
}

type Driver struct {
}

// 接口注入
func (d Driver) drive(car ICar) {
	car.run()
}

func main() {
	zhangsan := Driver{}
	benz := Benz{}

	zhangsan.drive(benz)

	bmw := BMW{}
	zhangsan.drive(bmw)
}
