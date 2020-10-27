// Author: xufei
// Date: 2019-07-29

package circular_queue

const (
	MAXSIZE = 100
)

type SQueue struct {
	front, rear int
	elements    [MAXSIZE]int
}
