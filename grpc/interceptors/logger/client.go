package logger

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"
)

// ClientLoggerInterceptor is a gRPC interceptor that logs the request and response of a gRPC call.
// It will log all important information, including the error if any.
// The level of the log is based on the error if any.
// For all non-error logs, it will use the debug level.
// For all error logs, it will use the error level.
func ClientLoggerInterceptor() connect.UnaryInterceptorFunc {
	iceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			startTime := time.Now().UTC()

			resp, err := next(ctx, req)
			requestDuration := time.Since(startTime)

			logger := log.Debug()
			if err != nil {
				logger = log.Error().Err(err)
			}

			status := "OK"
			if err != nil {
				status = "ERROR"
			}

			logger.Str("protocol", "grpc").
				Str("method", req.Spec().Procedure).
				Str("status", status).
				Dur("duration", requestDuration).
				Msg("sent a grpc call")

			return resp, err
		})
	}
	return connect.UnaryInterceptorFunc(iceptor)
}
