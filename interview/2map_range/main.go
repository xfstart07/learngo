// Author: xufei
// Date: 2019-10-23

package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() map[string]*student {
	m := make(map[string]*student)

	stus := []student{
		{Name: "zuo", Age: 20},
		{Name: "cheng", Age: 24},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus { // stu 的指针是不变的，每次遍历只是在拷贝 struct 的值
		m[stu.Name] = &stu // FIXME: 最后一个元素的地址
	}

	return m
}

func main() {
	students := pase_student()
	for k, v := range students {
		fmt.Printf("k=%s,v=%v \n", k, v)
	}
}
