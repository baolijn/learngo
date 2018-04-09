package engine

import "fmt"

type Scheduler interface {
	Submit(request Request)
}

type ConcurrentEngine struct {
	scheduler   Scheduler
	WorkerCount int
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		e.scheduler.Submit(r)
	}

	in := make(chan Request)
	out := make(chan ParseResult)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item: %v", item)
		}

		for _, request := range result.Requests {
			e.scheduler.Submit(request)
		}

	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
