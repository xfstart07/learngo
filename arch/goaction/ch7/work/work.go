package work

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGorutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGorutines)
	for i := 0; i < maxGorutines; i++ {
		go func() {
			// p.work 将阻塞，等待接收到 channel，当关闭 p.work 时，for 循环就结束了
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
