package sco

import (
	"net"

	"github.com/qensus-labs/go-stix/stix"
)

// IPv4Addr represents a STIX IPv4 Address Cyber Observable Object.
type IPv4Addr struct {
	stix.CommonProperties

	Value string `json:"value"`
}

// NewIPv4Addr creates a new IPv4 SCO.

func NewIPv4Addr(value string) (*IPv4Addr, error) {

	ip := net.ParseIP(value)

	if ip == nil || ip.To4() == nil {
		return nil, stix.ErrInvalidIPv4
	}

	normalized := ip.To4().String()

	return &IPv4Addr{
		CommonProperties: stix.NewCommonProperties(
			stix.TypeIPv4Addr,
			stix.NewDeterministicID(
				stix.TypeIPv4Addr,
				normalized,
			),
		),
		Value: normalized,
	}, nil
}
