package v124

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	max = -9999999999
)

// 思路：
func maxPathSum(root *TreeNode) int {
	ans := -9999999999

	findMaxPathSum(root, &ans)

	return ans
}

func findMaxPathSum(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	lv := checkMax(findMaxPathSum(root.Left, max), 0)
	rv := checkMax(findMaxPathSum(root.Right, max), 0)

	// 以这个节点为一个新路径
	if root.Val+lv+rv > *max {
		*max = root.Val + lv + rv
	}

	// 选择走左子树或者右子树
	if root.Val+lv > root.Val+rv {
		return root.Val + lv
	}

	return root.Val + rv
}

func checkMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewNode(tree []int, n int) *TreeNode {
	if n > len(tree) || tree[n-1] == 0 {
		return nil
	}

	node := &TreeNode{Val: tree[n-1]}

	node.Left = NewNode(tree, n*2)
	node.Right = NewNode(tree, n*2+1)

	return node
}
