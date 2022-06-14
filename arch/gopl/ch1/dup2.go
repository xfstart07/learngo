// Author: Xu Fei
// Date: 2018/8/17
package main

import (
	"os"
	"bufio"
	"fmt"
)

// ch1: dup2
// run: go run main.go text
func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, "%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d, %s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		if counts[input.Text()] > 1 {
			fmt.Println("重复行的文件名：", f.Name())
		}
	}
}
