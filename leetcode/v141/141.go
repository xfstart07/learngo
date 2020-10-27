// Author: xufei
// Date: 2020/3/31

package v141

type ListNode struct {
	Val  int
	Next *ListNode
}

// 哈希表
func hasCycleHash(head *ListNode) bool {
	hash := make(map[*ListNode]bool)
	for n := head; n != nil; n = n.Next {
		if hash[n] {
			return true
		}
		hash[n] = true
	}
	return false
}
