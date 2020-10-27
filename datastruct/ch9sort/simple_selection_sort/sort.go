// Author: xufei
// Date: 2019-08-05

package simple_selection_sort

// 简单选择排序
// 算法复杂度：O(N^2)
func SelectSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		min := i
		for j := i + 1; j < len(list); j++ {
			if list[min] > list[j] {
				min = j
			}
		}

		if min != i {
			list[i], list[min] = list[min], list[i] // i, min 交换
		}
	}
}
