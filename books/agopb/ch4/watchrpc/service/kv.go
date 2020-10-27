// Author: xufei
// Date: 2019-08-23

package service

import (
	"fmt"
	"math/rand"
	"net/rpc"
	"sync"
	"time"
)

const (
	ServiceName = "watch/KVStoreService"
)

type StoreInterface interface {
	Get(key string, value *string) error
	Set(kv [2]string, reply *struct{}) error
	Watch(timeoutSec int, keyChanged *string) error
}

func RegisterService(srv StoreInterface) error {
	return rpc.RegisterName(ServiceName, srv)
}

type KVStoreService struct {
	m      map[string]string           // kv 数据
	filter map[string]func(key string) // 过滤器函数
	mu     sync.Mutex                  // 互斥锁
}

func (s *KVStoreService) Watch(timeoutSec int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%3d", time.Now(), rand.Int())
	ch := make(chan string, 10)

	s.mu.Lock()
	s.filter[id] = func(key string) {
		ch <- key
	}
	s.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSec) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}

	//return nil // Unreachable code
}

func (s *KVStoreService) Get(key string, value *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if v, ok := s.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (s *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	key, value := kv[0], kv[1]

	// key 变更时调用 watch 的过滤器函数，以便进行通知
	if old := s.m[key]; old != value {
		for _, fn := range s.filter {
			fn(key)
		}
	}

	s.m[key] = value

	fmt.Println("set", kv)
	return nil
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}
