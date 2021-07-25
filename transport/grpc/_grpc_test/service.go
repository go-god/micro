package test

import "context"

// Service
type Service interface {
	Test(ctx context.Context, a string, b int64) (context.Context, string, error)
}

// TestRequest request
type TestRequest struct {
	A string
	B int64
}

// TestResponse response
type TestResponse struct {
	Ctx context.Context
	V   string
}
