// Author: xufei
// Date: 2019/4/29

package hummer3

type HInterface interface {
	Start()
	Stop()
	Alarm()
	EngineBoom()
	// 模板方法
	Run()
	IsAlarm() bool
}
