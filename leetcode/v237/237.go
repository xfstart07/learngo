// Author: xufei
// Date: 2020/3/24

package v237

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给定的 node 就是需要删的节点
// 思路：将删除的节点的下一个节点移到当前节点上
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
