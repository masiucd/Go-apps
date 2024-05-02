package read

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

const FOUR_KB = 4096

// ReadFileV2 reads the file and returns the data as a string
func ReadFileV2(f string) string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	// Create a new reader
	r := bufio.NewReader(file)
	// The buffer size in Go determines how much data you can read from a file at once.
	// If you set the buffer size to 4 bytes. So we will read 4 bytes of data from the file at a time.
	buf := make([]byte, FOUR_KB) //  4096 bytes = 4KB
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

// ReadFileV1 reads the file and returns the data as a byte slice
func ReadFileV1(f string) []byte {
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
