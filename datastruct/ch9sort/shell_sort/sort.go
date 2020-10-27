// Author: xufei
// Date: 2019-08-05

package shell_sort

// 希尔排序
// 算法复杂度：O(N^(3/2))
// 讲解: https://www.youtube.com/watch?v=yJjjkfnduiI
func ShellSort(list []int) {
	for inc := len(list) - 1; inc > 1; {
		inc = inc/3 + 1 // 增量序列

		for i := inc + 1; i <= len(list)-1; i++ {
			if list[i] < list[i-inc] {
				list[0] = list[i]

				j := i - inc
				for ; j > 0 && list[0] < list[j]; j -= inc {
					list[j+inc] = list[j]
				}
				list[j+inc] = list[0]
			}
		}
	}
}
