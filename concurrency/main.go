package main

import (
	"sync"
)

// In this code, we use a sync.WaitGroup to wait for the fn function to finish sending data before we start reading from the channel.
//  This ensures that the channel is closed before we start reading from it, preventing a deadlock.

// fn is a function that takes a channel as an argument
// and sends 10 integers to the channel
// fn can only write to the channel
func fn(ch1 chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1) // close the channel
	wg.Done()  // decrement the WaitGroup counter and signal that the goroutine is done
}

func main() {
	ch1 := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1) // add 1 to the WaitGroup, so that the main goroutine waits for 1 goroutine to finish
	go fn(ch1, &wg)

	wg.Wait() // wait for the goroutine to finish

	// read from the channel
	for c := range ch1 {
		println(c)
	}

}
