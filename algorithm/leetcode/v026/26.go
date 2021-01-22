// Author: xufei
// Date: 2020/3/30

package v026

// 迷惑点：在循环内操作数据，对循环本身有什么影响
// 从未部循环
func removeDuplicates(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
