// Author: xufei
// Date: 2020/3/30

package v160

type ListNode struct {
	Val  int
	Next *ListNode
}

// 题意：给出两个单向链表，找出两个链表中相交的起始节点, 起始节点指向同一个子链表
// 思路：
// * 创建两个指针 pA 和 pB，分别初始化为链表 A 和 B 的头结点。然后让它们向后逐结点遍历。
// * 当 pA 到达链表的尾部时，将它重定位到链表 B 的头结点 (你没看错，就是链表 B); 类似的，当 pB 到达链表的尾部时，将它重定位到链表 A 的头结点。
// * 若在某一时刻 pA 和 pB 相遇，则 pA/pB 为相交结点。
// * 想弄清楚为什么这样可行, 可以考虑以下两个链表: A={1,3,5,7,9,11} 和 B={2,4,9,11}，相交于结点 9。 由于 B.length (=4) < A.length (=6)，pB 比 pA 少经过 2 个结点，会先到达尾部。将 pB 重定向到 A 的头结点，pA 重定向到 B 的头结点后，pB 要比 pA 多走 2 个结点。因此，它们会同时到达交点。
// * 如果两个链表存在相交，它们末尾的结点必然相同。因此当 pA/pB 到达链表结尾时，记录下链表 A/B 对应的元素。若最后元素不相同，则两个链表不相交。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}

func newNode(nums []int) *ListNode {
	var head, tail *ListNode
	for i := 0; i < len(nums); i++ {
		n := &ListNode{Val: nums[i]}
		if i == 0 {
			head = n
			tail = n
		} else {
			tail.Next = n
			tail = n
		}
	}

	return head
}

/* 错误的实现

sa, sb := "", ""
	for n := headA; n != nil; n = n.Next {
		sa = strconv.Itoa(n.Val) + sa
	}
	for n := headB; n != nil; n = n.Next {
		sb = strconv.Itoa(n.Val) + sb
	}

	i, j := 0, 0
	flag := false
	for i < len(sa) && j < len(sb) {
		if sa[i] != sb[j] {
			break
		}
		flag = true
		i++
		j++
	}

	if !flag {
		return nil
	}

	headC := headA
	gap := len(sb) - len(sa)
	if gap < 0 {
		// sa如果是长的链条不需要多走
		gap = 0
	}
	for l := len(sa) - i + gap; l > 0; l-- {
		headC = headC.Next
	}

	return headC
*/
