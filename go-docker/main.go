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
		fn := fp.Filter(isGreaterThanOne)
		res := fn([]int{1, 2, 3, 4, 5})
		fmt.Println(res)
		w.Write([]byte("Hello World! from Go using Docker!"))
	})

	http.HandleFunc("/:id", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		message := fmt.Sprintf("Route entered with id: %s", id)
		w.Write([]byte(message))
	})

	http.ListenAndServe(":3000", nil)
}
