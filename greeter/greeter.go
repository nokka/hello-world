package greeter

import (
	"fmt"
	"regexp"
	"time"

	"github.com/nokka/hello-world/domain"
)

// Greeter encapsulates all the business logic about greeting people.
type Greeter struct{}

// The name regexp for validating names.
const nameRegexp = "^[a-zA-Z]+$"

// Greet will greet a person with the given name and validate it.
func (g Greeter) Greet(name string) (domain.Greeting, error) {
	match, _ := regexp.MatchString(nameRegexp, name)
	if !match {
		return domain.Greeting{}, fmt.Errorf("the given name can only contain letters a-z: %w", domain.ErrBadRequest)
	}

	return domain.Greeting{
		Greeting:  fmt.Sprintf("Hello there %s, nice to meet you!", name),
		GreetedAt: time.Now(),
	}, nil
}

// NewGreeter returns a new greeter with all dependencies set up.
func NewGreeter() Greeter {
	return Greeter{}
}
