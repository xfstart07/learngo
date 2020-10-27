// Author: xufei
// Date: 2019-08-01

package binary_sort_tree

type BiTNode struct {
	data        int
	left, right *BiTNode
}

type BiTree struct {
	root *BiTNode
}
