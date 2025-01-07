package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Generator functions are used to generate a sequence of values.
func Generator(start, end int) <-chan int {
	ch := make(chan int)
	if start >= end {
		return nil
	}
	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			ch <- i
		}
	}()
	return ch
}

// Consumer consumes the value produced by Generator
func Consumer() {
	ch := Generator(1, 20)
	for range ch {
		fmt.Println(<-ch)
	}

	// Simple example would be
	c := time.NewTicker(1 * time.Second)
	for v := range c.C {
		fmt.Println("ticker ticked", v)
	}
}

// Pipeline
func Pipeline() {
	natural := Generator(1, 20)
	square := make(chan int)
	cube := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer close(square)
		defer wg.Done()
		for {
			for v := range natural {
				square <- v * v
			}
		}
	}()
	go func() {
		defer wg.Done()
		defer close(cube)
		for v := range square {
			cube <- v * v * v
		}
	}()
	go func() {
		defer wg.Done()
		for i := range cube {
			fmt.Println("value : ", i)
		}
	}()
	wg.Wait()
	fmt.Println("All go routines have been executed")
}

// Fanin
// The Fan-in pattern in Go is a concurrency pattern where multiple channels are combined into a single channel
// You have multiple goroutines sending data to a single channel.
func FanIn(channels ...<-chan int) <-chan int { //recommended to make it a variadic function if same data type
	mergedCh := make(chan int)
	wg := &sync.WaitGroup{}
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				mergedCh <- v
				time.Sleep(1 * time.Second)
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(mergedCh)
	}()
	return mergedCh
}

func FanOut(numbers <-chan int, workers chan struct{}) { //recommended to make it a variadic function if same data type
	for i := range numbers {
		workers <- struct{}{}
		go func(i int, ch chan struct{}) {
			defer func() {
				<-ch
			}()
			fmt.Println("Woker ID:", i%len(workers)+1, "\tResult : ", i) // i % len(workers) + 1, ensures IDs range from 1 to the number of workers.
		}(i, workers)
	}
}

// Queue based Workerpool1 pattern
func Workerpool1(poolSize int) {
	//change this to set the number of workers
	ch := make(chan struct{}, poolSize)
	for range poolSize {
		ch <- struct{}{} //since a buffered channel is used this would block if the buffer is full/until a worker is available
		go operation(ch)
	}
}

func operation(ch chan struct{}) {
	defer func() {
		<-ch
	}()
	fmt.Println("Time:", time.Now().UTC())
	time.Sleep(2 * time.Second)
}

// Workerpools
func InitWorkers(workers int) chan<- int {
	jobQueue := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i <= workers; i++ {
		wg.Add(1)
		go func(id int, jobQueue <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobQueue {
				fmt.Printf("worker %d executed task %d\n", i, j)
			}
		}(i, jobQueue, &wg)
	}
	return jobQueue
}

func ScheduleJobs(workers int, jobQueue chan<- int) {
	go func() {
		for {
			job := rand.Intn(100) + 1 // Generate a random job ID (1-100)
			jobQueue <- job           // Send the job to the job queue
			// randomDelay := time.Duration(rand.Intn(1000)+500) * time.Millisecond // Random delay between 500ms and 1.5s
			time.Sleep(2 * time.Second)
		}
	}()
}

func FanOutFanIn(channels ...<-chan int) {
	wg := &sync.WaitGroup{}
	res := make(chan int)
	//Fan-Out
	for _, channel := range channels {
		wg.Add(1)
		go sum(channel, wg, res)
	}
	go func() {
		wg.Wait()
		close(res)
	}()
	result := 0
	for v := range res {
		result += v
	}
	fmt.Println("Sum of natural numbers from 1-30 : ", result)
}

func sum(in <-chan int, wg *sync.WaitGroup, res chan int) {
	defer wg.Done()
	total := 0
	for num := range in {
		total += num
	}
	// Fan-In
	res <- total
}
