package grpc

import "context"

// EncodeRequestStream request encode func interface
type EncodeRequestStream interface {
	EncodeRequest(context.Context, interface{}) (request interface{}, err error)
}

// EncodeRequestFunc impl EncodeRequestStream
type EncodeRequestFunc func(context.Context, interface{}) (request interface{}, err error)

// EncodeRequest function interface
func (fn EncodeRequestFunc) EncodeRequest(ctx context.Context, req interface{}) (request interface{}, err error) {
	return fn(ctx, req)
}

// DecodeRequestStream request decode func interface
type DecodeRequestStream interface {
	DecodeRequest(context.Context, interface{}) (request interface{}, err error)
}

// DecodeRequestFunc func impl DecodeRequestStream
type DecodeRequestFunc func(context.Context, interface{}) (request interface{}, err error)

// DecodeRequest
func (fn DecodeRequestFunc) DecodeRequest(ctx context.Context, req interface{}) (request interface{}, err error) {
	return fn(ctx, req)
}

// EncodeResponseStream response encode func interface
type EncodeResponseStream interface {
	EncodeResponse(context.Context, interface{}) (response interface{}, err error)
}

// EncodeResponseFunc func impl EncodeResponseStream
type EncodeResponseFunc func(context.Context, interface{}) (response interface{}, err error)

// EncodeResponse
func (fn EncodeResponseFunc) EncodeResponse(ctx context.Context, res interface{}) (response interface{}, err error) {
	return fn(ctx, res)
}

// DecodeResponseStream response decode func interface
type DecodeResponseStream interface {
	DecodeResponse(context.Context, interface{}) (response interface{}, err error)
}

// DecodeResponseFunc func impl DecodeResponseStream
type DecodeResponseFunc func(context.Context, interface{}) (response interface{}, err error)

// DecodeResponse
func (fn DecodeResponseFunc) DecodeResponse(ctx context.Context, res interface{}) (response interface{}, err error) {
	return fn(ctx, res)
}
