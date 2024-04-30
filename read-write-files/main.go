package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 10 is the ASCII code for newline
func main() {

	// data := readFileV1("file.txt")
	// simplePrint(data)
	// printEven(data)
	res := readFileV2("file.txt")
	fmt.Println(res)

}

func readFileV2(f string) string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	// Create a new reader
	r := bufio.NewReader(file)
	// The buffer size in Go determines how much data you can read from a file at once.
	// If you set the buffer size to 4 bytes. So we will read 4 bytes of data from the file at a time.
	buf := make([]byte, 4) // 4 is the buffer size
	var output strings.Builder
	for {
		n, err := r.Read(buf)
		if n > 0 {
			output.Write(buf[:n])
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err.Error())
		}
	}
	return output.String()
}

// readFileV1 reads the file and returns the data as a byte slice
func readFileV1(f string) []byte {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	return data
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

func printEven(data []byte) {
	var builder strings.Builder
	lineCount := 0
	for _, v := range data {
		if v == '\n' {
			// If it's an even line, print it
			if lineCount%2 == 0 {
				fmt.Println(builder.String())
			}
			// Reset the builder and increase the line count
			builder.Reset()
			lineCount++
		} else {
			// Add the character to the current line
			builder.WriteByte(v)
		}
	}
	// Print the last line if it's an even line and the data doesn't end with a newline
	if lineCount%2 == 0 {
		fmt.Println(builder.String())
	}
}
