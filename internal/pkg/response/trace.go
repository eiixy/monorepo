package response

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/otel/trace"
)

func AppendTraceId(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if header, ok := transport.FromServerContext(ctx); ok {
			spanCtx := trace.SpanContextFromContext(ctx)
			header.ReplyHeader().Set("X-Trace-Id", spanCtx.TraceID().String())
		}
		return handler(ctx, req)
	}
}
