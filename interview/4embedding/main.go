// Author: xufei
// Date: 2019-10-23

package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB() // 这里调用的接收者是 People 的方法。因为 go 是组合的方式而不是集成，这里并不知道接收者 Teacher
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA() // 相当于 t.People.ShowA()
}
