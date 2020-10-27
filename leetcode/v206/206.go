// Author: xufei
// Date: 2020/3/24

package v206

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：循环链表，当前节点的下一个节点就是目前 tail 链表，然后移动 tail
func reverseList(head *ListNode) *ListNode {
	var tail *ListNode

	for ; head != nil; head = head.Next {
		node := &ListNode{
			Val:  head.Val,
			Next: tail,
		}
		tail = node
	}

	return tail
}

// 递归, 将当前节点的下一个指向上一个节点
func reverseList2(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}

	head := reverse(cur, cur.Next)
	cur.Next = pre

	return head
}

func newNode(list []int) *ListNode {
	var head *ListNode

	for i := 0; i < len(list); i++ {
		node := &ListNode{
			Val:  list[i],
			Next: head,
		}

		head = node
	}

	return head
}
