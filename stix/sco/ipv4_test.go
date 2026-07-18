package sco_test

import (
	"testing"

	"github.com/qensus-labs/go-stix/stix/sco"
)

func TestIPv4Creation(t *testing.T) {

	ip, err := sco.NewIPv4Addr(
		"10.0.0.5",
	)

	if err != nil {
		t.Fatal(err)
	}

	if ip.Value != "10.0.0.5" {

		t.Fatalf(
			"unexpected value %s",
			ip.Value,
		)
	}

	if ip.GetType() != "ipv4-addr" {

		t.Fatalf(
			"unexpected type %s",
			ip.GetType(),
		)
	}
}

func TestIPv4Invalid(t *testing.T) {

	_, err := sco.NewIPv4Addr(
		"not-an-ip",
	)

	if err == nil {

		t.Fatal(
			"expected validation error",
		)
	}
}

func TestIPv4DeterministicID(t *testing.T) {

	ip1, _ := sco.NewIPv4Addr(
		"10.0.0.5",
	)

	ip2, _ := sco.NewIPv4Addr(
		"10.0.0.5",
	)

	if ip1.ID != ip2.ID {

		t.Fatal(
			"expected stable STIX ID",
		)
	}
}
