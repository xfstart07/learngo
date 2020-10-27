// Author: xufei
// Date: 2019-07-31

package fibonacci

var F = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

func Search(arr []int, n, key int) int {
	low := 1
	high := n

	// 计算 n 位于斐波那契数列的位置
	k := 0
	for n > F[k]-1 {
		k++
	}

	// 将不满的数值补全
	for i := n; i < F[k]-1; i++ {
		arr = append(arr, arr[n])
	}

	// 目前数值是 arr[0...(F[k]-1)-1]
	for low <= high {
		mid := low + F[k-1] - 1 // 计算当前分隔的下标
		if key < arr[mid] {
			high = mid - 1
			k = k - 1 // 斐波那契下标减一位
		} else if key > arr[mid] {
			low = mid + 1
			k = k - 2 // 斐波那契下标减 2 位
		} else {
			if mid <= n {
				return mid
			} else {
				return n // 若 mid>n 说明是补全数值，返回 n
			}
		}
	}

	return 0
}
