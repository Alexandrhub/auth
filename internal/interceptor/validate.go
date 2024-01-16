package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

func ValidateInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if v, ok := req.(validator); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}
