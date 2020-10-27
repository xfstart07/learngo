// Author: xufei
// Date: 2020/4/1

package v148

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：将链表转为数组，然后排序，在转为链表
// NOTE 优化: 归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := make([]int, 0)
	for n := head; n != nil; n = n.Next {
		list = append(list, n.Val)
	}
	sort.Ints(list)

	sortH := &ListNode{Val: list[0]}
	node := sortH
	for i := 1; i < len(list); i++ {
		node.Next = &ListNode{Val: list[i]}
		node = node.Next
	}

	return sortH
}

func newNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}

	head := &ListNode{Val: list[0]}
	node := head
	for i := 0; i < len(list); i++ {
		node.Next = &ListNode{Val: list[i]}
		node = node.Next
	}

	return head
}
