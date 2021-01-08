package telemetry

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

func init() {
	exporterName := os.Getenv("OTEL_EXPORTER")
	if exporterName == "otlp" {
		ctx := context.Background()
		exporter, err := otlp.NewExporter(ctx, otlp.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		cfg := trace.Config{
			DefaultSampler: trace.AlwaysSample(),
		}
		tp := trace.NewTracerProvider(
			trace.WithConfig(cfg),
			trace.WithBatcher(exporter),
		)
		if err != nil {
			log.Fatal(err)
		}
		otel.SetTracerProvider(tp)
	} else {
		exporter, err := stdout.NewExporter(stdout.WithPrettyPrint())
		if err != nil {
			log.Fatal(err)
		}
		cfg := trace.Config{
			DefaultSampler: trace.AlwaysSample(),
		}
		tp := trace.NewTracerProvider(
			trace.WithConfig(cfg),
			trace.WithSyncer(exporter),
		)
		if err != nil {
			log.Fatal(err)
		}
		otel.SetTracerProvider(tp)
	}
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}