package api

import (
	"net/http"
)

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
