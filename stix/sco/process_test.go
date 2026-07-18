package sco_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/sco"
)

func TestNewProcess(t *testing.T) {

	obj, err := sco.NewProcess(
		1234,
		"nginx",
		"nginx -g daemon off;",
	)

	if err != nil {
		t.Fatal(err)
	}

	if obj.Type != stix.TypeProcess {
		t.Fatalf(
			"expected %s got %s",
			stix.TypeProcess,
			obj.Type,
		)
	}

	if obj.PID != 1234 {
		t.Fatalf(
			"unexpected pid %d",
			obj.PID,
		)
	}

	if obj.Name != "nginx" {
		t.Fatalf(
			"unexpected name %s",
			obj.Name,
		)
	}
}

func TestNewProcessEmpty(t *testing.T) {

	_, err := sco.NewProcess(
		0,
		"",
		"",
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
