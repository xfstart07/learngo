// Author: xufei
// Date: 2020/3/31

package v1

// 水平扫描
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var ans []byte
	for i := 0; i < len(strs[0]); i++ {
		k := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 || strs[j][i] != k {
				return string(ans)
			}
		}
		ans = append(ans, k)
	}

	return string(ans)
}
