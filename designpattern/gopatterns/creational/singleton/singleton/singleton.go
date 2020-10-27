// Author: xufei
// Date: 2019/4/30

package singleton

import "sync"

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}
