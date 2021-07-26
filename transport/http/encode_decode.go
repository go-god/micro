package http

import (
	"context"
	"net/http"
)

// CreateRequestStream create client request interface
type CreateRequestStream interface {
	// CreateRequest creates an outgoing HTTP request based on the passed
	// request object. It's designed to be used in HTTP clients, for client-side
	// endpoints. It's a more powerful version of EncodeRequest, and can be used
	// if more fine-grained control of the HTTP request is required.
	CreateRequest(context.Context, interface{}) (*http.Request, error)
}

// CreateRequestFunc type is an adapter to allow the use of
// ordinary function as CreateRequestStream. If f is a function
// with the appropriate signature, CreateRequestFunc(f) is a
// CreateRequestFunc that calls f.
type CreateRequestFunc func(context.Context, interface{}) (*http.Request, error)

// CreateRequest calls f(ctx, err).
func (fn CreateRequestFunc) CreateRequest(ctx context.Context, req interface{}) (*http.Request, error) {
	return fn(ctx, req)
}

// EncodeRequestStream request encode func interface
type EncodeRequestStream interface {
	// EncodeRequest encodes the passed request object into the HTTP request
	// object. It's designed to be used in HTTP clients, for client-side
	// endpoints. One straightforward EncodeRequest could be something that JSON
	// encodes the object directly to the request body.
	EncodeRequest(context.Context, *http.Request, interface{}) error
}

// EncodeRequestFunc impl EncodeRequestStream
type EncodeRequestFunc func(context.Context, *http.Request, interface{}) error

// EncodeRequest fn call
func (fn EncodeRequestFunc) EncodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	return fn(ctx, r, request)
}

// DecodeRequestStream request decode func interface
type DecodeRequestStream interface {
	// DecodeRequest extracts a user-domain request object from an HTTP
	// request object. It's designed to be used in HTTP servers, for server-side
	// endpoints. One straightforward DecodeRequest could be something that
	// JSON decodes from the request body to the concrete request type.
	DecodeRequest(context.Context, *http.Request) (request interface{}, err error)
}

// DecodeRequestFunc decode request func
type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)

// DecodeRequest impl DecodeRequest
func (fn DecodeRequestFunc) DecodeRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return fn(ctx, r)
}

// EncodeResponseStream response encode func interface
type EncodeResponseStream interface {
	// EncodeResponse encodes the passed response object to the HTTP response
	// writer. It's designed to be used in HTTP servers, for server-side
	// endpoints. One straightforward EncodeResponse could be something that
	// JSON encodes the object directly to the response body.
	EncodeResponse(context.Context, http.ResponseWriter, interface{}) error
}

// EncodeResponseFunc encode response func
type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error

// EncodeResponse impl encode response
func (fn EncodeResponseFunc) EncodeResponse(ctx context.Context, w http.ResponseWriter, res interface{}) error {
	return fn(ctx, w, res)
}

// DecodeResponseStream response decode func interface
type DecodeResponseStream interface {
	// DecodeResponse extracts a user-domain response object from an HTTP
	// response object. It's designed to be used in HTTP clients, for client-side
	// endpoints. One straightforward DecodeResponse could be something that
	// JSON decodes from the response body to the concrete response type.
	DecodeResponse(context.Context, *http.Response) (response interface{}, err error)
}

// DecodeResponseFunc impl DecodeResponseStream
type DecodeResponseFunc func(context.Context, *http.Response) (response interface{}, err error)

// DecodeResponse fn(ctx,res)
func (fn DecodeResponseFunc) DecodeResponse(ctx context.Context, res *http.Response) (response interface{},
	err error) {
	return fn(ctx, res)
}
