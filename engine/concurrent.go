package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seed ...Request) {

	out := make(chan RequestResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(out, e.Scheduler)
	}

	for _, v := range seed {
		e.Scheduler.Submit(v)
	}

	cnt := 0
	for {
		result := <-out
		for _, v := range result.Items {
			cnt++
			log.Printf("got Item #%d: :%v", cnt, v)
		}
		for _, v := range result.Requests {
			e.Scheduler.Submit(v)
		}
	}
}

func (e *ConcurrentEngine) createWorker(out chan RequestResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result := Work(request)
			out <- result
		}
	}()
}
