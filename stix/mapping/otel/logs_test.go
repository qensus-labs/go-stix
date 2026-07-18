package otel_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/mapping/otel"
)

func TestMapAttributes(t *testing.T) {

	builder := stix.NewBuilder()

	attributes := map[string]string{

		"client.address": "10.0.0.5",

		"url.full": "https://example.com",
	}

	err := otel.MapAttributes(
		builder,
		attributes,
	)

	if err != nil {
		t.Fatal(err)
	}

	if builder.Count() != 3 {

		t.Fatalf(
			"expected 3 STIX objects, got %d",
			builder.Count(),
		)
	}
}
