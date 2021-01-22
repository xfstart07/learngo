// Author: xufei
// Date: 2020/4/2
package v033

// 循环搜索，O(n)
func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}
