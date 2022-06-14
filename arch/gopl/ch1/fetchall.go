// Author: Xu Fei
// Date: 2018/8/17
package main

import (
	"time"
	"os"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("总共耗时 %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cache-Control", "max-age=60")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	fmt.Println(resp.Header)
	fmt.Println(resp.Status)

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s, %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

// 1.10 的缓存练习还不清楚如何做？