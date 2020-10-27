// Author: Xu Fei
// Date: 2018/7/26
package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age int
}

func (p Person) Describe() { // value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state string
	country string
}

func (a *Address) Describe() { // pointer receiver
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()

	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()

	// 总结：带值接受器的方法同时接受指针和值类型的调用

	var d2 Describer
	a := Address{"Washington", "USA"}

	// Cannot use a (type Address) as type Describe in assignment
	//d2 = a

	d2 = &a

	d2.Describe()

	// 在赋值时不能使用（Address 类型）作为 Describe 类型，Address 没有实现 Describe

}