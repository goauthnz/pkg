package grpc_interceptors

import (
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/goauthnz/pkg/grpc/interceptors/logger"
	"github.com/goauthnz/pkg/grpc/interceptors/recover"
	"github.com/goauthnz/pkg/grpc/interceptors/timeout"
)

type InterceptorsChain []connect.Interceptor

const (
	DefaultTimeout = 10 * time.Second
)

func ServerDefaultChain() InterceptorsChain {
	return []connect.Interceptor{
		recover.RecoverInterceptor(),
	}
}

func ClientDefaultChain(t time.Duration) InterceptorsChain {
	return []connect.Interceptor{
		timeout.TimeoutInterceptor(t),
		logger.ClientLoggerInterceptor(),
	}
}
