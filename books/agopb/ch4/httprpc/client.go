// Author: xufei
// Date: 2019-08-22

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type result struct {
	ID     int         `json:"id"`
	Result string      `json:"result"`
	Error  interface{} `json:"error"`
}

func main() {
	params := struct {
		Method string        `json:"method"`
		Params []interface{} `json:"params"`
		ID     int           `json:"id"`
	}{
		Method: "jsonrpc/HelloService.Hello",
		Params: []interface{}{"hello"},
		ID:     1,
	}

	paramsBody, _ := json.Marshal(params)

	resp, err := http.Post("http://weixinote.dev:1234/jsonrpc", "", bytes.NewBuffer(paramsBody))
	if err != nil {
		log.Fatal(err)
	}

	var res result
	all, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(all, &res)
	fmt.Println(res)
}
