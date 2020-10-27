package v002

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		if !isCommonPrefix(strs[1:], i, strs[0][i]) {
			return strs[0][:i]
		}
	}

	return strs[0]
}

func isCommonPrefix(strs []string, idx int, k byte) bool {
	for i := 0; i < len(strs); i++ {
		if idx > len(strs[i])-1 {
			return false
		}

		if strs[i][idx] != k {
			return false
		}
	}

	return true
}
