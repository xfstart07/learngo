// Author: xufei
// Date: 2019/4/29

package hummer3

import "fmt"

type H1Model struct {
	AHummer
}

func (*H1Model) Start() {
	fmt.Println("悍马 H1 发动...")
}

func (*H1Model) Stop() {
	fmt.Println("悍马 H1 停车...")
}

func (*H1Model) Alarm() {
	fmt.Println("悍马 H1 鸣笛...")
}

func (*H1Model) EngineBoom() {
	fmt.Println("悍马 H1 引擎声...")
}

func (h *H1Model) SetAlarm(alarm bool) {
	h.AlarmFlag = true
}

func (h *H1Model) IsAlarm() bool {
	return h.AlarmFlag
}
