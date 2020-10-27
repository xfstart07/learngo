// Author: xufei
// Date: 2019-06-17

package p002

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：链表位数相加最大 18 ，当每位大于等于 10 时，就进一位。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := new(ListNode)
	curr := ans

	carry := 0
	for l1 != nil || l2 != nil {
		p1 := 0
		if l1 != nil {
			p1 = l1.Val
			l1 = l1.Next
		}
		p2 := 0
		if l2 != nil {
			p2 = l2.Val
			l2 = l2.Next
		}
		p := p1 + p2 + carry

		node := new(ListNode)
		if p >= 10 {
			node.Val = p - 10
			carry = 1
		} else {
			node.Val = p
			carry = 0
		}

		curr.Next = node
		curr = curr.Next
	}

	if carry == 1 {
		node := new(ListNode)
		node.Val = carry
		curr.Next = node
		curr = curr.Next
	}

	return ans.Next
}
