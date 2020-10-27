package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// str := `\xcds\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x04K123\x00\x00\x01\x00\x01`
	str := `9\x15\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x04K205\x00\x00\x01\x00\x01`

	// strPatten := `x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x03`
	strPatten := `x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x04`
	endPatten := `\x00\x00\x01\x00\x01`
	fmt.Println(strings.Index(str, strPatten))
	fmt.Println(strings.Index(str, strPatten) + len(strPatten))
	fmt.Println(strings.Index(str, endPatten))

	startIndex := strings.Index(str, strPatten) + len(strPatten)
	endIndex := strings.Index(str, endPatten)
	fmt.Println(str[startIndex:endIndex])

	clientIP := ""
	fmt.Println("包厢clientIP", clientIP)
	ipSplit := strings.Split(clientIP, ".")
	ipSegment := ""
	if len(ipSplit) > 1 {
		ipSegment = strings.Join(ipSplit[:3], ".")
	}
	fmt.Println("网段: ", ipSegment)

	s1 := `F\xd9\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x04K121\x00\x00\x01\x00\x01`
	s2 := `\xdar\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x04K125\x00\x00\x01\x00\x01`
	fmt.Println(len(s1), len(s2))

	sp := `\x00\x00\x01\x00\x01`
	spstr := `<\v\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x06isatap\x00\x00\x01\x00\x01`
	fmt.Println(strings.HasSuffix(spstr, sp))

	strPatten4 := "\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x04"
	fmt.Println(len(strPatten4))

	fmt.Println(time.Now().Unix())

	fmt.Printf("%q\n", fmt.Sprintf("%X", "\xf3\xbc"))
	fmt.Println(strconv.ParseInt("65", 16, 32))
	fmt.Println(strconv.ParseInt("52", 16, 32))

}
