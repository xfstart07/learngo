package v011

// https://leetcode-cn.com/explore/interview/card/bytedance/243/array-and-sorting/1035/
// 思路：
// nums[i] > nums[i-1] => f[i] = f[i-1]+1
// nums[i] <= nums[i-1] => f[i] = 1
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	f := make([]int, len(nums))
	max := 1
	f[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			f[i] = f[i-1] + 1
		} else {
			f[i] = 1
		}

		if f[i] > max {
			max = f[i]
		}
	}

	return max
}
