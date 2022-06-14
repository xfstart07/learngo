// Author: xufei
// Date: 2019/4/29

package hummer2

type AHummer struct {
	Hummer HInterface
}

// 父类实现模板方法，子类实现具体方法
func (h *AHummer) Run() {
	if h.Hummer == nil {
		return
	}

	h.Hummer.Start()
	h.Hummer.EngineBoom()
	h.Hummer.Alarm()
	h.Hummer.Stop()
}
