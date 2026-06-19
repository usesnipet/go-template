package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v == nil {
		return nil
	}
	return json.NewEncoder(w).Encode(v)
}

func WriteStatus(w http.ResponseWriter, status int) error {
	w.WriteHeader(status)
	return nil
}

func DecodeJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return NewHTTPError(http.StatusBadRequest, "request body is required")
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	return nil
}

func WriteError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	message := http.StatusText(status)

	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		status = httpErr.StatusCode
		if httpErr.Message != "" {
			message = httpErr.Message
		} else {
			message = http.StatusText(status)
		}
	} else if err != nil {
		message = err.Error()
	}

	_ = WriteJSON(w, status, ErrorResponse{
		Error:      message,
		StatusCode: status,
		StatusText: http.StatusText(status),
	})
}
