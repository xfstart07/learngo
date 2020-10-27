// Author: xufei
// Date: 2019-08-05

package merge_sort

// MergeSort 递归版归并排序
// 算法复杂度：O(N+log^N)
func MergeSort(list []int) []int {
	return MSort(list, 0, len(list)-1)
}

// 递归排序
func MSort(list []int, s, t int) []int {
	if s == t {
		return []int{list[s]}
	} else {
		mid := (s + t) / 2
		tr1 := MSort(list, s, mid)   // 排序左边
		tr2 := MSort(list, mid+1, t) // 排序右边

		return Merge(tr1, tr2) // 两两归并
	}
}

// 合并两个有序数组
func Merge(tr1, tr2 []int) []int {
	tr := make([]int, len(tr1)+len(tr2))

	var i, j, k int
	for ; i < len(tr1) && j < len(tr2); k++ {
		if tr1[i] < tr2[j] {
			tr[k] = tr1[i]
			i++
		} else {
			tr[k] = tr2[j]
			j++
		}
	}

	if i < len(tr1) {
		for ; i < len(tr1); i++ {
			tr[k] = tr1[i]
			k++
		}
	}

	if j < len(tr2) {
		for ; j < len(tr2); j++ {
			tr[k] = tr2[j]
			k++
		}
	}

	return tr
}
