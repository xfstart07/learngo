// Author: xufei
// Date: 2019-08-01

package avl

func (t *BitNode) findMin() *BitNode {
	if t.left != nil {
		return t.left.findMin()
	} else {
		return t
	}
}

func (t *BitNode) Insert(key int) *BitNode {
	if t == nil {
		return New(key)
	}

	newT := t
	if t.data > key {
		t.left = t.left.Insert(key)
		newT = t.adjust()
	} else if t.data < key {
		t.right = t.right.Insert(key)
		newT = t.adjust()
	}
	newT.height = newT.MaxHeight() + 1

	return newT
}

func (t *BitNode) Delete(key int) *BitNode {
	if t == nil {
		return nil
	}

	node := t
	if t.data > key {
		t.left = t.left.Delete(key)
		node = t.adjust()
	} else if t.data < key {
		t.right = t.right.Delete(key)
		node = t.adjust()
	} else {
		if t.left != nil && t.right != nil {
			t.data = t.right.findMin().data
			t.right = t.right.Delete(t.data)
		} else if t.left != nil {
			node = t.left
		} else {
			node = t.right
		}
	}

	return node
}

func (t *BitNode) adjust() *BitNode {
	// 如果左子树的高度比右子树的高度大于2
	if t.left.getHeight()-t.right.getHeight() == 2 {
		if t.left.left.getHeight()-t.left.right.getHeight() > 0 {
			return t.RightRotate()
		} else {
			return t.LeftThenRightRotate()
		}
	} else if t.left.getHeight()-t.right.getHeight() == -2 {
		if t.right.left.getHeight()-t.right.right.getHeight() < 0 {
			return t.LeftRotate()
		} else {
			return t.RightThenLeftRotate()
		}
	}
	return t
}
