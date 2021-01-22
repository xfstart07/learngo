package v016

import (
	"math"
	"sort"
)

// 排序+双指针枚举
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// abs(target-sum) < abs(target-ans)，表明 sum 更加接近 target
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-ans)) {
				ans = sum
			}

			if sum > target {
				k--
			} else if sum < target {
				j++
			} else { // sum == target
				return ans
			}
		}
	}

	return ans
}
