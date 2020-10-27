// Author: xufei
// Date: 2019-07-26

package static_link

const (
	MAXSIZE = 100
)

type StaticLinkList [MAXSIZE]component

type component struct {
	data, cur int
}
