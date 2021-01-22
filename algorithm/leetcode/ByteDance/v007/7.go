package v007

import (
	"strconv"
	"strings"
)

// 回溯，递归枚举
func restoreIpAddresses(s string) []string {
	ans := dfs(s, 4, []string{})
	return ans
}

// 1. s 不能大于 255
// 2. s 只有长度等于1市首位可以为0
func valid(s string) bool {
	m := len(s)
	if s[0] == '0' {
		return m == 1
	}
	dig, _ := strconv.Atoi(s)
	return dig <= 255
}

func dfs(s string, length int, ip []string) []string {
	if len(s) == 0 {
		return nil
	}

	if length == 1 {
		if !valid(s) {
			return nil
		}

		ip = append(ip, s)
		return []string{strings.Join(ip, ".")}
	}

	var ans []string

	k := ""
	sLen := len(s)
	if sLen > 3 {
		sLen = 3
	}
	for i := 0; i < sLen; i++ {
		k = k + string(s[i])
		if !valid(k) {
			break
		}

		tmp := append(ip, k)
		list := dfs(s[i+1:], length-1, tmp)
		if len(list) > 0 {
			ans = append(ans, list...)
		}
	}

	return ans
}
