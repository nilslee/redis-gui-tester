package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func DecodeRequestJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// 1. Decode the body
	err := json.NewDecoder(r.Body).Decode(dst)
	if err == nil {
		return nil
	}

	// 2. Classify errors
	var syntaxErr *json.SyntaxError
	var unmarshalErr *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxErr):
		http.Error(w, "Malformed JSON syntax", http.StatusBadRequest)
	case errors.As(err, &unmarshalErr):
		// 422 is better for valid JSON with the wrong types
		http.Error(w, "Invalid data types in JSON", http.StatusUnprocessableEntity)
	case errors.Is(err, io.EOF):
		http.Error(w, "Request body must not be empty", http.StatusBadRequest)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	return err
}
