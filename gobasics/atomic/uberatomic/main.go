package main

import (
	"encoding/json"
	"fmt"

	"go.uber.org/atomic"
)

type user struct {
	Age atomic.Int32 `json:"age"`
}

func main() {
	u := user{}
	u.Age.Store(100)

	body, _ := json.Marshal(u)
	fmt.Println(string(body))
}
