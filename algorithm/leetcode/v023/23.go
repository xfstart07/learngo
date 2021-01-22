package v023

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路，将所有链表元素放在一个数组中，然后排序，在组合成一个链表
// length = len(lists[1])..+..len(lists[n])
// O(length) + O(length*log(length))
func mergeKLists(lists []*ListNode) *ListNode {
	nums := make([]int, 0)
	for i := 0; i < len(lists); i++ {
		for node := lists[i]; node != nil; node = node.Next {
			nums = append(nums, node.Val)
		}
	}
	sort.Ints(nums)

	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	tail := head
	for i := 1; i < len(nums); i++ {
		n := &ListNode{Val: nums[i]}
		tail.Next = n
		tail = n
	}

	return head
}
