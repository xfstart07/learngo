package main

import (
	"fmt"
	"reflect"
)

func main() {
	type intAlias int
	a := 1
	var b intAlias
	b = (intAlias)(a)

	fmt.Println("a type = ", reflect.TypeOf(a), ", b type=", reflect.TypeOf(b)) // a type =  int , b type= main.intAlias
}
