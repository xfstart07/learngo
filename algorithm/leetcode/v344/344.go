// Author: xufei
// Date: 2020/3/24

package v344

// 从前后开始遍历，调换i,j位置的值，直到 i <= j
func reverseString(s []byte) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
