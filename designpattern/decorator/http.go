// Author: xufei
// Date: 2019/1/18

package main

import (
	"learngo/designpattern/decorator/http_api"

	"log"
	"net/http"
	"os"
)

func main() {
	ll := http_api.Logf(log.New(os.Stderr, "", log.LstdFlags))

	http.HandleFunc("/hello", http_api.Decorate(hello, ll, http_api.V1))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func hello(w http.ResponseWriter, req *http.Request) (interface{}, error) {
	return "ok", nil
}
