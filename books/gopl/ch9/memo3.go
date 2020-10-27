// Author: Xu Fei
// Date: 2018/8/24
package main

import (
	"fmt"
	"io/ioutil"
	"learngo/books/gopl/ch9/memo1"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	m := memo1.New(httpGetBody)
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)

		go func(url string) {
			start := time.Now()
			value, err := m.GetSafe2(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))

			n.Done()
		}(url)
	}

	n.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}
