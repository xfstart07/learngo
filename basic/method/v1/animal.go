// Author: xufei
// Date: 2019-05-23

package main

import "fmt"

type Animal interface {
	Bark()
}

type Dog struct {
}

func (d Dog) Bark() {
	fmt.Println("dog")
}

type Cat struct {
}

func (c *Cat) Bark() {
	fmt.Println("cat")
}

func Bark(a Animal) {
	a.Bark()
}

func getDog() Dog {
	return Dog{}
}

func getCat() Cat {
	return Cat{}
}

func getCatPtr() *Cat {
	return &Cat{}
}

func main() {
	dogPtr := &Dog{}
	dog := Dog{}
	dogPtr.Bark() // (1) 可以运行
	dog.Bark()    // (2) 可以运行

	Bark(dogPtr) // (3) 可以运行
	Bark(dog)    // (4) 可以运行

	catPtr := &Cat{}
	cat := Cat{}
	catPtr.Bark() // (5)
	cat.Bark()    // (6) 可以调用， go 会转成 (&cat).Bark() 的方式调用，因为 cat 是可以取地址

	Bark(catPtr) // (7) 可以运行，因为 (*cat) 类型实现了 Bark 方法
	//Bark(cat) // (8) 不能运行，因为 (cat) 类型没有实现 Bark 方法

	getDog().Bark() // (9)

	//getCat().Bark()    // (10) cannot call pointer method on getCat() ,cannot take the address of getCat()
	// getCat()  无法取地址

	// (10) 解决方案，声明一个变量接住返回值，这样就能取到地址
	catt := getCat()
	catt.Bark()

	getCatPtr().Bark() // (11) 可以运行，指针类型的返回值可以取到地址

}

// 文章 https://mp.weixin.qq.com/s/US7MnIJfekJRazioxyWQhg
