package timeout

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
)

// TimeoutInterceptor uses the context.WithTimeout to cancel the request if it takes too long.
func TimeoutInterceptor(timeout time.Duration) connect.UnaryInterceptorFunc {
	iceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(iceptor)
}
