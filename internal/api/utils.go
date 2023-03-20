package api

import (
	"net/http"
)

// Taken from a great blog post: https://go.dev/blog/error-handling-and-go
type AppError struct {
	Error   error
	Message string
	Code    int
}
type AppHandler func(http.ResponseWriter, *http.Request) *AppError

func (fn AppHandler) CustomServeHttp(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		http.Error(w, e.Message, e.Code)
	}
}
