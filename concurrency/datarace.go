package concurrency

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func Race() {
	wg.Add(2)
	go Increment()
	go Increment()
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func Increment() {
	for i := 0; i < 1000; i++ {
		counter++ // Race condition: multiple goroutines modifying counter
	}
	wg.Done()
}
