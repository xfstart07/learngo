// Author: xufei
// Date: 2021/1/8

package imutex

import (
	"fmt"
	"testing"
)

//NOTE 模拟锁的状态变化
var state int32

const (
	mutexLocked   = 1 << iota // mutex is locked  // 锁
	mutexWoken                // 唤醒
	mutexStarving             // 饥饿
)

func TestMutexState(t *testing.T) {
	fmt.Printf("原始状态: %v\n", state)
	state |= mutexLocked
	fmt.Printf("锁住状态: %v\n", state)

	state |= mutexWoken
	fmt.Printf("唤醒状态: %v\n", state)

	state |= mutexStarving
	fmt.Printf("饥饿状态: %v\n", state)
}

//https://github.com/golang/go/blob/41d8e61a6b9d8f9db912626eb2bbc535e929fefc/src/sync/mutex.go#L106
/*
	new := old
		old 不等于 mutexStarving
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}
		// old 存在状态，并且不等于 mutexLocked 或 mutexStarving
		if old&(mutexLocked|mutexStarving) != 0 {
			new += 1 << mutexWaiterShift
		}
		是饥饿状态，并且 old 等于 mutexLocked
		if starving && old&mutexLocked != 0 {
			new |= mutexStarving
		}
*/
