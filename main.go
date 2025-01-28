package main

import "learn-go/input"

func main() {
	// concurrency.Channels()
	// strings.StringMethods()

	// concurrency.Concurrency()
	// variables.Variables()

	// uncomment below line for Generator pattern
	// concurrency.Generator(1, 10)

	// uncomment below line for Consumer pattern
	// concurrency.Consumer()

	// uncomment below line for Consumer pattern
	// concurrency.Pipeline()

	// uncomment from 14-19 for fan-in pattern
	// ch1 := concurrency.Generator(1, 10)
	// ch2 := concurrency.Generator(11, 20)
	// res := concurrency.FanIn(ch1, ch2)
	// for value := range res {
	// 	fmt.Println("Value:", value)
	// }

	// uncomment from 17-19 for fan-out pattern
	// numbers := concurrency.Generator(1, 20)
	// workers := make(chan struct{}, 4)
	// concurrency.FanOut(numbers, workers)

	// uncomment from 37-43 for fan-out pattern
	// const workers = 4 //change this to set the number of workers
	// jobQueue := concurrency.InitWorkers(4)
	// for i := range 100 {
	// 	jobQueue <- i
	// }
	// concurrency.ScheduleJobs(workers, jobQueue)
	// uncomment 44 to see magic
	// select {}

	// uncomment below line for Queue based workerpool pattern
	// concurrency.Workerpool(10)

	// Fan-Out,Fan-In Pattern (combined)
	/*
		Distributes tasks to multiple worker goroutines.
		Combines the results from multiple worker goroutines into a single output channel.
	*/
	// ch1 := concurrency.Generator(1, 10)
	// ch2 := concurrency.Generator(11, 20)
	// ch3 := concurrency.Generator(21, 30)
	// concurrency.FanOutFanIn(ch1, ch2, ch3)

	// Profiling
	// uncomment the below line to start a server on port 9600, on which you can profile
	// profiling.Serve()

	input.Read()
}
