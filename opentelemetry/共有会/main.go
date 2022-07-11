package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"

	_ "go.opencensus.io/resource"
	_ "go.opencensus.io/trace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"

	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/otel"
)

func main() {
	tracerProvider, err := NewTracerProvider("otelhttp_client_trace")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tracerProvider)

	ctx := context.Background()
	ctx, span := tracerProvider.Tracer("main").Start(ctx, "main")
	defer span.End()

	if err := httpGet(ctx, "https://journal.lampetty.net/"); err != nil {
		log.Fatal(err)
	}
}

func httpGet(ctx context.Context, url string) error {
	ctx, span := otel.Tracer("main").Start(ctx, "httpGet")
	defer span.End()
	span.SetAttributes(attribute.Key("url").String(url))

	clientTrace := otelhttptrace.NewClientTrace(ctx)
	ctx = httptrace.WithClientTrace(ctx, clientTrace)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func NewTracerProvider(serviceName string) (*trace.TracerProvider, error) {
	// Port details: https://www.jaegertracing.io/docs/getting-started/
	collectorEndpointURI := "http://localhost:14268/api/traces"

	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(collectorEndpointURI)))
	if err != nil {
		return nil, err
	}

	r := NewResource(serviceName, "v1", "local")
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(r),
		trace.WithSampler(trace.TraceIDRatioBased(1)),
	), nil
}

func NewResource(serviceName string, version string, environment string) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(version),
			attribute.String("environment", environment),
		),
	)
	return r
}
