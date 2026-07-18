package otel_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix/mapping/otel"
)

func TestFromAttributes(t *testing.T) {

	attrs := otel.Attributes{
		"client.address": "10.0.0.5",
		"url.full":       "https://example.com",
	}

	event := otel.FromAttributes(attrs)

	if event.ClientAddress != "10.0.0.5" {
		t.Fatal("missing client address")
	}
}
