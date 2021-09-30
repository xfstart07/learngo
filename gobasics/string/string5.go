package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	person := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "Leon",
		Age:  27,
	}

	ss, _ := json.Marshal(person.Name)
	fmt.Println(string(ss))

	list := []string{"key1", "value1", "key2", "value2"}
	str := strings.Join(list, "=")
	fmt.Println(str)
}
