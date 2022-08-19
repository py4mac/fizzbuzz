// Package tracing holds jaeger implementation for otel
package tracing

import (
	"github.com/py4mac/fizzbuzz/config"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

// NewTracing return TraceProvider instance
func NewTracing(cfg *config.Config) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.Jaeger.Host)))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.Jaeger.ServiceName),
		)),
	)

	return tp, nil
}
