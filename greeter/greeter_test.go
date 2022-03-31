package greeter

import (
	"errors"
	"testing"

	"github.com/nokka/hello-world/domain"
)

func TestGreet(t *testing.T) {
	type args struct {
		name string
	}

	tests := []struct {
		name             string
		args             args
		expectedGreeting string
		expectedError    error
	}{
		{
			name: "valid greeting",
			args: args{
				name: "alice",
			},
			expectedGreeting: "Hello there alice, nice to meet you!",
			expectedError:    nil,
		},
		{
			name: "invalid greeting",
			args: args{
				name: "alice123",
			},
			expectedError: domain.ErrBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			greeter := Greeter{}

			greeting, err := greeter.Greet(tt.args.name)

			if err != nil && tt.expectedError == nil {
				t.Errorf("didn't expect an error, got = %v", err)
			}

			if tt.expectedError != nil && errors.Unwrap(err) != tt.expectedError {
				t.Errorf("Expected error to be = %v, got = %#v", tt.expectedError, errors.Unwrap(err))
			}

			if greeting.Greeting != tt.expectedGreeting {
				t.Errorf("expected greeting = %s; got = %s", tt.expectedGreeting, greeting.Greeting)
			}
		})
	}
}
