package transport

import "context"

// EncodeRequestStream request encode func interface
type EncodeRequestStream interface {
	EncodeRequest(context.Context, interface{}) (request interface{}, err error)
}

// DecodeRequestStream request decode func interface
type DecodeRequestStream interface {
	DecodeRequest(context.Context, interface{}) (request interface{}, err error)
}

// EncodeResponseStream response encode func interface
type EncodeResponseStream interface {
	EncodeResponse(context.Context, interface{}) (response interface{}, err error)
}

// DecodeResponseStream response decode func interface
type DecodeResponseStream interface {
	DecodeResponse(context.Context, interface{}) (response interface{}, err error)
}
