package scheduler

import "learngo/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerchan chan chan engine.Request
}

func (s QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

