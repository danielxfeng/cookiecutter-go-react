package util

import (
	"encoding/json"
	"net/http"
)

func ParseJSON[T any](r *http.Request) (T, error) {
	var data T
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}

func SendJSON[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// It would be a 500 so recovery handles the rest.
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
