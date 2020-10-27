// Author: xufei
// Date: 2019-08-05

package heap_sort

// 堆排序
// 算法复杂度：O(N*log^N)
func HeapSort(list []int) {
	for i := (len(list) - 1) / 2; i >= 0; i-- {
		adjust(list, i, len(list)-1)
	}

	for i := len(list) - 1; i > 0; i-- {
		list[0], list[i] = list[i], list[0]
		adjust(list, 0, i-1)
	}
}

// 调整：选出左右子结点大的结点位置，进行调整
func adjust(list []int, i, length int) {
	temp := list[i]

	// 从子节点开始调整
	for j := i * 2; j <= length; j *= 2 {
		if j < length && list[j] < list[j+1] {
			j++
		}

		if temp >= list[j] {
			break
		}
		list[i] = list[j]
		i = j
	}

	list[i] = temp
}
