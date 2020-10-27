// Author: xufei
// Date: 2020/4/1

package v236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 暴力递归
func lowestCommonAncestorV1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == p.Val && root.Val == q.Val {
		return root
	}

	ans, _ := FindP(root, p, q)

	return ans
}

// 1. 首先找到 p
// 2. 然后搜索子树是否有q，找到则标记然后返回
// 3. 没有则返回到父节点，重复第 2 步
func FindP(root, p, q *TreeNode) (ans *TreeNode, ok bool) {
	if root == nil {
		return nil, false
	}

	if root.Val == p.Val {
		okq := searchNode(root, q)
		if okq {
			return root, true
		}
		return nil, true
	}

	ansl, okl := FindP(root.Left, p, q)
	// 在左边已经找到
	if ansl != nil && okl {
		return ansl, okl
	}

	// 左边找到了p，但没有找到 q，那么在右边找 q
	if okl {
		okq := searchNode(root, q)
		if okq {
			return root, true
		}
		ok = okl
	}

	ansr, okr := FindP(root.Right, p, q)
	// 在右边已经找到
	if ansr != nil && okr {
		return ansr, okr
	}
	// 在右边找到了p，但没有找到q，那么在左边找q
	if okr {
		okq := searchNode(root, q)
		if okq {
			return root, true
		}
		ok = okr
	}

	return
}

func searchNode(tree, q *TreeNode) bool {
	if tree == nil {
		return false
	}
	if tree.Val == q.Val {
		return true
	}

	return searchNode(tree.Left, q) || searchNode(tree.Right, q)
}

func NewNode(tree []int, idx int) *TreeNode {
	if idx > len(tree) || tree[idx-1] < 0 {
		return nil
	}

	node := &TreeNode{Val: tree[idx-1]}

	node.Left = NewNode(tree, idx*2)
	node.Right = NewNode(tree, idx*2+1)

	return node
}
