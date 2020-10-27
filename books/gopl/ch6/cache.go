// Author: Xu Fei
// Date: 2018/8/21
package main

import (
	"sync"
	"fmt"
	"time"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	cache.mapping["key1"] = "value1"
	cache.mapping["key2"] = "value2"
	go func() {
		fmt.Println(Lookup("key1"))
	}()
	go func() {
		fmt.Println(Lookup("key2"))
	}()

	time.Sleep(3*time.Second)
}
