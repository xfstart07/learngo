package main

import (
	"fmt"
	"strings"
)

func main() {
	// 拆分，合并
	splitStrs := strings.Split("a,b,c", ",")
	fmt.Println(splitStrs) // [a b c]

	// n = 2 表示从第2个开始不切，= 0 返回 nil 没有切, < -1 全切
	splitnStrs := strings.SplitN("a,b,c", ",", 2)
	fmt.Println(splitnStrs) // [a b,c]

	joinStrs := strings.Join([]string{"a", "b", "c"}, ",")
	fmt.Println(joinStrs) // a,b,c

	repeatStr := strings.Repeat("na", 2)
	fmt.Println(repeatStr) // nana

	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("切割的字符串: %v\n", sl) // [The quick brown fox jumps over the lazy dog]

	// 子字符串相关
	fmt.Println("------------------子字符串相关---------------")
	s := "hello go world"
	fmt.Println(strings.HasPrefix(s, "hello")) // true
	fmt.Println(strings.HasSuffix(s, "world")) // true

	fmt.Println(strings.Contains(s, "go")) // true
	// 包含任意一个字符
	fmt.Println(strings.ContainsAny(s, "ad")) // true 包含 d

	fmt.Println(strings.Index(s, "o"))         // 4 索引位置
	fmt.Println(strings.LastIndex(s, "o"))     // 10 索引位置
	fmt.Println(strings.IndexByte(s, 'o'))     // 4
	fmt.Println(strings.LastIndexByte(s, 'o')) // 10
	fmt.Println(strings.IndexAny(s, "or"))     // 4
	fmt.Println(strings.LastIndexAny(s, "or")) // 11

	fmt.Println("------------------替换相关------------")

	fmt.Println(strings.Replace(s, "go", "to", -1)) // hello to world
	fmt.Println(strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r + 1
		}
		return r
	}, "'Twas brillig and the slithy gopher...")) // 'Txbt csjmmjh boe uif tmjuiz hpqifs...

	// 统计个数
	fmt.Println(strings.Count(s, "o")) // 3
}
