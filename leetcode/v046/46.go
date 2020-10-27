// Author: xufei
// Date: 2020/3/31

package v046

// 递归所有排列
func permute(nums []int) [][]int {
	list := make([][]int, 0)
	arrange(&list, nums, 0, len(nums)-1)
	return list
}

func arrange(list *[][]int, nums []int, cur, end int) {
	if cur == end {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*list = append(*list, tmp)
	} else {
		for i := cur; i <= end; i++ {
			nums[cur], nums[i] = nums[i], nums[cur]
			arrange(list, nums, cur+1, end)
			nums[cur], nums[i] = nums[i], nums[cur]
		}
	}
}
