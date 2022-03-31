package httpencoder

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nokka/hello-world/domain"
)

func TestRespond(t *testing.T) {
	tests := []struct {
		name                string
		code                int
		payload             interface{}
		expectedBody        string
		expectedContentType string
		expectedCode        int
	}{
		{
			name: "successful encoding",
			code: http.StatusOK,
			payload: struct {
				SomeField string `json:"some_field"`
			}{"hello"},
			expectedBody:        `{"some_field":"hello"}`,
			expectedContentType: "application/json; charset=utf-8",
			expectedCode:        http.StatusOK,
		},
		{
			name:                "successful encoding empty payload",
			code:                http.StatusOK,
			payload:             nil,
			expectedBody:        "",
			expectedContentType: "",
			expectedCode:        http.StatusOK,
		},
		{
			name:                "unsuccessful encoding invalid payload",
			code:                http.StatusOK,
			payload:             func() {},
			expectedBody:        "",
			expectedContentType: "",
			expectedCode:        http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := NewEncoder()
			w := httptest.NewRecorder()

			encoder.Respond(context.TODO(), w, tt.payload, tt.code)

			if got, want := w.Header().Get("Content-Type"), tt.expectedContentType; got != want {
				t.Errorf(`expected content type to be = %s, got = %s`, want, got)
			}

			if w.Code != tt.expectedCode {
				t.Errorf("expected status code to be = %d, got = %d", tt.expectedCode, w.Code)
			}

			b, err := ioutil.ReadAll(w.Body)
			if err != nil {
				t.Errorf("reading response body error = %q, want = %q", err, tt.expectedBody)
			}

			// Trim the newline added by the json encoder, otherwise we can't compary the body.
			decoded := strings.TrimSuffix(string(b), "\n")
			if decoded != tt.expectedBody {
				t.Errorf("expected body to be = %q, got = %q", tt.expectedBody, decoded)
			}
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode int
	}{
		{
			name:         "resource not found",
			err:          fmt.Errorf("resource not found: %w", domain.ErrNotFound),
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "bad request data",
			err:          fmt.Errorf("unable to parse UUID: %w", domain.ErrBadRequest),
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "default to internal server error",
			err:          errors.New("error contains no domain error"),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:         "temporary error occured",
			err:          fmt.Errorf("something went temporarily wrong: %w", domain.ErrTemporary),
			expectedCode: http.StatusServiceUnavailable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := NewEncoder()
			w := httptest.NewRecorder()

			encoder.Error(context.TODO(), w, tt.err)

			if w.Code != tt.expectedCode {
				t.Errorf("expected status code to be = %d, got = %d", tt.expectedCode, w.Code)
			}
		})
	}
}
