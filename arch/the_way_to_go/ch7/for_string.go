package main

import (
	"fmt"
)

func main() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d %c\n", i, c)
	}
}
