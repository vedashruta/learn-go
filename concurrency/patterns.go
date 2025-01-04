package concurrency

import "fmt"

func generator() {

}

// Pipeline
func Pipeline() {
	natural := make(chan int)
	square := make(chan int)
	go func() {
		for i := range 100 {
			natural <- i
		}
		close(natural)
	}()
	go func() {
		for {
			i, ok := <-natural
			if !ok {
				break
			}
			square <- i * i
		}
		close(square)
	}()
	for {
		i, ok := <-square
		if !ok {
			break
		}
		fmt.Println(i)
	}
}
