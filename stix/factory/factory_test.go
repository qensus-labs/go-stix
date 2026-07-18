package factory_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/factory"
)

func TestFactoryCreatesObjects(t *testing.T) {

	builder := stix.NewBuilder()

	ip, err := factory.IPv4(
		builder,
		"10.0.0.5",
	)

	if err != nil {
		t.Fatal(err)
	}

	url, err := factory.URL(
		builder,
		"https://example.com",
	)

	if err != nil {
		t.Fatal(err)
	}

	factory.ObservedData(
		builder,
		[]string{
			ip.GetID(),
			url.GetID(),
		},
	)

	if builder.Count() != 3 {
		t.Fatalf(
			"expected 3 objects, got %d",
			builder.Count(),
		)
	}
}
