package sco

import (
	"errors"
	"net"

	"github.com/qensus-labs/go-stix/stix"
)

// IPv6Addr represents a STIX IPv6 address SCO.
type IPv6Addr struct {
	stix.CommonProperties

	Value string `json:"value"`
}

// NewIPv6Addr creates a new IPv6 address SCO.
func NewIPv6Addr(value string) (*IPv6Addr, error) {

	ip := net.ParseIP(value)

	if ip == nil || ip.To4() != nil {
		return nil, errors.New("invalid IPv6 address")
	}

	return &IPv6Addr{
		CommonProperties: stix.NewCommonProperties(
			stix.TypeIPv6Addr,
			stix.NewDeterministicID(
				stix.TypeIPv6Addr,
				value,
			),
		),
		Value: value,
	}, nil
}
