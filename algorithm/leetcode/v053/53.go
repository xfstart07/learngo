// Author: xufei
// Date: 2020/3/30

package v053

// 贪心
// 1. 当 i + i-1 > i 时则加上 i-1
// 2. max(i)
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
