// Author: xufei
// Date: 2020/3/31

package v088

import "sort"

// 拷贝 num2数据到 nums1[m:] 中，然后排序
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
