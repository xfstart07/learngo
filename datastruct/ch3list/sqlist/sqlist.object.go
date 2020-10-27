// Author: xufei
// Date: 2019-07-25

package sqlist

// GetElem 获取第 i 位数据
func (list *SqList) GetElem(i int, e *int) bool {
	if list.length == 0 || i < 1 || i > list.length {
		return false
	}

	*e = list.data[i-1]
	return true
}

// ListInsert 在第 i 位 插入 e，思路：将 i+1...n 的数据向后移一位，然后将 e 放入位置 i 中。
func (list *SqList) ListInsert(i int, e int) bool {
	var k int
	if list.length == MAXSIZE {
		return false
	}
	if i < 1 || i > list.length+1 {
		return false
	}
	if i <= list.length {
		for k = list.length; k >= i-1; k-- {
			list.data[k+1] = list.data[k]
		}
	}
	list.data[i-1] = e
	list.length++

	return true
}

// ListDelete 删除线性表第 i 位元素，思路：将 i+1...n 的元素向前移一位
func (list *SqList) ListDelete(i int, e *int) bool {
	var k int
	if list.length == 0 {
		return false
	}
	if i < 1 || i > list.length {
		return false
	}

	*e = list.data[i-1]
	if i < list.length {
		for k = i; k < list.length; k++ {
			list.data[k-1] = list.data[k]
		}
	}
	list.length--
	return true
}
