// Author: xufei
// Date: 2020/11/16

package v239

func maxSlidingWindow(nums []int, k int) []int {
	var queue, max []int
	for i := 0; i < len(nums); i++ {
		// 超过范围的退出队列
		for len(queue) > 0 && i-queue[0] >= k {
			queue = queue[1:]
		}
		// 小于后加入的值退出队列
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)

		if i < k-1 {
			continue
		}

		max = append(max, nums[queue[0]])
	}
	return max
}
