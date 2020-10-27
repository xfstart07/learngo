// Author: xufei
// Date: 2020/7/2

package department

import (
	"context"
	"learngo/microservices/gokit/one-part/protobuf-spec/common"

	"github.com/go-kit/kit/endpoint"
)

func makeListEndpoint(srv *DepartmentManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		list, err := srv.List()
		if err != nil {
			return nil, err
		}

		resp := &common.DepartmentList{}
		for _, d := range list {
			resp.List = append(resp.List, d)
		}

		return resp, nil
	}
}

func makeCreateEndpoint(srv *DepartmentManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*common.Department)
		return &common.Empty{}, srv.Create(req)
	}
}

func makePersonChangeEndpont(srv *DepartmentManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*common.PersonChangeRequest)
		return &common.Empty{}, srv.PersonnelChange(req.User, req.Company)
	}
}
