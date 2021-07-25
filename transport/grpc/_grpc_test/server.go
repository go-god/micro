package test

import (
	"context"
	"fmt"

	"github.com/go-god/micro/endpoint"
	grpctransport "github.com/go-god/micro/transport/grpc"
	"github.com/go-god/micro/transport/grpc/_grpc_test/pb"
)

// service impl
type service struct{}

// Test
func (service) Test(ctx context.Context, a string, b int64) (context.Context, string, error) {
	return nil, fmt.Sprintf("%s = %d", a, b), nil
}

// NewService create service entity
func NewService() Service {
	return service{}
}

func makeTestEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TestRequest)
		newCtx, v, err := svc.Test(ctx, req.A, req.B)
		return &TestResponse{
			V:   v,
			Ctx: newCtx,
		}, err
	}
}

// serverBinding impl Test gRPC interface.
type serverBinding struct {
	pb.UnimplementedTestServer

	test grpctransport.Handler
}

// Test grpc Test method impl
func (b *serverBinding) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	_, response, err := b.test.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return response.(*pb.TestResponse), nil
}

// NewBinding bind svc
func NewBinding(svc Service) *serverBinding {
	return &serverBinding{
		test: grpctransport.NewServer(
			makeTestEndpoint(svc),
			testEncodeDecodeEntry,
			testEncodeDecodeEntry,

			// grpc transport server options
			grpctransport.ServerBefore(
				extractCorrelationID,
			),
			grpctransport.ServerBefore(
				displayServerRequestHeaders,
			),
			grpctransport.ServerAfter(
				injectResponseHeader,
				injectResponseTrailer,
				injectConsumedCorrelationID,
			),
			grpctransport.ServerAfter(
				displayServerResponseHeaders,
				displayServerResponseTrailers,
			),
		),
	}
}
