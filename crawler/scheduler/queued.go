// Author: xufei
// Date: 2018/12/10

package scheduler

import (
	"learngo/crawler/engine"
)

type QueuedSchedule struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedSchedule) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedSchedule) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedSchedule) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

func (*QueuedSchedule) ConfigureWorkerChan(chan engine.Request) {
	panic("implement me")
}
