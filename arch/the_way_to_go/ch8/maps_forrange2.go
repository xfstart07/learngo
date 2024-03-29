package main

import (
	"fmt"
)

func main() {
	// Version A
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Println("Version A", items)

	// Version B: Not Good!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Println("Version B", items2)
}
