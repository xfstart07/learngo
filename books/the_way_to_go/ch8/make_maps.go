package main

import (
	"fmt"
)

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.21
	mapAssigned["two"] = 3

	fmt.Println(mapLit["one"])
	fmt.Println(mapCreated["key2"])
	fmt.Println(mapLit["two"])
	fmt.Println(mapLit["ten"])
	// 1
	// 3.21
	// 3
	// 0
}
