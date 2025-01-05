package concurrency

import "fmt"

func Channels() {
	/*
		Blocking on Send
	*/
	ch1 := make(chan int) // Unbuffered channel
	go func() {
		ch1 <- 42 // Blocks until the value is received
		fmt.Println("Sent")
	}()
	fmt.Println(<-ch1)

	ch2 := make(chan int, 1) // Buffered channel
	go func() {
		ch2 <- 42 // will not be blocked
		ch2 <- 43 // Blocks until the value is received
		fmt.Println("Sent")
	}()
	fmt.Println(<-ch2)

	/*
		Blocking on Receive
	*/
	ch3 := make(chan int, 1) // Buffered channel
	go func() {
		fmt.Println("Sent")
		ch3 <- 42
	}()
	fmt.Println("Receiving")
	fmt.Println(<-ch3) // Blocks until a value is sent

	/*
		Non-Blocking Behavior
	*/
	ch4 := make(chan int)
	select {
	case v := <-ch4:
		fmt.Println("Received:", v) // Won't execute (no value to receive)
	default:
		fmt.Println("No data") // Executes immediately will not be blocked
	}

	/*
		Reading from a closed channel
	*/
	// If channels buffer is not empty
	ch5 := make(chan int, 2)
	ch5 <- 10
	ch5 <- 20
	close(ch5)

	fmt.Println(<-ch5) // Prints: 10
	fmt.Println(<-ch5) // Prints: 20
	fmt.Println(<-ch5) // Prints: 0 (channel is closed and empty),this will not be blocked

	chStr1 := make(chan string, 2)
	chStr1 <- "abc"
	chStr1 <- "def"
	close(chStr1)

	fmt.Println(<-chStr1)
	fmt.Println(<-chStr1)
	fmt.Println(<-chStr1)

	// Channel Buffer is empty
	ch6 := make(chan int)
	close(ch6)
	v, ok := <-ch6
	fmt.Println(v, ok)

	chStr2 := make(chan string)
	close(chStr2)
	s, ok := <-chStr2
	fmt.Println(s, ok)

	/*
		Iterating Over a Channel
	*/
	// Case 1: Channel is closed and buffer is not-empty.

	ch7 := make(chan int, 2)
	ch7 <- 10
	ch7 <- 20
	close(ch7)
	for v := range ch7 {
		fmt.Println(v)
	}

	// Same Example with UnBuffered channel
	ch8 := make(chan int, 2)
	go func() {
		ch8 <- 10
		ch8 <- 20
		close(ch8)
	}()
	for v := range ch8 {
		fmt.Println(v)
	}

	// Case 2: channel is closed and buffer is empty.

	ch9 := make(chan int, 2)
	close(ch9)
	for v := range ch9 {
		fmt.Println(v) //No output
	}
	fmt.Println("Execution will reach here") //Execution will reach this line

	// Case 3:  Channel is not closed(open) and the buffer is not empty

	// UnBuffered Channel
	// Use a goroutine with an unbuffered channel,since the receiver is not yet ready, send would block forever.
	ch10 := make(chan int)
	go func() {
		ch10 <- 10
	}()
	for v := range ch10 {
		fmt.Println(v) //Prints 10 & then Deadlock
	}
	//  Buffered Channel
	ch11 := make(chan int, 1)
	ch11 <- 100
	for v := range ch11 {
		fmt.Println(v) //Prints 10 & then Deadlock
	}
	fmt.Println("Execution will not reach here") //Execution will not reach this line
	// Reason for Deadlock in both cases: The channel is not closed, and the for range loop waits
	// indefinitely for more values.

	// Case 4:  Channel is not closed and the buffer is empty
	ch12 := make(chan int)
	for v := range ch12 {
		fmt.Println(v) //No output
	}
	fmt.Println("Execution will not reach here") //Execution will not reach this line

	// Closing an Already Closed Channel
	ch13 := make(chan int)
	go func() {
		ch13 <- 10
		close(ch13)
	}()
	<-ch13
	close(ch13)

	// What will be the output of the following
	ch14 := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch14 <- i
	}
	for range ch14 {
		fmt.Println(<-ch14)
	}
	// Output : 2 4 6 8 10
	// The line for range ch14 consumes values from the channel, and within the loop,
	// you also explicitly read from the channel using <-ch14.
	// This results in skipping every other value because:
	// for range ch14:
	// Automatically reads a value from the channel and discards it.
	// <-ch14:
	// Reads the next value from the channel within the loop.

	//Proper way
	ch15 := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch14 <- i
	}
	for v := range ch15 {
		fmt.Println(v)
	}
}
