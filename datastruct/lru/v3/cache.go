// Author: xufei
// Date: 2019-08-07

package v3

import (
	"container/list"
)

type Value struct {
	value int
	ele   *list.Element
}

// O(1)
// 执行用时 : 184 ms
// 内存消耗 : 12.3 MB
type LRUCache struct {
	keys     *list.List
	hash     map[int]*Value
	capacity int // 容量
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		hash:     make(map[int]*Value),
		keys:     list.New(),
	}

	return cache
}

func (lru *LRUCache) Get(key int) int {
	v := lru.hash[key]
	if v != nil {
		lru.keys.MoveToFront(v.ele)

		return v.value
	}

	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if lru.capacity == 0 {
		return
	}

	v := lru.hash[key]
	if v != nil {
		v.value = value
		lru.keys.MoveToFront(v.ele)

		return
	}

	if lru.keys.Len() == lru.capacity {
		least := lru.keys.Back()

		key := least.Value.(int)
		lru.keys.Remove(least)
		delete(lru.hash, key)
	}

	lru.hash[key] = &Value{
		value: value,
		ele:   lru.keys.PushFront(key),
	}
}
