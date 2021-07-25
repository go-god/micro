package main

import (
	"context"
	"fmt"
	"log"

	test "github.com/go-god/micro/transport/grpc/_grpc_test"
	"github.com/go-god/micro/transport/grpc/_grpc_test/pb"

	"google.golang.org/grpc"
)

const (
	hostPort string = "localhost:8002"
)

func main() {
	cc, err := grpc.Dial(hostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to Dial: %+v", err)
	}

	defer cc.Close()

	client := test.NewClient(cc)

	var (
		a   = "the answer to life the universe and everything"
		b   = int64(42)
		cID = "request-1"
		ctx = test.SetCorrelationID(context.Background(), cID)
	)

	responseCTX, v, err := client.Test(ctx, a, b)
	if err != nil {
		log.Fatalf("unable to Test: %+v", err)
	}
	if want, have := fmt.Sprintf("%s = %d", a, b), v; want != have {
		log.Fatalf("want %q, have %q", want, have)
	}

	if want, have := cID, test.GetConsumedCorrelationID(responseCTX); want != have {
		log.Fatalf("want %q, have %q", want, have)
	}

	// client call of grpc pb itself
	grpcClient := pb.NewTestClient(cc)
	resp, err := grpcClient.Test(ctx, &pb.TestRequest{
		A: a,
		B: b,
	})

	log.Println("resp: ", resp)
	log.Println("resp error:  ", err)
}

/*
 Client found correlationID "request-1" in context, set metadata header
        Client >> Request Headers:
                correlation-id: request-1
        Client << Response Headers:
                content-type: application/grpc
                my-response-header: has-a-value
        Client << Response Trailers:
                my-response-trailer: has-a-value-too
                correlation-id-consumed: request-1
        Client received consumed correlationID "request-1" in metadata trailer, set context
2021/07/25 21:33:23 resp:  v:"the answer to life the universe and everything = 42"
2021/07/25 21:33:23 resp error:   <nil>
*/
