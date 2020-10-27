// Author: xufei
// Date: 2019-07-26

package main

import "log"

func Fbi(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	}

	return Fbi(i-1) + Fbi(i-2)
}

func main() {
	for i := 0; i < 20; i++ {
		log.Printf("i = %d, value = %d\n", i, Fbi(i))
	}
}
