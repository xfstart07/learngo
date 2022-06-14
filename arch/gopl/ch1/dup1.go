// Author: Xu Fei
// Date: 2018/8/16
package main

import (
	"bufio"
	"os"
	"fmt"
)

// ch1.3
// run: go run main.go < text
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("count=%d, text=%s\n", n, line)
		}
	}

}
