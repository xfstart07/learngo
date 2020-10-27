// Author: xufei
// Date: 2020/3/24

package v557

// 遍历字符串，遇见空格，将前面记录的字符串倒置遍历添加到结果字符串，最后判断是否还有记录的字符串未添加
func reverseWords(s string) string {
	var v []byte
	r := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			reverse(v)
			r = append(r, v...)
			r = append(r, ' ')

			v = v[:0]
		} else {
			v = append(v, s[i])
		}
	}

	if len(v) > 0 {
		reverse(v)
		r = append(r, v...)
	}

	return string(r)
}

func reverse(s []byte) {
	ls := len(s)
	for i := 0; i < ls/2; i++ {
		j := ls - i - 1
		s[i], s[j] = s[j], s[i]
	}
}
