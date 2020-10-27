// Author: xufei
// Date: 2019-07-29

package suffix_expression

import (
	"fmt"
	"learngo/datastruct/ch4stack_queue/stack_link"
	"strconv"
)

// 数字进栈，遇到符号，出栈 2 个数字，进行运算后将结果入栈
func Suffix(expressions []string) (int, error) {
	link := stack_link.InitStack()

	for _, value := range expressions {
		switch value {
		case "+":
			var val1, val2 int
			if !link.Pop(&val1) {
				return 0, fmt.Errorf("expression error")
			}
			if !link.Pop(&val2) {
				return 0, fmt.Errorf("expression error")
			}

			link.Push(val1 + val2)
		case "-":
			var val1, val2 int
			if !link.Pop(&val1) {
				return 0, fmt.Errorf("expression error")
			}
			if !link.Pop(&val2) {
				return 0, fmt.Errorf("expression error")
			}

			link.Push(val1 - val2)
		case "*":
			var val1, val2 int
			if !link.Pop(&val1) {
				return 0, fmt.Errorf("expression error")
			}
			if !link.Pop(&val2) {
				return 0, fmt.Errorf("expression error")
			}

			link.Push(val1 * val2)
		case "/":
			var val1, val2 int
			if !link.Pop(&val1) {
				return 0, fmt.Errorf("expression error")
			}
			if !link.Pop(&val2) {
				return 0, fmt.Errorf("expression error")
			}

			link.Push(int(val1 / val2))
		default:
			val, err := strconv.Atoi(value)
			if err != nil {
				return 0, err
			}
			link.Push(val)
		}
	}

	var ans int
	if !link.Pop(&ans) {
		return 0, fmt.Errorf("expression error")
	}
	return ans, nil
}
