// Author: xufei
// Date: 2019-08-01

package avl

// https://www.jianshu.com/p/1c07007408f3
type BitNode struct {
	data        int
	height      int
	left, right *BitNode
}

type BiTree struct {
	root *BitNode
}

func New(data int) *BitNode {
	return &BitNode{
		data:   data,
		height: 1,
	}
}
