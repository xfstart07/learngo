// Author: xufei
// Date: 2020/11/19

package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"learngo/pkg/xgo"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

var (
	ReadTimeout = 10 * time.Second
)

type etcdv3Registry struct {
	client     *clientv3.Client
	prefix     string        // 服务前缀
	serviceTTL time.Duration // 服务过期时间
	mu         sync.Mutex
	endpoints  Endpoints
}

// RegisterService 服务注册
func (reg *etcdv3Registry) RegisterService(ctx context.Context, info *ServiceInfo) error {
	var readCtx context.Context
	var readCancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		readCtx, readCancel = context.WithTimeout(ctx, ReadTimeout)
		defer readCancel()
	}

	key := reg.registerKey(info)
	value := reg.registerValue(info)

	opOptions := make([]clientv3.OpOption, 0)
	if reg.serviceTTL.Seconds() > 0 {
		// 设置租约时间
		ctx := context.Background()
		// TODO: 申请了租约还需要考虑续约的逻辑，etcd sessions
		grant, err := reg.client.Grant(ctx, int64(reg.serviceTTL.Seconds()))
		if err != nil {
			return err
		}
		opOptions = append(opOptions, clientv3.WithLease(grant.ID))
	}
	_, err := reg.client.Put(readCtx, key, value, opOptions...)
	if err != nil {
		return err
	}
	log.Printf("register service: %v, %v", key, value)

	return nil
}

// UnregisterService 注销服务
func (reg *etcdv3Registry) UnregisterService(ctx context.Context, info *ServiceInfo) error {
	key := reg.registerKey(info)
	return reg.unregister(ctx, key)
}

func (reg *etcdv3Registry) unregister(ctx context.Context, key string) error {
	if _, ok := ctx.Deadline(); ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, ReadTimeout)
		defer cancel()
	}

	_, err := reg.client.Delete(ctx, key)
	return err
}

// ListServices 获取服务列表
func (reg *etcdv3Registry) ListServices(ctx context.Context, name string, schema string) (services []*ServiceInfo, err error) {
	key := fmt.Sprintf("/%s/%s/%s://", reg.prefix, name, schema)
	get, err := reg.client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	for _, ks := range get.Kvs {
		var info ServiceInfo
		if err := json.Unmarshal(ks.Value, &info); err != nil {
			log.Printf("信息错误: %v", err)
			continue
		}

		services = append(services, &info)
	}

	return
}

// WatchServices 观察服务变化
func (reg *etcdv3Registry) WatchServices(ctx context.Context, name string, schema string) (chan Endpoints, error) {
	prefix := fmt.Sprintf("/%s/%s/%s://", reg.prefix, name, schema)
	get, err := reg.client.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	reg.updateAddr(get.Kvs...)

	var addresses = make(chan Endpoints, 10)
	addresses <- reg.endpoints

	watch := reg.client.Watch(ctx, prefix, clientv3.WithPrefix())
	xgo.Go(func() {
		for event := range watch {
			for _, ev := range event.Events {
				switch ev.Type {
				case mvccpb.PUT:
					reg.updateAddr(ev.Kv)
				case mvccpb.DELETE:
					reg.deleteAddr(ev.Kv)
				}

				select {
				case addresses <- reg.endpoints:
				default:
					log.Println("invalid")
				}
			}
		}
	})

	return addresses, nil
}

func (reg *etcdv3Registry) updateAddr(kvs ...*mvccpb.KeyValue) {
	reg.mu.Lock()
	defer reg.mu.Unlock()
	for _, ks := range kvs {
		var info ServiceInfo
		if err := json.Unmarshal(ks.Value, &info); err != nil {
			log.Printf("信息错误: %v", err)
			continue
		}
		parse, err := url.Parse(info.Address)
		if err != nil {
			log.Printf("地址错误: %v", err)
			continue
		}
		reg.endpoints.Nodes[parse.String()] = info
	}
}

func (reg *etcdv3Registry) deleteAddr(kvs ...*mvccpb.KeyValue) {
	reg.mu.Lock()
	defer reg.mu.Unlock()
	for _, ks := range kvs {
		var info ServiceInfo
		if err := json.Unmarshal(ks.Value, &info); err != nil {
			log.Printf("信息错误: %v", err)
			continue
		}
		parse, err := url.Parse(info.Address)
		if err != nil {
			log.Printf("地址错误: %v", err)
			continue
		}
		delete(reg.endpoints.Nodes, parse.String())
	}
}

func (reg *etcdv3Registry) Close() error {
	var sy sync.WaitGroup
	for _, info := range reg.endpoints.Nodes {
		sy.Add(1)

		go func() {
			defer sy.Done()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			err := reg.unregister(ctx, reg.registerKey(&info))
			if err != nil {
				log.Printf("unregister service err: %v", err)
			} else {
				log.Printf("unregister service ok: %v", info)
			}

			cancel()
		}()
	}
	sy.Wait()

	return nil
}

func (reg *etcdv3Registry) registerKey(info *ServiceInfo) string {
	return GetServiceKey(reg.prefix, info)
}

func (reg *etcdv3Registry) registerValue(info *ServiceInfo) string {
	return GetServiceValue(info)
}

func newEtcdRegistry(prefix string, client *clientv3.Client) *etcdv3Registry {
	return &etcdv3Registry{
		client:    client,
		prefix:    prefix,
		mu:        sync.Mutex{},
		endpoints: Endpoints{Nodes: make(map[string]ServiceInfo)},
	}
}
