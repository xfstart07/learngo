// Author: xufei
// Date: 2020/3/31

package v020

// 思路：使用栈数据结构保存，当栈顶匹配上则出栈
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	var stack []byte

	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if len(stack) > 0 && ((stack[len(stack)-1] == '(' && s[i] == ')') ||
			(stack[len(stack)-1] == '{' && s[i] == '}') ||
			(stack[len(stack)-1] == '[' && s[i] == ']')) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
