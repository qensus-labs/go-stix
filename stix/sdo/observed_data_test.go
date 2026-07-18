package sdo_test

import (
	"testing"
	"time"

	"github.com/qensus-labs/go-stix/stix/sdo"
)

func TestObservedDataCreation(t *testing.T) {

	now := time.Now()

	objectRefs := []string{
		"ipv4-addr--123",
		"url--456",
	}

	observed := sdo.NewObservedData(
		now,
		now,
		objectRefs,
	)

	if observed.GetType() != "observed-data" {

		t.Fatalf(
			"unexpected type %s",
			observed.GetType(),
		)
	}

	if observed.NumberObserved != 2 {

		t.Fatalf(
			"expected 2 objects, got %d",
			observed.NumberObserved,
		)
	}

	if len(observed.ObjectRefs) != 2 {

		t.Fatal(
			"expected object references",
		)
	}
}
