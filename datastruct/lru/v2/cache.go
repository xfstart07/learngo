// Author: xufei
// Date: 2019-08-07

package v2

import "fmt"

// 双向链表
type Node struct {
	key, value int
	prev, next *Node
}

// O(1)
// 执行用时 : 204 ms
// 内存消耗 : 26.9 MB
type LRUCache struct {
	front, rear     *Node
	hash            map[int]*Node
	count, capacity int // 容量
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		hash:     make(map[int]*Node),
		front:    &Node{key: -1, value: -1},
		rear:     &Node{key: -1, value: -1},
	}

	// 添加头尾的哨兵，减少删除的判断
	cache.front.next = cache.rear
	cache.rear.prev = cache.front

	return cache
}

func (this *LRUCache) Get(key int) int {
	node := this.hash[key]
	if node != nil {
		this.remove(node)
		this.addRear(node)

		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node := this.hash[key]
	if node != nil {
		this.remove(node)

		node.value = value
		this.addRear(node)
		return
	}

	node = &Node{key: key, value: value}
	if this.count == this.capacity {
		this.removeHead()
	}
	this.addRear(node)
}

func (this *LRUCache) print() {
	for node := this.front; node != nil; node = node.next {
		fmt.Println(node.key, node.value)
	}
}

func (this *LRUCache) addRear(node *Node) {
	this.rear.prev.next = node
	node.prev = this.rear.prev

	node.next = this.rear
	this.rear.prev = node

	this.hash[node.key] = node
	this.count++
}

func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	// 避免内存泄漏
	node.next = nil
	node.prev = nil

	delete(this.hash, node.key)
	this.count--
}

func (this *LRUCache) removeHead() {
	node := this.front.next
	this.front.next = node.next
	node.next.prev = this.front

	delete(this.hash, node.key)
	this.count--
}
