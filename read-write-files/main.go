package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 10 is the ASCII code for newline
func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	simplePrint(data)

}

const newLine = 10

// simplePrint prints the data to the console
func simplePrint(data []byte) {
	var output string
	for _, v := range data {
		if v == newLine {
			// Print a newline character
			output += "\n"
		} else {
			// Print the character
			output += string(v)
		}
	}
	fmt.Println(output)
}
