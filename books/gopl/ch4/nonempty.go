// Author: Xu Fei
// Date: 2018/8/20
package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _,s:= range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
}