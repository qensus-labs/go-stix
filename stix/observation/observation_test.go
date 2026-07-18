package observation_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix/observation"
)

func TestObservation(t *testing.T) {
	obs := observation.Observation{
		ClientAddress: "10.0.0.5",
		URL:           "https://example.com",
	}

	if obs.ClientAddress != "10.0.0.5" {
		t.Fatal("unexpected client address")
	}

	if obs.URL != "https://example.com" {
		t.Fatal("unexpected URL")
	}
}
