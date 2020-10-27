// Author: xufei
// Date: 2019-05-14

package main

import "fmt"

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (o *Addition) Apply(leftValue, rightValue int) int {
	return leftValue + rightValue
}

type Multiplication struct{}

func (o Multiplication) Apply(leftValue, rightValue int) int {
	return leftValue * rightValue
}

func main() {
	add := Operation{&Addition{}}
	fmt.Println("Add:", add.Operate(3, 5))

	mult := Operation{Multiplication{}}
	fmt.Println("Mult: ", mult.Operate(3, 5))
}
