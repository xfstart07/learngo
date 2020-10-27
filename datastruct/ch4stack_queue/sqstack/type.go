// Author: xufei
// Date: 2019-07-26

package sqstack

const (
	MAXSIZE = 100
)

type SqStack struct {
	data [MAXSIZE]int
	top  int
}
