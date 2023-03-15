package api

import (
	"models/appleFactory"
	"net/http"

	"github.com/gorilla/mux"
)

func handleApples(w http.ResponseWriter, r *http.Request) *appError {
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			return &appError{http.ErrNotSupported, "Record not Found", 404}
		}
		record, err := appleFactory.FetchByID(id)
		if err != nil {
			return &appError{err.Error(), "Record Not Found", 404}
		}
		return record, nil
	default:
		return &appError{http.ErrNotSupported, "Not Supported!", 500}
	}
}

func handleTrees(w http.ResponseWriter, r *http.Request) *appError {
	return &appError{http.ErrNotSupported, "Not Supported!", 500}
}
