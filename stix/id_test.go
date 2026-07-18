package stix_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix"
)

func TestDeterministicID(t *testing.T) {

	id1 := stix.NewDeterministicID(
		"ipv4-addr",
		"10.0.0.5",
	)

	id2 := stix.NewDeterministicID(
		"ipv4-addr",
		"10.0.0.5",
	)

	if id1 != id2 {
		t.Fatalf(
			"expected deterministic IDs to match",
		)
	}
}

func TestRandomID(t *testing.T) {

	id1 := stix.NewRandomID(
		"indicator",
	)

	id2 := stix.NewRandomID(
		"indicator",
	)

	if id1 == id2 {
		t.Fatalf(
			"expected random IDs to differ",
		)
	}
}
