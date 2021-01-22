// Author: xufei
// Date: 2020/3/31

package v078

import "fmt"

// 思路：递归枚举各种长度的子集，元素可以选择加入或者不加入，当长度等于枚举长度则加入结果集
func subsets(nums []int) [][]int {
	list := make([][]int, 0)

	list = append(list, []int{})
	// 枚举子集长度
	for l := 1; l <= len(nums); l++ {
		subset(&list, 0, l, []int{}, nums)
	}

	fmt.Println(list)
	return list
}

func subset(list *[][]int, i, l int, sub, nums []int) {
	if l == len(sub) {
		// 加入结果集，因为slice 是引用类型，所以讲sub进行拷贝，以免修改sub时的影响
		tmp := make([]int, len(sub))
		copy(tmp, sub)
		*list = append(*list, tmp)
	}

	for j := i; j < len(nums); j++ {
		sub = append(sub, nums[j]) // 加入到子集
		subset(list, j+1, l, sub, nums)
		sub = sub[:len(sub)-1] // 不加入
	}
}
