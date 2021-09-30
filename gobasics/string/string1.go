package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// http://blog.csdn.net/wangshubo1989/article/details/74557066

	// 1. 转换相关的方法
	s := "Welcome to The WORld of go!"
	fmt.Println("origin", s) // origin Welcome to The WORld of go!

	upperS := strings.ToUpper(s)
	fmt.Println("to_upper", upperS) // to_upper WELCOME TO THE WORLD OF GO!

	lowerS := strings.ToLower(s)
	fmt.Println("to_lower", lowerS) // to_lower welcome to the world of go!

	titleS := strings.ToTitle(s)
	fmt.Println("to_title", titleS) // to_title WELCOME TO THE WORLD OF GO!

	var sc unicode.SpecialCase

	fmt.Println(strings.ToUpperSpecial(sc, s))
	fmt.Println(strings.ToLowerSpecial(sc, s))
	fmt.Println(strings.ToTitleSpecial(sc, s))
	// WELCOME TO THE WORLD OF GO!
	// welcome to the world of go!
	// WELCOME TO THE WORLD OF GO!

	// 2. 比较相关的方法

	s1 := "Welcome to The WORld of go!"
	s2 := "Welcome go The WORld of go!"

	// compare 0 if a==b, -1 if a < b, and +1 if a > b.
	// 比较两个字符串大小
	fmt.Println("compare", strings.Compare(s1, s2)) // 1

	// equalFold 不区分大小写比较，返回 true, false
	fmt.Println("equalFold", strings.EqualFold("Go", "go"))  // true
	fmt.Println("equalFold", strings.EqualFold("Go", "go1")) // false

	// 3. 清理裁剪相关的方法

	st := "Goodbye, world!"

	fmt.Println(strings.Trim(st, "!"))      // Goodbye, world
	fmt.Println(strings.TrimRight(st, "!")) // Goodbye, world
	fmt.Println(strings.TrimLeft(st, "!"))  // Goodbye, world!

	fmt.Println(strings.TrimPrefix(st, "Go")) // odbye, world!
	fmt.Println(strings.TrimSuffix(st, "d!")) // Goodbye, worl

	st = " Goodbye, world!  "
	fmt.Println(strings.TrimSpace(st)) // Goodbye, world!


}
