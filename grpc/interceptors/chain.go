package grpc_interceptors

import (
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/goauthnz/pkg/grpc/interceptors/logger"
	"github.com/goauthnz/pkg/grpc/interceptors/recover"
	"github.com/goauthnz/pkg/grpc/interceptors/timeout"
)

type Interceptors []connect.Interceptor

const (
	DefaultTimeout = 10 * time.Second
)

func ServerDefaultInterceptors() Interceptors {
	return []connect.Interceptor{
		recover.RecoverInterceptor(),
	}
}

func ClientDefaultInterceptors(t time.Duration) Interceptors {
	return []connect.Interceptor{
		timeout.TimeoutInterceptor(t),
		logger.ClientLoggerInterceptor(),
	}
}
