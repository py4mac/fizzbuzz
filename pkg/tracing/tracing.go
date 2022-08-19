// Package tracing holds jaeger implementation for otel
package tracing

import (
	"fmt"

	"github.com/py4mac/fizzbuzz/config"
	"github.com/py4mac/fizzbuzz/pkg/constants"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

// NewTracing return TraceProvider instance
func NewTracing(cfg *config.Config) (*tracesdk.TracerProvider, error) {
	endpoint := fmt.Sprintf("http://%s:%s/api/traces", cfg.Jaeger.Host, cfg.Jaeger.Port)
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(constants.Name),
		)),
	)

	return tp, nil
}
