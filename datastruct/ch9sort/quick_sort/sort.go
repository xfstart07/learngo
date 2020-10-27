// Author: xufei
// Date: 2019-08-05

package quick_sort

// 快排
// 算法复杂度: O(N*log^n)
func QSort(list []int, low, high int) {
	if low < high {
		pivot := Partition(list, low, high)

		QSort(list, low, pivot-1)
		QSort(list, pivot+1, high)
	}
}

func Partition(list []int, low, high int) int {
	pivotKey := list[low]

	for low < high {
		for low < high && list[high] >= pivotKey { // 在尾部找小于 pivotKey
			high--
		}
		list[low], list[high] = list[high], list[low] // 交换

		for low < high && list[low] <= pivotKey { // 在头部找大于 pivotKey
			low++
		}
		list[low], list[high] = list[high], list[low] // 交换
	}
	return low
}
