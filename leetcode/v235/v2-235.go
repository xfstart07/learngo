// Author: xufei
// Date: 2020/3/25

package v235

// 二叉搜索树的特点：左子树 小于 树根，右子树 大于等于 树根
// 所有，当 p,q 小于 树根，那么就递归左子树，因为根据二叉搜索树的特点，最近公共祖先肯定在左子树
// 同样，当 p,q 大于 树根，那么递归右子树寻找
// 不然当前节点就是 p,q 的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
