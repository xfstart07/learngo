// Author: xufei
// Date: 2020/3/5

package handler

import (
	"context"
	"learngo/gomicro/second-part/prime-srv/service"
	"learngo/gomicro/second-part/proto/prime"
)

type handler struct {
}

func (h handler) GetPrime(ctx context.Context, req *prime.PrimeRequest, resp *prime.PrimeRepose) error {
	n := req.Input

	primes := service.GetPrime(n)

	resp.Output = primes

	return nil
}

func NewHandler() prime.PrimeHandler {
	return handler{}
}
