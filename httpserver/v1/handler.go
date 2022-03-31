package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nokka/hello-world/domain"
)

//go:generate moq -out greeter_mocks.go . encoder greeter

type encoder interface {
	Respond(ctx context.Context, w http.ResponseWriter, payload interface{}, statusCode int)
	Error(ctx context.Context, w http.ResponseWriter, err error)
}

type greeter interface {
	Greet(name string) (domain.Greeting, error)
}

type Handler struct {
	encoder encoder
	greeter greeter
}

func (h Handler) Routes(r chi.Router) {
	r.Get("/hello/{name}", h.sayHello)
}

type helloResponse struct {
	Greeting  string    `json:"greeting"`
	GreetedAt time.Time `json:"greeted_at"`
}

func (h Handler) sayHello(w http.ResponseWriter, r *http.Request) {
	// Get request context.
	ctx := r.Context()

	// Get the name from the URL param {name}.
	name := chi.URLParam(r, "name")

	// Call the greeter to construct a proper greeting.
	greeting, err := h.greeter.Greet(name)
	if err != nil {
		h.encoder.Error(ctx, w, err)
		return
	}

	resp := helloResponse{
		Greeting:  greeting.Greeting,
		GreetedAt: greeting.GreetedAt,
	}

	h.encoder.Respond(ctx, w, resp, http.StatusOK)
}

// NewHandler returns a new v1 handler with all dependencies.
func NewHandler(encoder encoder, greeter greeter) Handler {
	return Handler{
		encoder: encoder,
		greeter: greeter,
	}
}
