package httpserver

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type handler interface {
	Routes(r chi.Router)
}

type Server struct {
	address   string
	timeout   time.Duration
	v1Handler handler
}

// NewServer returns a new HTTP server with all dependencies setup.
func NewServer(address string, timeout time.Duration, v1Handler handler) Server {
	return Server{
		address:   address,
		timeout:   timeout,
		v1Handler: v1Handler,
	}
}

// Open setups a TCP listener on the servers address and serves http requests.
func (s Server) Open() error {
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	server := http.Server{
		Handler: http.TimeoutHandler(s.handler(), s.timeout, "http request timed out"),
	}

	log.Printf("http server listening on %s", s.address)

	return server.Serve(ln)
}

func (s Server) handler() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Health endpoint.
	router.Use(middleware.Heartbeat("/health"))

	router.Route("/api", func(r chi.Router) {
		r.Route("/v1", s.v1Handler.Routes)
	})

	return router
}
