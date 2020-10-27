package main

import (
	"fmt"
)

func main() {
	invokeF()
	invokeS()
}

func invokeF() {
	// f := Father{}

	f := Son{}
	m := map[string]string{
		"a": "1",
		"b": "2",
	}

	fmt.Println(f.Father.doSomething(m))
}

func invokeS() {
	f := Son{}
	m := map[string]interface{}{
		"a": "1",
		"b": "2",
	}

	fmt.Println(f.doSomething(m))
}

type Father struct {
}

func (f Father) doSomething(m map[string]string) []string {
	fmt.Println("父...")
	values := make([]string, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

type Son struct {
	Father
}

// 重载方法
func (s Son) doSomething(m map[string]interface{}) []string {
	fmt.Println("子...")
	values := make([]string, len(m))
	for _, v := range m {
		values = append(values, v.(string))
	}
	return values
}

/*
前置条件：执行需要什么条件
就是方法中的输入参数


后置条件：执行后需要反馈什么
就是方法中的返回数据



父类方法的输入参数是 maps[string]string，子类的输入参数是 maps[string]interface{}，
也就是说子类的输入参数类型的范围扩大了，子类代替父类传递到调用中，子类的方法永远都不会被执行。


父类的一个方法的返回值是一个类型T，子类的相同方法的返回值是S，那么里氏替换原则就要求S必须小于等于T。

子类的输入参数宽于或等于父类的输入参数

*/
