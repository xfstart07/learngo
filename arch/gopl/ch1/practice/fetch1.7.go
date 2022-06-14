// Author: Xu Fei
// Date: 2018/8/17
package main

import (
	"os"
	"fmt"
	"net/http"
	"log"
	"io"
)

// 1.5
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch %v\n", err)
			os.Exit(1)
		}

		write, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close() // 关闭 Body 流，防止资源泄露
		if err != nil {
			log.Fatalf("fetch reading %v\n", err)
		}

		fmt.Printf("write-count=%d", write)
	}
}
