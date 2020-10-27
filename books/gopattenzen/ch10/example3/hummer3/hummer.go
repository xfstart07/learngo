// Author: xufei
// Date: 2019/4/29

package hummer3

type AHummer struct {
	Hummer    HInterface
	AlarmFlag bool
}

// 子类实现，父类调用模板方法
func (h *AHummer) Run() {
	if h.Hummer == nil {
		return
	}

	h.Hummer.Start()
	h.Hummer.EngineBoom()
	if h.Hummer.IsAlarm() {
		h.Hummer.Alarm()
	}

	h.Hummer.Stop()
}

func (h *AHummer) IsAlarm() bool {
	return true
}
