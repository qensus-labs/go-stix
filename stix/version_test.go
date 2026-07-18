package stix_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix"
)

func TestVersion(t *testing.T) {

	if stix.Version == "" {
		t.Fatal("expected version to be defined")
	}
}
