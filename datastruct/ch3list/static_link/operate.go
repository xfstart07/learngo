// Author: xufei
// Date: 2019-07-26

package static_link

// InitList 初始化链表
// 0 表示空指针
func InitList(space *StaticLinkList) bool {
	for i := 0; i < MAXSIZE-1; i++ {
		space[i].cur = i + 1
	}
	space[MAXSIZE-1].cur = 0 // 目前静态链表为空，最后一个元素的 cur 为 0，指向表头所在位置，分配从 1 开始分配
	return true
}

// ListInsert 插入
func ListInsert(list *StaticLinkList, i, e int) bool {
	if i < 1 || i > ListLength(list)+1 {
		return false
	}

	j := MallocSLL(list)
	if j == 0 {
		return false
	}

	k := MAXSIZE - 1
	list[j].data = e
	for l := 1; l <= i-1; l++ { // 执行 i-1 次后，list[k].cur 就是第 i 个位置
		k = list[k].cur
	}

	list[j].cur = list[k].cur
	list[k].cur = j
	return true
}

// ListLength 删除
func ListDelete(list *StaticLinkList, i int) bool {
	if i < 1 || i > ListLength(list) {
		return false
	}
	k := MAXSIZE - 1
	for j := 1; j <= i-1; j++ { // 执行 i-1 次后，list[k].cur 就是第 i 个位置
		k = list[k].cur
	}
	j := list[k].cur // j 等于第 i 位置
	list[k].cur = list[j].cur
	FreeSSL(list, j) // 释放第 i 位置的空间

	return true
}

// MallocSLL 分配结点
// 若备用空间链表非空，则返回分配结点，否则返回 0
func MallocSLL(space *StaticLinkList) int {
	i := space[0].cur
	if space[0].cur > 0 { // 0 表示分配的结点无后续空间
		space[0].cur = space[i].cur
	}
	return i
}

// 释放位置的空间
func FreeSSL(list *StaticLinkList, i int) {
	list[i].cur = list[0].cur
	list[0].cur = i
}

// ListLength 记录链表长度
func ListLength(list *StaticLinkList) int {
	j := 0
	for i := list[MAXSIZE-1].cur; i > 0; i = list[i].cur {
		j++
	}

	return j
}
