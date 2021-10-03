package engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkerChannel(chan Request)
}
func (engine *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	engine.Scheduler.ConfigureMasterWorkerChannel(in)
	for i := 0; i < engine.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		engine.Scheduler.Submit(r)
	}
	count := 0
	for {
		resp := <- out
		for _, property := range resp.Properties {
			count++
			fmt.Println("get property: ", property, count)
		}
		for _, req := range resp.Requests {
			engine.Scheduler.Submit(req)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <- in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()

}