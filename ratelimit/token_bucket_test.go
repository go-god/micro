package ratelimit

import (
	"context"
	"strings"
	"testing"
	"time"

	"golang.org/x/time/rate"

	"github.com/go-god/micro/endpoint"
)

// TestNewRateLimiter test rate limit.
func TestNewRateLimiter(t *testing.T) {
	limit := rate.NewLimiter(rate.Every(time.Minute), 1)
	testSuccessThenFailure(t, NewRateLimiter(limit)(endpoint.NoEndpoint), ErrLimited.Error())
}

// TestNewDelayingLimiter NewDelayingLimiter wait limit.
func TestNewDelayingLimiter(t *testing.T) {
	limit := rate.NewLimiter(rate.Every(time.Minute), 1)
	testSuccessThenFailure(t, NewDelayingLimiter(limit)(endpoint.NoEndpoint), "exceed context deadline")
}

func testSuccessThenFailure(t *testing.T, e endpoint.Endpoint, failContains string) {
	ctx, cxl := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cxl()

	// First request should succeed.
	if _, err := e(ctx, struct{}{}); err != nil {
		t.Errorf("unexpected: %v\n", err)
	}

	// Next request should fail.
	if _, err := e(ctx, struct{}{}); !strings.Contains(err.Error(), failContains) {
		t.Errorf("expected `%s`: %v\n", failContains, err)
	}

}
