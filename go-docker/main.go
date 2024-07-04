package main

import (
	"fmt"
	"net/http"

	"github.com/repeale/fp-go"
)

func isGreaterThanOne(n int) bool {
	return n > 1
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World! from Go using Docker!"))
	})
	fn := fp.Filter(isGreaterThanOne)
	res := fn([]int{1, 2, 3, 4, 5})
	fmt.Println(res)
	fmt.Println("Hello World! from Go using Docker!")
	http.ListenAndServe(":3000", nil)
}
