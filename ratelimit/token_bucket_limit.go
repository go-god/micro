package ratelimit

import (
	"context"

	"github.com/go-god/micro/endpoint"
)

// Waiter dictates how long a request must be delayed.
// The Limiter from "golang.org/x/time/rate" already implements this interface,
// one is able to use that in NewDelayingLimiter without any modifications.
type Waiter interface {
	Wait(ctx context.Context) error
}

// WaiterFunc is an adapter that lets a function operate as if it implements Waiter
type WaiterFunc func(ctx context.Context) error

// Wait makes the adapter implement Waiter
func (f WaiterFunc) Wait(ctx context.Context) error {
	return f(ctx)
}

// NewDelayingLimiter returns an endpoint.Middleware that acts as a request throttler.
// Requests that would exceed the maximum request rate are delayed via the Waiter function
func NewDelayingLimiter(w Waiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			if err := w.Wait(ctx); err != nil {
				return nil, err
			}

			return next(ctx, request)
		}
	}
}
