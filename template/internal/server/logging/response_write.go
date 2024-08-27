package logging

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func Trace(ctx context.Context, spanName string) trace.Span {
	tr := otel.GetTracerProvider().Tracer("example/server")
	_, span := tr.Start(ctx, spanName, trace.WithSpanKind(trace.SpanKindServer))
	return span
}
