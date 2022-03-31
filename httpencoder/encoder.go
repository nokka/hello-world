package httpencoder

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nokka/hello-world/domain"
)

// NewEncoder creates an instance of Encoder.
func NewEncoder() Encoder {
	return Encoder{}
}

// ResponseEncoder handles serialization over HTTP responses.
type Encoder struct{}

// ErrorResponse represents an error response sent over HTTP.
type errorResponse struct {
	Message string `json:"message"`
}

// Error utilizes the error parameter to write it over HTTP and set the corresponding status code.
func (e Encoder) Error(ctx context.Context, w http.ResponseWriter, err error) {
	resp := errorResponse{
		Message: err.Error(),
	}

	// Default to http 500.
	statusCode := http.StatusInternalServerError

	// Unwrap the error to see if it contains any other error code.
	switch errors.Unwrap(err) {
	case domain.ErrBadRequest:
		statusCode = http.StatusBadRequest
	case domain.ErrNotFound:
		statusCode = http.StatusNotFound
	case domain.ErrTemporary:
		statusCode = http.StatusServiceUnavailable
	default:
		statusCode = http.StatusInternalServerError
	}

	e.Respond(ctx, w, resp, statusCode)
}

// Respond will respond with an encoded payload if it exists otherwise just with the given status code.
func (e Encoder) Respond(ctx context.Context, w http.ResponseWriter, payload interface{}, statusCode int) {
	if payload != nil {
		encoded, err := json.Marshal(payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Write header first because the status code needs to be written last (see below).
		w.Header().Set("Content-type", "application/json; charset=utf-8")

		w.WriteHeader(statusCode)

		// Write the json encoded payload.
		w.Write(encoded)
	} else {
		// If there's no payload just write the header.
		w.WriteHeader(statusCode)
	}
}
