// Author: xufei
// Date: 2020/3/24

package v104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归深入，计算出左右深度后取最大值
// 注意边界值，树的深度可以为0
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	l := maxDepth(root.Left) + 1
	r := maxDepth(root.Right) + 1
	if l >= r {
		return l
	}

	return r
}

func NewNode(tree []int, n int) *TreeNode {
	if n > len(tree) || tree[n-1] <= 0 {
		return nil
	}

	node := &TreeNode{Val: tree[n-1]}

	node.Left = NewNode(tree, n*2)
	node.Right = NewNode(tree, n*2+1)

	return node
}
