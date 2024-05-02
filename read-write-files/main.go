package main

import (
	"fmt"
	"go-apps/read-write-files.com/read"
)

func main() {

	var version string
	fmt.Println("Select what how you would like to read the file")
	fmt.Println("1. Read the file as a byte slice")
	fmt.Println("2. Read the file as a string")
	fmt.Scan(&version)

	switch version {
	case "1":
		fmt.Println("Reading the file as a string")
		data := read.ReadFileV1("file.txt")
		fmt.Println(string(data))
	case "2":
		fmt.Println("Reading the file as a byte slice")
		data := read.ReadFileV2("file.txt")
		fmt.Println(data)
	default:
		fmt.Println("Invalid option")
	}
}
