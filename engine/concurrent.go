package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seed ...Request) {

	in := make(chan Request)
	out := make(chan RequestResult)
	e.Scheduler.ConfigMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(in, out)
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

func (e *ConcurrentEngine) createWorker(in chan Request, out chan RequestResult) {
	go func() {
		for {
			request := <-in
			result := Work(request)
			out <- result
		}
	}()
}
