package main

import (
	"log"
	"strings"
)

func main() {
	// str := "\x17\x00\x00\x00   192.168.7.27\x1f\x00\x00\x00"
	str := "\x17\x00\x00\x00   192.168.7.27\x1e\x00\x00\x00"

	if strings.Contains(str, "\x17\x00\x00\x00") && (strings.Contains(str, "\a\x00\x00\x00") || strings.Contains(str, "\x06\x00\x00\x00")) {
		log.Println("正确")
	} else {
		log.Println("不正确")
	}
}
