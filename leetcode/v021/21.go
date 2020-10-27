// Author: xufei
// Date: 2020/3/27

package v021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	for l1 != nil || l2 != nil {
		n, ok := compare(l1, l2)

		if tail != nil {
			tail.Next = n
		}
		tail = n
		if head == nil {
			head = n
		}

		if ok {
			l1 = l1.Next
		} else {
			l2 = l2.Next
		}
	}

	return head
}

func compare(l1, l2 *ListNode) (*ListNode, bool) {
	if l1 == nil {
		return l2, false
	}
	if l2 == nil {
		return l1, true
	}

	if l1.Val > l2.Val {
		return l2, false
	}

	return l1, true
}

func newNode(nums []int) *ListNode {
	l := &ListNode{Val: nums[0]}
	head := l
	for i := 1; i < len(nums); i++ {
		l.Next = &ListNode{Val: nums[i]}
		l = l.Next
	}

	return head
}
