// Author: xufei
// Date: 2019-07-26

package stack_link

type StackNode struct {
	data int
	next *StackNode
}

type LinkStack struct {
	top   *StackNode
	count int
}
