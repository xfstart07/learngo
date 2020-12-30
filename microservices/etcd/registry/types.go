// Author: xufei
// Date: 2020/11/19

package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type ServiceInfo struct {
	Name    string `json:"name"`    // 服务名称
	Schema  string `json:"schema"`  // 服务协议，http, grpc
	Address string `json:"address"` // 地址
}

type Endpoints struct {
	// 服务节点列表
	Nodes map[string]ServiceInfo
}

type Registry interface {
	RegisterService(context.Context, *ServiceInfo) error
	UnregisterService(context.Context, *ServiceInfo) error
	ListServices(context.Context, string, string) ([]*ServiceInfo, error)
	WatchServices(context.Context, string, string) (chan Endpoints, error)
	io.Closer
}

//GetServiceKey ..
func GetServiceKey(prefix string, s *ServiceInfo) string {
	// 格式: microservices/app_name/grpc/ip:port
	return fmt.Sprintf("/%s/%s/%s://%s", prefix, s.Name, s.Schema, s.Address)
}

//GetServiceValue ..
func GetServiceValue(s *ServiceInfo) string {
	val, _ := json.Marshal(s)
	return string(val)
}
