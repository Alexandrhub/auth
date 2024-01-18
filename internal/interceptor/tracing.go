package interceptor

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const traceIDKey = "x-trace-id"

func ServerTracingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, info.FullMethod)
	defer span.Finish()

	spanCtx, ok := span.Context().(jaeger.SpanContext)
	if ok {
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(traceIDKey, spanCtx.TraceID().String()))

		header := metadata.New(map[string]string{traceIDKey: spanCtx.TraceID().String()})
		err := grpc.SendHeader(ctx, header)
		if err != nil {
			return nil, err
		}
	}

	resp, err := handler(ctx, req)
	if err != nil {
		ext.Error.Set(span, true)
		span.SetTag("err", err.Error())
	} else {
		// just for testing, in prod resp will be too large
		span.SetTag("res", resp)
	}

	return resp, err
}
