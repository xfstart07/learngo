// Author: xufei
// Date: 2019/3/21

package buildin

import (
	"fmt"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	s := make([]int, 2, 3)
	fmt.Printf("%+v\n", s)
	s1 := s[1:] // cap = end - index
	fmt.Printf("%+v, %d, %d\n", s1, len(s1), cap(s1))

	s2 := s[2:2]
	fmt.Printf("%+v, %d, %d\n", s1, len(s2), cap(s2))
}
