// Author: xufei
// Date: 2020/5/10

package greeter

import (
	"context"
	"learngo/gomicro/thiry-part/proto/learning"
)

type Handler struct{}

func (h *Handler) Hi(ctx context.Context, request *learning.Request, response *learning.Response) error {
	response.Msg = "Hello Greeter" + request.Name

	return nil
}
