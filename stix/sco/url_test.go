package sco_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix/sco"
)

func TestURLCreation(t *testing.T) {

	u, err := sco.NewURL(
		"https://example.com/path",
	)

	if err != nil {

		t.Fatal(err)
	}

	if u.Value != "https://example.com/path" {

		t.Fatalf(
			"unexpected URL %s",
			u.Value,
		)
	}

	if u.GetType() != "url" {

		t.Fatalf(
			"unexpected type %s",
			u.GetType(),
		)
	}
}

func TestInvalidURL(t *testing.T) {

	_, err := sco.NewURL(
		"not-a-url",
	)

	if err == nil {

		t.Fatal(
			"expected URL validation error",
		)
	}
}
