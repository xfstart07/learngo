package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Println("字符串字节长度", len(str))
	fmt.Println("字符串字符个数", utf8.RuneCountInString(str))

	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Println("字符串字节长度2", len(str2))
	fmt.Println("字符串字符个数", utf8.RuneCountInString(str2))

}
