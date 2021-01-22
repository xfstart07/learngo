// Author: xufei
// Date: 2020/11/12

package v025

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	curr := head
	for prev != tail {
		nex := curr.Next
		curr.Next = prev
		prev = curr
		curr = nex
	}
	return tail, head
}

// 内存消耗：3M，过多
func reverseKGroup1(head *ListNode, k int) *ListNode {
	var newHead, tmp *ListNode

	if k == 0 {
		return head
	}

	// 计算长度
	length := 0
	for n := head; n != nil; n = n.Next {
		length++
	}
	if k > length {
		return head
	}

	tmp = new(ListNode)
	newHead = tmp
	nodes := make([]int, 0, k)
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node.Val)

		if len(nodes) == k {
			for i := len(nodes) - 1; i >= 0; i-- {
				tmp.Next = new(ListNode)
				tmp.Next.Val = nodes[i]
				tmp = tmp.Next
			}
			nodes = nodes[:0]
		}
	}

	for i := 0; i < len(nodes); i++ {
		tmp.Next = new(ListNode)
		tmp.Next.Val = nodes[i]
		tmp = tmp.Next
	}

	return newHead.Next
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

//func checkNodes(n1, n2 *ListNode) bool {
//
//}
