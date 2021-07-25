package test

import (
	"context"

	"google.golang.org/grpc"

	"github.com/go-god/micro/endpoint"
	grpctransport "github.com/go-god/micro/transport/grpc"
	"github.com/go-god/micro/transport/grpc/_grpc_test/pb"
)

type clientBinding struct {
	test endpoint.Endpoint
}

// Test func
func (c *clientBinding) Test(ctx context.Context, a string, b int64) (context.Context, string, error) {
	response, err := c.test(ctx, TestRequest{A: a, B: b})
	if err != nil {
		return nil, "", err
	}
	r := response.(*TestResponse)
	return r.Ctx, r.V, nil
}

// NewClient new client
func NewClient(cc *grpc.ClientConn) Service {
	return &clientBinding{
		test: grpctransport.NewClient(
			cc,
			"pb.Test",
			"Test",
			testEncodeDecodeEntry,
			testEncodeDecodeEntry,
			&pb.TestResponse{},
			grpctransport.ClientBefore(
				injectCorrelationID,
			),
			grpctransport.ClientBefore(
				displayClientRequestHeaders,
			),
			grpctransport.ClientAfter(
				displayClientResponseHeaders,
				displayClientResponseTrailers,
			),
			grpctransport.ClientAfter(
				extractConsumedCorrelationID,
			),
		).Endpoint(),
	}
}
