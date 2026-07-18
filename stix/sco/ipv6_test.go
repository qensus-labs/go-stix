package sco_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/sco"
)

func TestNewIPv6Addr(t *testing.T) {

	obj, err := sco.NewIPv6Addr(
		"2001:db8::1",
	)

	if err != nil {
		t.Fatal(err)
	}

	if obj.Type != stix.TypeIPv6Addr {
		t.Fatalf(
			"expected %s got %s",
			stix.TypeIPv6Addr,
			obj.Type,
		)
	}

	if obj.Value != "2001:db8::1" {
		t.Fatalf(
			"unexpected value %s",
			obj.Value,
		)
	}

	if obj.ID == "" {
		t.Fatal("expected deterministic ID")
	}
}

func TestNewIPv6AddrRejectIPv4(t *testing.T) {

	_, err := sco.NewIPv6Addr(
		"192.168.1.10",
	)

	if err == nil {
		t.Fatal(
			"expected IPv4 address to be rejected",
		)
	}
}
