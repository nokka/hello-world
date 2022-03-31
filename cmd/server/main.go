package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nokka/hello-world/greeter"
	"github.com/nokka/hello-world/httpencoder"
	"github.com/nokka/hello-world/httpserver"
	v1 "github.com/nokka/hello-world/httpserver/v1"
)

func main() {
	var (
		httpAddress = envString("HTTP_ADDRESS", ":8085")
	)

	// Services.
	greeterService := greeter.NewGreeter()

	// Encoder to encode all http responses and errors.
	httpEncoder := httpencoder.NewEncoder()

	// Handler that will handle all HTTP reqeusts on the /v1 routes.
	v1Handler := v1.NewHandler(httpEncoder, greeterService)

	// Channel to receive errors on from different go routines, such as the http server.
	errorChannel := make(chan error)

	// Open HTTP server and send it on the error channel, will stall until any error is returned.
	go func() {
		server := httpserver.NewServer(httpAddress, 2*time.Second, v1Handler)
		errorChannel <- server.Open()
	}()

	// Capture interupts, to handle them gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorChannel <- fmt.Errorf("got terminating signal: %s", <-c)
	}()

	// Wait for errors on the error channel, this will stall until an error is received.
	if err := <-errorChannel; err != nil {
		log.Fatal(err)
	}
}

func envString(key string, fallback string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return fallback
}
