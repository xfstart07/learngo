// Author: xufei
// Date: 2019-07-31

package binary

func Search(arr []int, key int) int {
	high := len(arr) - 1
	for low := 0; low <= high; {
		mid := (low + high) / 2
		if key < arr[mid] {
			high = mid - 1
		} else if key > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return 0
}

// 插值查找
func Interpolation(arr []int, key int) int {
	high := len(arr) - 1
	for low := 0; low <= high; {
		mid := low + (high-low)*(key-arr[low])/(arr[high]-arr[low])
		if key < arr[mid] {
			high = mid - 1
		} else if key > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return 0
}
