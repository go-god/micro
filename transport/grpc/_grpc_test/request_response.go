package test

import (
	"context"

	"github.com/go-god/micro/transport/grpc/_grpc_test/pb"
)

type testEncodeDecode struct{}

var testEncodeDecodeEntry = testEncodeDecode{}

// EncodeRequest
func (testEncodeDecode) EncodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(TestRequest)
	return &pb.TestRequest{A: r.A, B: r.B}, nil
}

// DecodeRequest
func (testEncodeDecode) DecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb.TestRequest)
	return TestRequest{A: r.A, B: r.B}, nil
}

// EncodeResponse
func (testEncodeDecode) EncodeResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*TestResponse)
	return &pb.TestResponse{V: r.V}, nil
}

// DecodeResponse
func (testEncodeDecode) DecodeResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	r := resp.(*pb.TestResponse)
	return &TestResponse{V: r.V, Ctx: ctx}, nil
}
