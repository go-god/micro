package grpc

type contextKey int

const (
	// ContextKeyRequestMethod ctx key request method
	ContextKeyRequestMethod contextKey = iota
	ContextKeyIsClientStream
	ContextKeyIsServerStream
)
