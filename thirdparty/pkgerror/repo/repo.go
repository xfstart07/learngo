// Author: xufei
// Date: 2020/7/23

package repo

import pkgerr "github.com/pkg/errors"

func Get() error {
	err := find()

	return pkgerr.WithStack(err)
}

func Update() error {
	err := find()

	return pkgerr.Wrap(err, "update fail")
}

func find() error {
	return pkgerr.New("not found")
}
