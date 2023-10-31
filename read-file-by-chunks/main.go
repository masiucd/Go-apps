package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Chunk size
	const maxSize = 12
	// Create buffer

	var xs [][]byte

	for {
		b := make([]byte, maxSize)
		// read content to buffer
		total, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		xs = append(xs, b)
		fmt.Println(" Reading  ", string(b[:total]))
	}
	printXs(xs)
}

func printXs(xs [][]byte) {
	for _, x := range xs {
		fmt.Println(string(x))
	}
}
