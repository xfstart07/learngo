package v001

// https://leetcode-cn.com/explore/interview/card/bytedance/242/string/1012/
// 无重复字符的最长子串
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	num := count(s)
	max := 1
	for i := 1; i < len(s); i++ {
		if i-num[i]+1 > max {
			max = i - num[i] + 1
		}

		for j := num[i] - 1; j >= 0; j-- {
			if !check(s[j+1:i+1], s[j]) {
				break
			}

			if i-j+1 > max {
				max = i - j + 1
			}
		}
	}

	return max
}

func count(s string) []int {
	num := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		j := num[i-1]
		if check(s[j:i], s[i]) {
			num[i] = num[i-1]
		} else {
			num[i] = i
		}
	}

	return num
}

func check(s string, k byte) bool {
	for i := 0; i < len(s); i++ {
		if k == s[i] {
			return false
		}
	}
	return true
}
