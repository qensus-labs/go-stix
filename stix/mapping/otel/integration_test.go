package otel_test

import (
	"testing"

	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/mapping/otel"
)

func TestCollectorLogRecordToSTIX(t *testing.T) {

	record := plog.NewLogRecord()

	record.Attributes().PutStr(
		otel.ClientAddress,
		"10.0.0.5",
	)

	record.Attributes().PutStr(
		otel.URLFull,
		"https://example.com",
	)

	obs := otel.FromLogRecord(record)

	builder := stix.NewBuilder()

	err := otel.MapObservation(builder, obs)
	if err != nil {
		t.Fatal(err)
	}

	if builder.Count() != 3 {
		t.Fatalf("expected 3 STIX objects, got %d", builder.Count())
	}
}
