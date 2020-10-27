// Author: xufei
// Date: 2020/7/2

package department

import "context"

func DecodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func EndcodeGrpcResponse(_ context.Context, resp interface{}) (interface{}, error) {
	return resp, nil
}
