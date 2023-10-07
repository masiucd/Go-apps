package main

import (
	"log"
	"os"
)

func readFile(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func getAbsolutePath() (string, error) {
	abs, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return abs, nil
}

func main() {
	abs, err := getAbsolutePath()
	if err != nil {
		panic(err)
	} else {
		log.Println(abs)
	}
	file, err := readFile(abs + "/" + "file.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(file)
	}

}
