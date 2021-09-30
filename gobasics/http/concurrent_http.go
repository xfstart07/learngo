package main

import "fmt"

func main() {
	name := "五星"
	for i := 1; i < 99; i++ {
		value := fmt.Sprintf("%v,%v,%v", i, fmt.Sprintf("K00000%02d", i), fmt.Sprintf("%s%d", name, i))
		fmt.Println(value)
	}
}
