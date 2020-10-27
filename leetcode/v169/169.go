// Author: xufei
// Date: 2020/3/27

package v169

import "sort"

// 用哈希表累计所有元素出现的次数，然后找出大于 n/2 的元素
// 用空间换时间
func majorityElement(nums []int) int {
	h := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		h[nums[i]]++
	}

	key := len(nums) / 2
	for k, v := range h {
		if v > key {
			return k
		}
	}

	return 0
}

//  排序，但是速度慢
func majorityElementSort(nums []int) int {
	sort.Ints(nums)

	return nums[len(nums)/2]
}
