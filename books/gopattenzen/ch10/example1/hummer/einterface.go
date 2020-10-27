// Author: xufei
// Date: 2019/4/29

package hummer

type HInterface interface {
	Start()
	Stop()
	Alarm()
	EngineBoom()
	Run()
}
