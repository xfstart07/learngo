package v005

// 动态规划
// 1. 枚举长度1..N的情况;
// 2. f[i][j] 表示长度为 i，从 j 起始的字符串是否是回文，等于 f[i-2][j+1] && s[j] == s[j+i-1]
func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	f := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		f[i] = make([]bool, len(s))
		if i == 0 || i == 1 {
			for j := 0; j < len(s); j++ {
				f[i][j] = true
			}
		}
	}

	max := 1
	sMax := s[0:1]
	for i := 2; i <= len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			x := j
			y := j + i - 1

			if f[i-2][j+1] && s[x] == s[y] {
				f[i][j] = true
			}

			if f[i][j] && i > max {
				max = i
				sMax = s[j : j+i]
			}
		}
	}

	return sMax
}
