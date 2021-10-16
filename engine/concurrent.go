package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	PropertyChan chan Property
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
	outputChannel := make(chan ParseResult)
	engine.Scheduler.Run()
	for i := 0; i < engine.WorkerCount; i++ {
		createWorker(engine.Scheduler.WorkerChan(), outputChannel, engine.Scheduler)
	}

	for _, r := range seeds {
		engine.Scheduler.Submit(r)
	}

	for {
		resp := <- outputChannel
		for _, property := range resp.Properties {
			go func() {
				if resp.Store {
					engine.PropertyChan <- property
				}
			}()
		}
		for _, req := range resp.Requests {
			engine.Scheduler.Submit(req)
		}
	}
}

func createWorker(inputChannel chan Request, outputChannel chan ParseResult,  notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(inputChannel)
			request := <- inputChannel
			result, err := Execute(request)
			if err != nil {
				continue
			}
			outputChannel <- result
		}
	}()
}