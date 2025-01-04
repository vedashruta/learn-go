package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Monitor GoRoutines
var (
	deposits = make(chan int)
	balances = make(chan int)
	once     = &sync.Once{}
	counter  int
	wg       sync.WaitGroup
)

func init() {
	go teller() //Monitor go routine
	once.Do(initOnlyOnce)
	time.Sleep(4 * time.Second)
	once.Do(initOnlyOnce)
}

func teller() {
	var bal int
	for {
		select {
		case amount := <-deposits:
			fmt.Println("inside case 1")
			bal += amount
			notify(bal)
		case balances <- bal:
			fmt.Println("inside case 2")
			notify(bal)
		}
	}
}

func notify(bal int) {
	fmt.Println("balance:", bal)
}

func dep(amount int) {
	deposits <- amount
}

func bal() int {
	return <-balances
}

func Concurrency() {

	// //Buffered Channel
	// ch1 := make(chan struct{}, 1)

	// //UnBuffered Channel
	// ch2 := make(chan struct{})

	// var ch3 chan struct{}

	// blockingChannels()
	// blockingChannels2()
	// dep(100)
	// bal()
	// dep(200)
	// dep(500)
	// dep(400)
	// dep(1)

	//Race condition with -race flag enabled
	wg.Add(2)
	go Increment()
	go Increment()
	wg.Wait()
	fmt.Println("Counter:", counter)

}

// We can use buffered channel as a semaphore
func blockingChannels() {
	ch1 := make(chan struct{}, 1)
	go func() {
		defer func() {
			ch1 <- struct{}{}
		}()
		fmt.Println("Executing go routine 1")
		for i := 0; i < 1000; i++ {
			if i%2 == 0 {
				fmt.Println(i) // Print odd numbers
			}
		}
	}()
	<-ch1
	fmt.Println("Goroutine 1 executed")
	go func() {
		defer func() {
			ch1 <- struct{}{}
		}()
		fmt.Println("Executing go routine 1")
		for i := 0; i < 1000; i++ {
			if i%2 != 0 {
				fmt.Println(i) // Print odd numbers
			}
		}
	}()
	<-ch1
	fmt.Println("Goroutine 2 executed")
	fmt.Println("All Goroutines are executed")
}

func blockingChannels2() {
	ch := make(chan struct{}, 1)
	now := time.Now()
	var num int
	fmt.Println("started at : ", now)
	wg := &sync.WaitGroup{}
	for i := range 20 {
		wg.Add(1)
		go func() {
			ch <- struct{}{}
			defer func() {
				<-ch
			}()
			defer wg.Done()
			fmt.Println("go routine ", i, " running\nupdating variable to ", i)
			num = i
			fmt.Println("current value of num", num)
			time.Sleep(2 * time.Second)
		}()
	}
	wg.Wait()
	fmt.Println("completed at : ", time.Since(now))
}

func initOnlyOnce() {
	fmt.Println(time.Now())
	fmt.Println("Function is initialized")
}

func Increment() {
	for i := 0; i < 1000; i++ {
		counter++ // Race condition: multiple goroutines modifying counter
	}
	wg.Done()
}
