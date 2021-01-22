// Author: xufei
// Date: 2020/3/31

package v215

import "sort"

// 排序，第k大 = 第 (n-k+1) 小
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
