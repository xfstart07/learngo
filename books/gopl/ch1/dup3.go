// Author: Xu Fei
// Date: 2018/8/17
package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

//run: go run main.go text
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d %s\n", n, line)
		}
	}
}