// Author: xufei
// Date: 2020/11/13

package v739

func dailyTemperatures(T []int) []int {
	stack := []int{0}
	num := make([]int, len(T))
	for i := 1; i < len(T); i++ {
		for {
			if len(stack) == 0 {
				break
			}

			j := len(stack) - 1 // 头部值
			if T[i] > T[stack[j]] {
				num[stack[j]] = i - stack[j]
				if j > 0 {
					j = j - 1
				}
				stack = stack[:j] // 退出栈
			} else {
				stack = append(stack, i) // 加入栈
				break
			}
		}
		stack = append(stack, i)
	}

	return num
}
