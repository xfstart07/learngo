package v015

import (
	"sort"
)

// 思路：首先把数组进行排序，然后线性枚举i，那么剩下2个j，k相加要等于-nums[i]，这样就保证i,j,k相加等于0
func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)

	if len(nums) < 3 {
		return ans
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]

		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})

				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[j]+nums[k] > target {
				k--
			} else {
				j++
			}
		}
	}

	return ans
}
