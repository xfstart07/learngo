// Author: xufei
// Date: 2019/4/30

package data

type StorageType int

// 利用 const 建立 Store 的类型
const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case DiskStorage:
		return newDiskStorage()
	case MemoryStorage:
		return newMemoryStorage()
	default:
		return newTempStorage()
	}
}
