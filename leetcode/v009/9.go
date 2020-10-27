// Author: xufei
// Date: 2020/3/30

package v009

import (
	"fmt"
)

// 思路：转换成字符串，从头和尾逐一对比，注意一个数的边界值。
func isPalindrome(x int) bool {
	s := fmt.Sprintf("%v", x)
	if len(s) <= 1 {
		return true
	}

	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
