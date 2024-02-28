package main

import (
	"fmt"
	"time"
)

func count(message string) {
	for i := range 10 {
		fmt.Println(message, i)
	}
}

func main() {
	// start time to measure the time taken
	start := time.Now()

	go count("concurrent")
	count("blocking")
	fmt.Println("Done")
	time.Sleep(time.Second * 1)

	// end time to measure the time taken
	end := time.Now()
	fmt.Println("Time taken:", end.Sub(start))
}
