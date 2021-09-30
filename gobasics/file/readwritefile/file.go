// Author: xufei
// Date: 30 Apr 2019

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/a1/Dropbox/Code/gopath/src/learngo/basic/file/readwritefile/songs")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	write, err := os.Create("/Users/a1/Dropbox/Code/gopath/src/learngo/basic/file/readwritefile/write.json")
	if err != nil {
		panic(err)
	}
	defer write.Close()

	//all, _ := ioutil.ReadAll(file)
	//fmt.Println(string(all))

	songs := []string{}
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		songs = append(songs, string(line))
	}
	fmt.Println(songs)
	bytes, _ := json.Marshal(songs)
	_, e := write.Write(bytes)
	if e != nil {
		panic(e)
	}
}
