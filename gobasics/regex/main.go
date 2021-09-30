package main

import (
	"fmt"
	"regexp"
)

var reg *regexp.Regexp
var pattern string
var source string

func regexpMatch() {
	pattern = `((?:(?:25[0-5]|2[0-4]\d|(?:1\d{2}|[1-9]?\d))\.){3}(?:25[0-5]|2[0-4]\d|(?:1\d{2}|[1-9]?\d)))`
	source = "M-SEARCH * HTTP/1.1\r\nHost: [FF02::C]:1900\r\nST: uuid:898f9738-d930-4db4-a3cf-147590ce3150\r\nMan: \"ssdp:discover\"\r\nMX: 3\r\n\r\n"

	reg = regexp.MustCompile(pattern)
	allMatch := reg.FindAllString(source, -1)
	fmt.Println(allMatch)
}

func main() {
	regexpMatch()
}
