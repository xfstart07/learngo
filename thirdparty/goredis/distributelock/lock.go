// Author: xufei
// Date: 2020/11/4

package distributelock

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	"github.com/go-redis/redis"
)

// 分布式锁的实现之 redis 篇
// https://xiaomi-info.github.io/2019/12/17/redis-distributed-lock/
// https://mp.weixin.qq.com/s/qJK61ew0kCExvXrqb7-RSg

type DistributeLock interface {
	Lock() (bool, error)
	Unlock() error
	Refresh(ttl time.Duration) error
}

type redisLock struct {
	cli        *redis.Client
	key, value string
	expr       time.Duration
}

func (r *redisLock) Lock() (bool, error) {
	r.value = r.randomValue()

	ok, err := r.cli.SetNX(r.key, r.value, r.expr).Result()
	return ok, err
}

func (r *redisLock) Unlock() error {
	if r.cli.Get(r.key).Val() == r.value {
		return r.cli.Del(r.key).Err()
	}
	return fmt.Errorf("value is not consistent")
}

func (r *redisLock) Refresh(ttl time.Duration) error {
	if r.cli.Get(r.key).Val() == r.value {
		return r.cli.PExpire(r.key, ttl).Err()
	}
	return fmt.Errorf("value is not consistent")
}

func (r *redisLock) randomValue() string {
	tmp := make([]byte, 16)
	_, _ = io.ReadFull(rand.Reader, tmp)

	return base64.RawURLEncoding.EncodeToString(tmp)
}

func NewLock(cli *redis.Client, key string, expr time.Duration) *redisLock {
	return &redisLock{
		cli:  cli,
		key:  key,
		expr: expr,
	}
}
