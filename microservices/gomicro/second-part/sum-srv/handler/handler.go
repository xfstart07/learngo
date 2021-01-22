// Author: xufei
// Date: 2020/3/5

package handler

import (
	"context"
	"learngo/gomicro/second-part/proto/sum"
	"learngo/gomicro/second-part/sum-srv/service"
)

type handler struct {
}

func (h handler) GetSum(ctx context.Context, req *sum.SumRequest, resp *sum.SumRepose) error {

	n := req.Input
	ret := service.GetSum(n)

	resp.Output = ret
	return nil
}

func NewHandler() sum.SumHandler {
	return &handler{}
}
