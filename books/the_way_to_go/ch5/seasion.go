package main

import "time"
import "fmt"

func main() {
	t := time.Now()

	switch t.Month() {
	case 1, 2, 3:
		fmt.Println("春天")
	case 4, 5, 6:
		fmt.Println("夏天")
	case 7, 8, 9:
		fmt.Println("秋天")
	case 10, 11, 12:
		fmt.Println("冬天")
	}

}
