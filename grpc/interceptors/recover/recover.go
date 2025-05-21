package recover

import (
	"context"
	"fmt"
	"runtime"

	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"
)

type PanicError struct {
	Panic any
	Stack []byte
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("panic caught: %v\n\n%s", e.Panic, e.Stack)
}

// RecoverInterceptor recovers from panics and returns an error.
func RecoverInterceptor() connect.UnaryInterceptorFunc {
	iceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (_ connect.AnyResponse, err error) {
			defer func() {
				if r := recover(); r != nil {
					log.Error().Msg("grpc.interceptors.recover: panic caught")
					err = recoverFrom(ctx, r)
				}
			}()

			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(iceptor)
}

func recoverFrom(_ context.Context, p any) error {
	stack := make([]byte, 64<<10)
	stack = stack[:runtime.Stack(stack, false)]
	return &PanicError{Panic: p, Stack: stack}
}
