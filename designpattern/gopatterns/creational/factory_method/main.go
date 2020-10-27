// Author: xufei
// Date: 2019/4/30

package main

import (
	"learngo/designpattern/gopatterns/creational/factory_method/data"
	"os"
)

func main() {
	dir := os.TempDir()
	_, err := os.Create(dir + "temp.txt")
	if err != nil {
		panic(err)
	}

	store := data.NewStore(data.MemoryStorage)
	r, err := store.Open(dir + "temp.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	_, err = r.Write([]byte("hello world"))
	if err != nil {
		panic(err)
	}
}
