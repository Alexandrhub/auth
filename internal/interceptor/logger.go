package interceptor

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/alexandrhub/auth/internal/logger"
)

func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	now := time.Now()

	resp, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(), zap.String("method", info.FullMethod), zap.Any("req", req))
	}

	logger.Info(
		"request",
		zap.String("method", info.FullMethod),
		zap.Duration("duration", time.Since(now)),
		zap.Any("req", req),
		zap.Any("resp", resp),
	)

	return resp, err
}
