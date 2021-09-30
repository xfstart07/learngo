// Author: xufei
// Date: 2020/11/19

package main

import (
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func BuildClient(endpoints []string) *clientv3.Client {
	cfg := clientv3.Config{
		Endpoints:   []string{"http://weixinote.dev:2379"},
		DialTimeout: 3 * time.Second,
	}

	c, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
