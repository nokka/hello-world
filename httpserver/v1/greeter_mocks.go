// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package v1

import (
	"context"
	"github.com/nokka/hello-world/domain"
	"net/http"
	"sync"
)

// Ensure, that encoderMock does implement encoder.
// If this is not the case, regenerate this file with moq.
var _ encoder = &encoderMock{}

// encoderMock is a mock implementation of encoder.
//
// 	func TestSomethingThatUsesencoder(t *testing.T) {
//
// 		// make and configure a mocked encoder
// 		mockedencoder := &encoderMock{
// 			ErrorFunc: func(ctx context.Context, w http.ResponseWriter, err error)  {
// 				panic("mock out the Error method")
// 			},
// 			RespondFunc: func(ctx context.Context, w http.ResponseWriter, payload interface{}, statusCode int)  {
// 				panic("mock out the Respond method")
// 			},
// 		}
//
// 		// use mockedencoder in code that requires encoder
// 		// and then make assertions.
//
// 	}
type encoderMock struct {
	// ErrorFunc mocks the Error method.
	ErrorFunc func(ctx context.Context, w http.ResponseWriter, err error)

	// RespondFunc mocks the Respond method.
	RespondFunc func(ctx context.Context, w http.ResponseWriter, payload interface{}, statusCode int)

	// calls tracks calls to the methods.
	calls struct {
		// Error holds details about calls to the Error method.
		Error []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// W is the w argument value.
			W http.ResponseWriter
			// Err is the err argument value.
			Err error
		}
		// Respond holds details about calls to the Respond method.
		Respond []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// W is the w argument value.
			W http.ResponseWriter
			// Payload is the payload argument value.
			Payload interface{}
			// StatusCode is the statusCode argument value.
			StatusCode int
		}
	}
	lockError   sync.RWMutex
	lockRespond sync.RWMutex
}

// Error calls ErrorFunc.
func (mock *encoderMock) Error(ctx context.Context, w http.ResponseWriter, err error) {
	if mock.ErrorFunc == nil {
		panic("encoderMock.ErrorFunc: method is nil but encoder.Error was just called")
	}
	callInfo := struct {
		Ctx context.Context
		W   http.ResponseWriter
		Err error
	}{
		Ctx: ctx,
		W:   w,
		Err: err,
	}
	mock.lockError.Lock()
	mock.calls.Error = append(mock.calls.Error, callInfo)
	mock.lockError.Unlock()
	mock.ErrorFunc(ctx, w, err)
}

// ErrorCalls gets all the calls that were made to Error.
// Check the length with:
//     len(mockedencoder.ErrorCalls())
func (mock *encoderMock) ErrorCalls() []struct {
	Ctx context.Context
	W   http.ResponseWriter
	Err error
} {
	var calls []struct {
		Ctx context.Context
		W   http.ResponseWriter
		Err error
	}
	mock.lockError.RLock()
	calls = mock.calls.Error
	mock.lockError.RUnlock()
	return calls
}

// Respond calls RespondFunc.
func (mock *encoderMock) Respond(ctx context.Context, w http.ResponseWriter, payload interface{}, statusCode int) {
	if mock.RespondFunc == nil {
		panic("encoderMock.RespondFunc: method is nil but encoder.Respond was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		W          http.ResponseWriter
		Payload    interface{}
		StatusCode int
	}{
		Ctx:        ctx,
		W:          w,
		Payload:    payload,
		StatusCode: statusCode,
	}
	mock.lockRespond.Lock()
	mock.calls.Respond = append(mock.calls.Respond, callInfo)
	mock.lockRespond.Unlock()
	mock.RespondFunc(ctx, w, payload, statusCode)
}

// RespondCalls gets all the calls that were made to Respond.
// Check the length with:
//     len(mockedencoder.RespondCalls())
func (mock *encoderMock) RespondCalls() []struct {
	Ctx        context.Context
	W          http.ResponseWriter
	Payload    interface{}
	StatusCode int
} {
	var calls []struct {
		Ctx        context.Context
		W          http.ResponseWriter
		Payload    interface{}
		StatusCode int
	}
	mock.lockRespond.RLock()
	calls = mock.calls.Respond
	mock.lockRespond.RUnlock()
	return calls
}

// Ensure, that greeterMock does implement greeter.
// If this is not the case, regenerate this file with moq.
var _ greeter = &greeterMock{}

// greeterMock is a mock implementation of greeter.
//
// 	func TestSomethingThatUsesgreeter(t *testing.T) {
//
// 		// make and configure a mocked greeter
// 		mockedgreeter := &greeterMock{
// 			GreetFunc: func(name string) (domain.Greeting, error) {
// 				panic("mock out the Greet method")
// 			},
// 		}
//
// 		// use mockedgreeter in code that requires greeter
// 		// and then make assertions.
//
// 	}
type greeterMock struct {
	// GreetFunc mocks the Greet method.
	GreetFunc func(name string) (domain.Greeting, error)

	// calls tracks calls to the methods.
	calls struct {
		// Greet holds details about calls to the Greet method.
		Greet []struct {
			// Name is the name argument value.
			Name string
		}
	}
	lockGreet sync.RWMutex
}

// Greet calls GreetFunc.
func (mock *greeterMock) Greet(name string) (domain.Greeting, error) {
	if mock.GreetFunc == nil {
		panic("greeterMock.GreetFunc: method is nil but greeter.Greet was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockGreet.Lock()
	mock.calls.Greet = append(mock.calls.Greet, callInfo)
	mock.lockGreet.Unlock()
	return mock.GreetFunc(name)
}

// GreetCalls gets all the calls that were made to Greet.
// Check the length with:
//     len(mockedgreeter.GreetCalls())
func (mock *greeterMock) GreetCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockGreet.RLock()
	calls = mock.calls.Greet
	mock.lockGreet.RUnlock()
	return calls
}