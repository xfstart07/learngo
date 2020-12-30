// Author: xufei
// Date: 2018/12/17

package v1

import "fmt"

type DivisionDemo struct{}

type Args struct {
	A, B int
}

func (DivisionDemo) Div(args *Args, result *float64) error {
	if args.B == 0 {
		return fmt.Errorf("divisor by zero")
	}

	*result = float64(args.A) / float64(args.B)

	return nil
}
