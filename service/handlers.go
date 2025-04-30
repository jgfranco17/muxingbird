package service

import (
	"encoding/json"
	"net/http"
)

func CreateNewMockHandler(statusCode int, content any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(content)
		if err != nil {
			http.Error(w, "Failed to serve mock endpoint", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
	}
}
