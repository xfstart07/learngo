// Author: xufei
// Date: 2019-07-26

package infix2suffix

type StackNode struct {
	data string
	next *StackNode
}

type LinkStack struct {
	top   *StackNode
	count int
}

type Expression struct {
	exps, ans []string
	link      *LinkStack
}
