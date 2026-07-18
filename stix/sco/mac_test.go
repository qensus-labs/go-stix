package sco_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/sco"
)

func TestNewMACAddr(t *testing.T) {

	obj, err := sco.NewMACAddr(
		"AA:BB:CC:DD:EE:FF",
	)

	if err != nil {
		t.Fatal(err)
	}

	if obj.Type != stix.TypeMACAddr {
		t.Fatalf(
			"expected %s got %s",
			stix.TypeMACAddr,
			obj.Type,
		)
	}

	if obj.Value != "aa:bb:cc:dd:ee:ff" {
		t.Fatalf(
			"unexpected value %s",
			obj.Value,
		)
	}
}

func TestNewMACAddrInvalid(t *testing.T) {

	_, err := sco.NewMACAddr(
		"invalid",
	)

	if err == nil {
		t.Fatal("expected error")
	}
}
