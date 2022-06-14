// Author: xufei
// Date: 2019/5/10

package main

import (
	"fmt"
	"time"
)

type (
	// 事件
	Event struct {
		Data int64
	}

	// 观察者抽象接口
	Observer interface {
		OnNotify(Event)
	}

	// 被观察者抽象接口
	Notifier interface {
		Register(Observer)
		Deregister(Observer)
		Notify(Event)
	}
)

type (
	// 观察者对象
	eventObserver struct {
		id int
	}

	// 被观察者对象
	eventNotifier struct {
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Notify(e Event) {
	for ob := range o.observers {
		ob.OnNotify(e)
	}
}

func main() {
	// 声明被观察者
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	// 注册有多少个观察者
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C

	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			// 通知观察者发生的事件
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}
