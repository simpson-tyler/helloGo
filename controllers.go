package main

import (
	"net/http"
)

func handleApples(w http.ResponseWriter, r *http.Request) *appError {
	switch r.Method {
	case http.MethodGet:
		appleId := r.Response.Body
		record, err := appleFactory.FetchByID(appleId)

		if err != nil {
			return &appError{err.Error(), "Record Not Found", 404}
		}
	default:
		return &appError{http.ErrNotSupported, "Not Supported!", 500}
	}
}

func handleTrees(w http.ResponseWriter, r *http.Request) *appError {
	return &appError{http.ErrNotSupported, "Not Supported!", 500}
}
