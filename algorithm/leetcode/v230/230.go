// Author: xufei
// Date: 2020/3/31

package v230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将搜索树按照前序遍历从小到大一次，将值加入到数组中，这样list[k]就是第k小值。
func kthSmallest(root *TreeNode, k int) int {
	list := make([]int, 0)
	preorder(root, &list)
	return list[k-1]
}

func preorder(node *TreeNode, list *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		preorder(node.Left, list)
	}
	*list = append(*list, node.Val)
	if node.Right != nil {
		preorder(node.Right, list)
	}
}

func newTree(tree []int, n int) *TreeNode {
	if n > len(tree) || tree[n-1] < 0 {
		return nil
	}

	node := &TreeNode{Val: tree[n-1]}

	node.Left = newTree(tree, n*2)
	node.Right = newTree(tree, n*2+1)

	return node
}
