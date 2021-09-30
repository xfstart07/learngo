package main

import (
	"fmt"
)

func main() {
  
  s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
  s4 := s3[3:6]

  fmt.Println(len(s4)) 
  fmt.Println(cap(s4))
}