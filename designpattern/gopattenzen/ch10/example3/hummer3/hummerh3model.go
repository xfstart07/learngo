// Author: xufei
// Date: 2019/4/29

package hummer3

import "fmt"

type H3Model struct {
}

func (*H3Model) Start() {
	fmt.Println("悍马 H3 发动...")
}

func (*H3Model) Stop() {
	fmt.Println("悍马 H3 停车...")
}

func (*H3Model) Alarm() {
	fmt.Println("悍马 H3 鸣笛...")
}

func (*H3Model) EngineBoom() {
	fmt.Println("悍马 H3 引擎声...")
}
