package v238

// nums[i] = 左边[i-1] * 右边[i+1]
// 计算 left[i] = left[i-1] * nums[i]
// 计算 right[i] = right[i+1] * nums[i]
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	left := make([]int, len(nums))
	right := make([]int, len(nums))

	left[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i] * left[i-1]
	}

	right[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = nums[i] * right[i+1]
	}

	for i := 0; i < len(nums); i++ {
		lv := 1
		if i > 0 {
			lv = left[i-1]
		}

		rv := 1
		if i < len(nums)-1 {
			rv = right[i+1]
		}

		ans[i] = lv * rv
	}

	return ans
}
