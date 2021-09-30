// Author: Xu Fei
// Date: 2018/8/29
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

var myClient *http.Client

func startWebserver() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 50)
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

	defaultRoundTripper := http.DefaultTransport
	defaultTransportPtr, ok := defaultRoundTripper.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}
	defaultTransport := *defaultTransportPtr
	defaultTransport.MaxIdleConns = 1000
	defaultTransport.MaxIdleConnsPerHost = 100

	// start a webserver in a goroutine
	startWebserver()

	for i := 0; i < 100; i++ {
		go startLoadTest()
	}

	time.Sleep(time.Second * 2400)

}
