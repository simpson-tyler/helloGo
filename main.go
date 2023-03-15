package main

import (
	"log"
	"net/http"
)

// Top Level Exception Handler to return 500 on global exception
// Taken from a great blog post: https://go.dev/blog/error-handling-and-go
type appError struct {
	Error   error
	Message string
	Code    int
}
type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		http.Error(w, e.Message, e.Code)
	}
}

func main() {
	http.HandleFunc("/apples", appHandler(handleApples))
	http.HandleFunc("/trees", appHandler(handleTrees))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
