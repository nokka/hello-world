package domain

import "time"

// Greeting holds all information related to greeting someone.
type Greeting struct {
	Greeting  string
	GreetedAt time.Time
}
