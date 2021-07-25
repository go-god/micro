package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	"github.com/go-god/micro"
	test "github.com/go-god/micro/transport/grpc/_grpc_test"
	"github.com/go-god/micro/transport/grpc/_grpc_test/pb"
)

const (
	hostPort string = "localhost:8002"
)

var waitTime = 5 * time.Second

func main() {
	var (
		server  = grpc.NewServer()
		service = test.NewService()
	)

	sc, err := net.Listen("tcp", hostPort)
	if err != nil {
		log.Fatalf("unable to listen: %+v", err)
	}

	go func() {
		pb.RegisterTestServer(server, test.NewBinding(service))
		err = server.Serve(sc)
		if err != nil {
			log.Fatalf("grpc start serve error: %s", err.Error())
		}
	}()

	log.Printf("grpc server run on:%s pid:%d", hostPort, os.Getppid())

	// intercept interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, micro.InterruptSignals...)

	sig := <-sigChan
	signal.Stop(sigChan)

	log.Printf("recv signal %s", sig.String())
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// if your application should wait for other services
	// to finalize based on context cancellation.
	go server.GracefulStop()
	<-ctx.Done()

	log.Println("server shutting down")
}

/*
 Server received correlationID "request-1" in metadata header, set context
        Server << Request Headers:
                :authority: localhost:8002
                content-type: application/grpc
                user-agent: grpc-go/1.39.0
                correlation-id: request-1
        Server found correlationID "request-1" in context, set consumed trailer
        Server >> Response Headers:
                my-response-header: has-a-value
        Server >> Response Trailers:
                my-response-trailer: has-a-value-too
                correlation-id-consumed: request-1
        Server << Request Headers:
                content-type: application/grpc
                user-agent: grpc-go/1.39.0
                :authority: localhost:8002
        Server >> Response Headers:
                my-response-header: has-a-value
        Server >> Response Trailers:
                my-response-trailer: has-a-value-too
*/
