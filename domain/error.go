package domain

// Error is a package specific error type that lets us define immutable errors.
type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// ErrBadRequest is returned when the request was invalid.
	ErrBadRequest = Error("invalid request data")

	// ErrNotFound is returned when a resource can't be find.
	ErrNotFound = Error("resource was not found")

	// ErrTemporary is returned when something has gone wrong but it's only temporary.
	ErrTemporary = Error("temporary error")

	// ErrInternal is returned when the error is unspecified.
	ErrInternal = Error("internal error")
)
