package transport

import (
	"context"
	"errors"
	"log"
	"testing"

	glog "github.com/go-god/micro/log"
)

func TestLogErrorHandler(t *testing.T) {
	logger := glog.LoggerFunc(log.Printf)

	errorHandler := NewLogErrorHandler(logger)

	err := errors.New("request error")

	errorHandler.Handle(context.Background(), err)
}
