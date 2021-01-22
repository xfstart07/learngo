package v003

// 思路：计算s1的26个字母数量，然后枚举s2同样长度的字符串26个字母是否和s1相等
func checkInclusion(s1 string, s2 string) bool {
	c1 := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		c1[s1[i]-'a']++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		c2 := make([]int, 26)
		for j := 0; j < len(s1); j++ {
			c2[s2[i+j]-'a']++
		}

		if check(c1, c2) {
			return true
		}
	}

	return false
}

// 判断26个字母个数相等情况
func check(c1 []int, c2 []int) bool {
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			return false
		}
	}

	return true
}
