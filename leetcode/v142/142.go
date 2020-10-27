// Author: xufei
// Date: 2020/4/1

package v142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	c := isCycle(head)
	if c == nil {
		return nil
	}

	n := head
	for n != c {
		n = n.Next
		c = c.Next
	}

	return n
}

func isCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return slow
		}
	}

	return nil
}

func newNode(list []int, pos int) (*ListNode, *ListNode) {
	if len(list) == 0 {
		return nil, nil
	}

	var posNode *ListNode
	head := &ListNode{Val: list[0]}
	node := head
	for i := 1; i < len(list); i++ {
		n := &ListNode{Val: list[i]}
		if i == pos {
			posNode = n
		}
		node.Next = n
		node = n
	}
	node.Next = posNode

	return head, posNode
}
