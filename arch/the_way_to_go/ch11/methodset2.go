package main

import (
	"fmt"
)

type List []int

// 定义在值上
func (l List) Len() int {
	return len(l)
}

// 定义在指针上
func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	// 空值
	var lst List
	// 调用 CountInto(lst, 1, 10) 报错
	// cannot use lst (type List) as type Appender in argument to CountInto:
	// List does not implement Appender (Append method has pointer receiver)
	if LongEnough(lst) {
		fmt.Println("- lst is long enough\n")
	}

	// 指针
	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Println("- plst is long enough\n")
	}
}
