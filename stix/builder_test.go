package stix_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix"
)

type testObject struct {
	stix.CommonProperties

	Value string `json:"value"`
}

func newTestObject(
	id string,
) *testObject {

	return &testObject{

		CommonProperties: stix.NewCommonProperties(
			"test-object",
			id,
		),

		Value: "test",
	}
}

func TestBuilderDeduplication(
	t *testing.T,
) {

	builder := stix.NewBuilder()

	object1 := newTestObject(
		"test--123",
	)

	object2 := newTestObject(
		"test--123",
	)

	builder.Add(object1)
	builder.Add(object2)

	if builder.Count() != 1 {

		t.Fatalf(
			"expected one object, got %d",
			builder.Count(),
		)
	}
}

func TestBundleJSON(
	t *testing.T,
) {

	builder := stix.NewBuilder()

	builder.Add(
		newTestObject(
			"test--123",
		),
	)

	json, err := builder.JSON()

	if err != nil {

		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if len(json) == 0 {

		t.Fatal(
			"expected JSON output",
		)
	}
}
