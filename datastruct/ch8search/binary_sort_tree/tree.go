// Author: xufei
// Date: 2019-08-01

package binary_sort_tree

func (t *BiTree) Search(node *BiTNode, key int, parent *BiTNode) (*BiTNode, bool) {
	if t.root == nil {
		return nil, false
	}
	if node == nil {
		return parent, false
	}
	if node.data == key {
		return node, true
	}

	if node.data > key {
		return t.Search(node.left, key, node)
	} else {
		return t.Search(node.right, key, node)
	}
}

func (t *BiTree) Insert(key int) bool {
	parent, ok := t.Search(t.root, key, nil)
	if ok {
		return false
	}

	node := new(BiTNode)
	node.data = key

	if t.root == nil { // 整棵树为空，插入根
		t.root = node
	} else if parent.data > key {
		parent.left = node
	} else {
		parent.right = node
	}

	return true
}

func (t *BiTree) Delete(node, parent *BiTNode, key int) bool {
	if node == nil {
		return false
	}

	if node.data == key {
		return t.DeleteNode(node, parent)
	} else if node.data > key {
		return t.Delete(node.left, node, key)
	} else {
		return t.Delete(node.right, node, key)
	}
}

func (t *BiTree) DeleteNode(node, parent *BiTNode) bool {
	if node.left == nil && node.right == nil {
		if parent == nil { // root not parent
			t.root = nil
			return true
		}
		if err := node.replaceNode(parent, nil); err != nil {
			return false
		}
		return true
	}

	if node.right == nil {
		if err := node.replaceNode(parent, node.left); err != nil {
			return false
		}
	} else if node.left == nil {
		if err := node.replaceNode(parent, node.right); err != nil {
			return false
		}
	} else {
		s := node.left
		q := node
		for ; s.right != nil; s = s.right {
			q = s
		}

		node.data = s.data
		if q != node {
			q.right = s.left
		} else {
			q.left = s.left
		}
	}

	return true
}
