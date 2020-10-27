// Author: xufei
// Date: 2020/4/1

package v061

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：计算链表的断点在哪里？
// 1. len > k，那么位置是第 len-k 个
// 2. k > len，那么位置是第 len - (k%len) 个
// 知道断点后那么新链表就是断开断点，后面部分的链表的尾部接上原头部
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	// 计算长度，记录尾部
	lh := 1
	tail := head
	for tail.Next != nil {
		lh++
		tail = tail.Next
	}

	// k 是 len 的倍数，则不需要修改链表
	if (k % lh) == 0 {
		return head
	}

	// 计算断点位置
	step := lh - k
	if k > lh {
		step = lh - (k % lh)
	}

	n := head
	for n != nil {
		step--
		if step == 0 {
			break
		}
		n = n.Next
	}

	//if n == nil {
	//	return head
	//}

	newhead := n.Next
	if newhead != nil {
		n.Next = nil
	}
	tail.Next = head

	return newhead
}

func newNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}

	head := &ListNode{Val: list[0]}
	tail := head
	for i := 1; i < len(list); i++ {
		tail.Next = &ListNode{Val: list[i]}
		tail = tail.Next
	}

	return head
}
