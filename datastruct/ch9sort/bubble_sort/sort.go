// Author: xufei
// Date: 2019-08-05

package bubble_sort

// BubbleSort0 冒泡排序初级版
// 算法复杂度：O(N^2)
func BubbleSort0(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i] // i, j 交换
			}
		}
	}
}

// BubbleSort1 冒泡排序
// 思路：将最小的从尾部不断冒上来
// 算法复杂度：O(N^2)
func BubbleSort1(list []int) {
	for i := 0; i < len(list); i++ {
		for j := len(list) - 2; j >= i; j-- {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

// BubbleSort2 冒泡排序优化
// 算法复杂度：O(N^2)
func BubbleSort2(list []int) {
	flag := true
	for i := 0; i < len(list) && flag; i++ {
		flag = false
		for j := len(list) - 2; j >= i; j-- {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				flag = true
			}
		}
	}
}
