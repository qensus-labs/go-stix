package otel_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix/mapping/otel"
)

func TestLogEventModel(t *testing.T) {

	event := otel.LogEvent{
		ClientAddress: "10.0.0.5",
		URL:           "https://example.com",
	}

	if event.ClientAddress == "" {
		t.Fatal("expected client address")
	}

	if event.URL == "" {
		t.Fatal("expected URL")
	}
}
