package v004

// 题意：
// 中位数查找，1. 当两个数组的长度和是奇数，中位数等于 length/2 ，
// 2. 长度和是偶数，中位数等于 (length/2 + length/2+1)/2
// 解法：那么题意可以转化为寻找第 length/2 大的值。
// 利用二分法的思路，每次删除 k/2 个数，直到 k == 1，那么就寻找到了 第 k 值。时间复杂度 O(log(m+n))
// https://www.youtube.com/watch?v=Be-u2dLh8Aw

// 本题难点：
// 1. 利用二分法的思路，不过做过类似题目的估计也能考虑到；
// 2. 转化成查找第K大值的思路；
// 3. 利用二分法来做删除划分。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	len := len(nums1) + len(nums2)
	if len%2 == 0 {
		k1 := findKth(nums1, 0, nums2, 0, len/2) // 这里-1是因为下标是从0开始
		k2 := findKth(nums1, 0, nums2, 0, len/2+1)

		return float64(k1+k2) / 2
	} else {
		return float64(findKth(nums1, 0, nums2, 0, len/2+1))
	}
}

func findKth(nums1 []int, pos1 int, nums2 []int, pos2 int, k int) int {
	if pos1 >= len(nums1) {
		return nums2[pos2+k-1]
	}
	if pos2 >= len(nums2) {
		return nums1[pos1+k-1]
	}

	if k == 1 {
		if nums1[pos1] > nums2[pos2] {
			return nums2[pos2]
		} else {
			return nums1[pos1]
		}
	}

	di := k / 2

	item1 := 2147483647 // 设置最大值，当数组不够删时表示
	// nums1 够删
	if pos1+di <= len(nums1) {
		item1 = nums1[pos1+di-1]
	}

	item2 := 2147483647
	// nums2 不够删
	if pos2+di <= len(nums2) {
		item2 = nums2[pos2+di-1]
	}

	if item1 > item2 {
		return findKth(nums1, pos1, nums2, pos2+di, k-di)
	}

	return findKth(nums1, pos1+di, nums2, pos2, k-di)
}
