package ratelimit

import (
	"context"
	"errors"

	"github.com/go-god/micro/endpoint"
)

// Limiter dictates whether or not a request is acceptable to run.
// If Limit function return true, the request will be rejected.
// Otherwise, the request will pass.
//
// The Limiter from "golang.org/x/time/rate" already implements this interface,
// one is able to use that in WithRateLimit without any modifications.
type Limiter interface {
	Allow() bool
}

// LimiterFunc is an adapter that lets a function operate as if it implements Limiter
type LimiterFunc func() bool

// Allow makes the adapter implement Limiter
func (f LimiterFunc) Allow() bool {
	return f()
}

var (
	// ErrLimited is returned in the request path when the rate limiter is
	// triggered and the request is rejected.
	ErrLimited = errors.New("rate limit exceeded")
)

// NewRateLimiter returns an endpoint.Middleware that acts as a rate limiter.
// Requests that would exceed the maximum request rate are simply rejected with an error.
func NewRateLimiter(limiter Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			if !limiter.Allow() {
				return nil, ErrLimited
			}

			return next(ctx, request)
		}
	}
}
