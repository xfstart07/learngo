// Author: xufei
// Date: 2020/3/31

package v141

// 快慢指针，慢指针走一步，快指针走两步，当链表存在环时，快指针会赶上慢指针
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
