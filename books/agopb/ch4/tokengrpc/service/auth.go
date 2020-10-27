// Author: xufei
// Date: 2019-08-29

package service

import (
	"context"
	"fmt"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

func (*Authentication) RequireTransportSecurity() bool {
	return false
}

func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var appid, appkey string

	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	if appid != a.User || appkey != a.Password {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}
