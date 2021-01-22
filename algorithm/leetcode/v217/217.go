// Author: xufei
// Date: 2020/3/30

package v217

// 哈希
func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if numMap[nums[i]] {
			return true
		}
		numMap[nums[i]] = true
	}

	return false
}
