package engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (engine *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	engine.Scheduler.Run()
	for i := 0; i < engine.WorkerCount; i++ {
		createWorker(engine.Scheduler.WorkerChan(), out, engine.Scheduler)
	}

	for _, r := range seeds {
		engine.Scheduler.Submit(r)
	}

	for {
		resp := <- out
		for _, property := range resp.Properties {
			fmt.Println("get property: ", property)
		}
		for _, req := range resp.Requests {
			engine.Scheduler.Submit(req)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult,  notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			request := <- in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}