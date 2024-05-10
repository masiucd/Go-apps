package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-apps/auth.com/db"
	"go-apps/auth.com/routes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// dir := http.Dir("./static")
	// fs := http.FileServer(dir)

	db.ConnectDB()
	mux := routes.Init()
	loggedMux := NewLoggingMiddleware(mux)

	fmt.Println("Server is running on port 4000")
	http.ListenAndServe(":4000", loggedMux)

}

type LoggingMiddleware struct {
	handler http.Handler
}

func (l *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Log request details before handling the request
	log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)

	// Pass the request to the actual handler
	l.handler.ServeHTTP(w, r)

	// Log response details after handling the request
	duration := time.Since(start)
	log.Printf("Request processed in %s", duration)
}

func NewLoggingMiddleware(handler http.Handler) http.Handler {
	return &LoggingMiddleware{handler}
}
