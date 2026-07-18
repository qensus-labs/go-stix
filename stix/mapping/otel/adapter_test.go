package otel_test

import (
	"testing"

	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/qensus-labs/go-stix/stix/mapping/otel"
)

func TestFromLogRecord(t *testing.T) {

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

	if obs.ClientAddress != "10.0.0.5" {
		t.Fatal("unexpected client address")
	}

	if obs.URL != "https://example.com" {
		t.Fatal("unexpected URL")
	}
}
