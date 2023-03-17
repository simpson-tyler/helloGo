package api

import (
	"encoding/json"
	"helloGo/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleApples(w http.ResponseWriter, r *http.Request) *AppError {
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)
		id_str, ok := vars["id"]
		id, atoi_err := strconv.Atoi(id_str)
		if !ok || atoi_err != nil {
			return &AppError{http.ErrNotSupported, "Record not Found", 404}
		}

		record, err := model.FetchAppleByID(uint(id))
		if err != nil {
			return &AppError{err, "Record Not Found", 404}
		}
		apple_enc, err := json.Marshal(record)
		if err != nil {
			return &AppError{err, "Could not encode to Json", 404}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(apple_enc)
	default:
		return &AppError{http.ErrNotSupported, "Not Supported!", 500}
	}

	return nil
}

func handleTrees(w http.ResponseWriter, r *http.Request) *AppError {
	return &AppError{http.ErrNotSupported, "Not Supported!", 500}
}
