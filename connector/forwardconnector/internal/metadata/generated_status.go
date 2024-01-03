// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

const (
	Type                      = "forward"
	TracesToTracesStability   = component.StabilityLevelBeta
	MetricsToMetricsStability = component.StabilityLevelBeta
	LogsToLogsStability       = component.StabilityLevelBeta
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("otelcol/forward")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("otelcol/forward")
}