// Author: Xu Fei
// Date: 2018/8/29
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func startWebserver() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":8080", nil)

}

func startLoadTest() {
	count := 0
	for {
		resp, err := http.Get("http://weixinote.dev:8080/")
		if err != nil {
			panic(fmt.Sprintf("Got error: %v", err))
		}
		resp.Body.Close()
		log.Printf("Finished GET request #%v", count)
		count += 1
	}

}

func main() {

	// start a webserver in a goroutine
	startWebserver()

	startLoadTest()
}
