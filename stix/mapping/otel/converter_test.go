package otel_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/mapping/otel"
	"github.com/qensus-labs/go-stix/stix/observation"
)

func TestMapObservation(t *testing.T) {

	builder := stix.NewBuilder()

	obs := observation.Observation{
		ClientAddress: "10.0.0.5",
		URL:           "https://example.com",
	}

	err := otel.MapObservation(
		builder,
		obs,
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
