// Author: xufei
// Date: 2019/4/30

package data

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}
