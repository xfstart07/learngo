// Author: xufei
// Date: 2019-08-06

package tree

const (
	MaxTreeSize = 100
)

// 结点
type Node struct {
	data   int
	parent int
}

type Tree struct {
	nodes     [MaxTreeSize]Node
	root, num int // 根位置和结点数
}

// 二叉树结点
type BiTNode struct {
	data        int
	left, right *BiTNode
}
