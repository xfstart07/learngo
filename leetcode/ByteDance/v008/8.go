package v008

import "sort"

// 首先将 nums 排序，形成一个有序数组。
// 枚举第一个数字 i，那么剩下两个值相加就要等于 -nums[i]；
// 枚举剩下两个数字j，k，因为数组已经有序，那么j，k一定是一小一大的两个数字。
// 移动的原则：
// 1. 当 j+k > target 时，k 移动
// 2. 当 j+k < target 时，j 移动
// 3. 当 j+k = target 时，记录组合，移动 j，k，并且判断重复性（j == j-1, k == k+1）。
func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) <= 2 {
		return ans
	}

	// 将数组排序，这样就可以枚举数组中的每个值
	// 枚举 i，剩下的目标值就是 -nums[i]
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		// 当 i == i-1，那么不能重复使用，i 移动
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[j]+nums[k] == target {
				val := []int{nums[i], nums[j], nums[k]}
				ans = append(ans, val)

				// 移动
				// 当 j == j-1，那么不能重复使用，j 移动
				// 当 k == k+1，那么不能重复使用，k 移动
				j++
				k--
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[j]+nums[k] > target {
				// 当相加大于目标值，那么k移动
				k--
			} else {
				// 当相加小于目标值，那么j移动
				j++
			}
		}
	}

	return ans
}
