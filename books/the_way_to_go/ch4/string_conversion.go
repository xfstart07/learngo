package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("数字的长度 %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Println("转换的数字", an)

	newS = strconv.Itoa(an)
	fmt.Println("转换的字符串", newS)
}
