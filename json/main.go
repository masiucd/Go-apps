package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Human interface {
	speak()
}

func (p Person) speak() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	p := Person{
		Name: "John",
		Age:  30,
	}

	// Convert struct to JSON
	j, err := convertToJson(p)
	if err != nil {
		panic(err)
	}
	absolutePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = write(absolutePath+"/person.json", j)
	if err != nil {
		panic(err)
	}

	frank, err := parseJson(j)
	if err != nil {
		panic(err)

	}
	frank.speak()

}

func convertToJson(item Human) ([]byte, error) {
	// Convert struct to JSON
	j, err := json.Marshal(item)
	return j, err
}

func write(path string, item []byte) error {
	err := os.WriteFile(path, item, 0644)
	return err
}

func parseJson(item []byte) (Human, error) {
	p := Person{}
	err := json.Unmarshal([]byte(`{"Name":"Frank","Age":29}`), &p)
	return p, err
}
