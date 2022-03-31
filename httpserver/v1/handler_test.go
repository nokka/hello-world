package v1

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nokka/hello-world/domain"
)

func TestSayHello(t *testing.T) {
	type fields struct {
		encoder encoder
		greeter *greeterMock
	}

	type args struct {
		name string
	}

	type want struct {
		statusCode int
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		want         want
		greeterCalls int
	}{
		{
			name: "happy path",
			want: want{
				statusCode: http.StatusOK,
			},
			fields: fields{
				encoder: &encoderMock{
					RespondFunc: func(ctx context.Context, w http.ResponseWriter, payload interface{}, statusCode int) {
						w.WriteHeader(statusCode)
					},
				},
				greeter: &greeterMock{
					GreetFunc: func(name string) (domain.Greeting, error) {
						return domain.Greeting{}, nil
					},
				},
			},
			args: args{
				name: "alice",
			},
			greeterCalls: 1,
		},
		{
			name: "error path",
			want: want{
				statusCode: http.StatusBadRequest,
			},
			fields: fields{
				encoder: &encoderMock{
					ErrorFunc: func(ctx context.Context, w http.ResponseWriter, err error) {
						w.WriteHeader(http.StatusBadRequest)
					},
				},
				greeter: &greeterMock{
					GreetFunc: func(name string) (domain.Greeting, error) {
						return domain.Greeting{}, errors.New("something went wrong")
					},
				},
			},
			args: args{
				name: "bob",
			},
			greeterCalls: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				encoder: tt.fields.encoder,
				greeter: tt.fields.greeter,
			}

			res := httptest.NewRecorder()
			rq, _ := http.NewRequest(http.MethodGet, "/hello/"+tt.name, nil)

			r := chi.NewRouter()
			h.Routes(r)

			r.ServeHTTP(res, rq)

			if res.Code != tt.want.statusCode {
				t.Errorf("sayHello() got status code = %v, want %v", res.Code, tt.want.statusCode)
			}

			if len(tt.fields.greeter.GreetCalls()) != tt.greeterCalls {
				t.Errorf("expected greeter.Greet to be called exactly %d times but was called %d times",
					tt.greeterCalls,
					len(tt.fields.greeter.GreetCalls()),
				)
			}
		})
	}
}
