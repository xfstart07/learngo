// Author: Xu Fei
// Date: 2018/8/20
package main

import (
	"sort"
	"fmt"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s,"")
}

func compare(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}

	sort1 := sortString(s1)
	sort2 := sortString(s2)

	return strings.Compare(sort1, sort2)
}

func main() {
	s1 := "qwerty"
	s2 := "ytrewq"

	fmt.Println(compare(s1, s2))
}
