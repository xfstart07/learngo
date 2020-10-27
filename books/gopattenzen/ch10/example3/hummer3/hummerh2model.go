// Author: xufei
// Date: 2019/4/29

package hummer3

import "fmt"

type H2Model struct {
}

func (*H2Model) Start() {
	fmt.Println("悍马 H2 发动...")
}

func (*H2Model) Stop() {
	fmt.Println("悍马 H2 停车...")
}

func (*H2Model) Alarm() {
	fmt.Println("悍马 H2 鸣笛...")
}

func (*H2Model) EngineBoom() {
	fmt.Println("悍马 H2 引擎声...")
}

func (h *H2Model) IsAlarm() bool {
	return false
}
