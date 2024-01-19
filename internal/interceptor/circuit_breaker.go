package interceptor

import (
	"context"
	"errors"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CircuitBreakerInterceptor struct {
	cb *gobreaker.CircuitBreaker
}

func NewCircuitBreakerInterceptor(cb *gobreaker.CircuitBreaker) *CircuitBreakerInterceptor {
	return &CircuitBreakerInterceptor{
		cb: cb,
	}
}

func (c *CircuitBreakerInterceptor) Unary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	resp, err := c.cb.Execute(
		func() (any, error) {
			return handler(ctx, req)
		},
	)
	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return nil, status.Error(codes.Unavailable, "service is unavailable")
		}

		return nil, err
	}

	return resp, nil
}
