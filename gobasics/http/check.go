package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://github.com/nats-io/go-nats/blob/master/README.md"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	for i := 0; i < 65536; i++ {
		// result, err := http.Get(url)
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("请求错误", err)
			continue
		}

		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取错误", err)
			continue
		}

		fmt.Println(string(result))
		resp.Body.Close()

	}
}
