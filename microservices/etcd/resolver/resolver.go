// Author: xufei
// Date: 2020/11/19

package resolver

import (
	"context"
	"log"

	"learngo/microservices/etcd/registry"
	"learngo/pkg/xgo"

	"google.golang.org/grpc/resolver"
)

func Register(schema string, reg registry.Registry) {
	resolver.Register(&Builder{
		schema: schema,
		reg:    reg,
	})
}

type Builder struct {
	schema string
	reg    registry.Registry
}

func (b *Builder) Scheme() string {
	return b.schema
}

func (b *Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	endpoints, err := b.reg.WatchServices(context.Background(), target.Endpoint, "grpc")
	if err != nil {
		return nil, err
	}

	var stop = make(chan struct{})
	xgo.Go(func() {
		for {
			select {
			case endpoint := <-endpoints:
				var state = resolver.State{
					Addresses: make([]resolver.Address, 0),
				}

				for _, node := range endpoint.Nodes {
					addr := resolver.Address{
						Addr:       node.Address,
						ServerName: target.Endpoint,
					}
					state.Addresses = append(state.Addresses, addr)
				}

				log.Printf("最新的地址: %+v", state)
				cc.UpdateState(state)
			case <-stop:
				return
			}
		}
	})

	return &baseResolver{stop: stop}, nil
}

type baseResolver struct {
	stop chan struct{}
}

func (b *baseResolver) ResolveNow(opts resolver.ResolveNowOptions) {}

func (b *baseResolver) Close() {
	b.stop <- struct{}{}
}
