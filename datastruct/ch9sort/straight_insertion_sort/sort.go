// Author: xufei
// Date: 2019-08-05

package straight_insertion_sort

// 直接插入排序
// 算法复杂度：O(N^2)
func InsertSort(list []int) {
	for i := 1; i < len(list); i++ {
		if list[i] < list[i-1] {
			sentry := list[i]
			j := i - 1
			for ; j >= 0 && list[j] > sentry; j-- {
				list[j+1] = list[j]
			}
			list[j+1] = sentry
		}
	}
}
