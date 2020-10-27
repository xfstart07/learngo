package v010

import "fmt"

func search(nums []int, target int) int {
	// 找到旋转点
	// 判断 target 在旋转点的左边还是右边，
	// 二分查找

	mid := findMid(nums)
	fmt.Println("mid: ", mid)

	if mid+1 >= len(nums) {
		return binarySearch(nums, 0, len(nums)-1, target)
	} else if target >= nums[0] && target <= nums[mid] {
		return binarySearch(nums, 0, mid, target)
	}

	return binarySearch(nums, mid+1, len(nums)-1, target)
}

func findMid(nums []int) int {
	i := 0
	j := len(nums) - 1

	for i < j {
		mid := (i + j) / 2
		if nums[mid] > nums[mid+1] { // mid > mid+1，旋转点
			return mid
		} else if nums[mid] > nums[len(nums)-1] { // mid > len(nums) 那么 mid 在高位部分，旋转点应该在右边
			i = mid
		} else { // mid <= len(nums)，那么 mid 在低位部分，旋转点应该在左边
			j = mid
		}
	}

	return len(nums) - 1
}

func binarySearch(nums []int, star, end int, target int) int {
	fmt.Println("search:", star, " ", end)
	for star <= end {
		mid := (star + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			star = mid + 1
		}
	}

	return -1
}
