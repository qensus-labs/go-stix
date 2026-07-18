package sco

import (
	"errors"
	"net"
	"strings"

	"github.com/qensus-labs/go-stix/stix"
)

// MACAddr represents a STIX mac-addr SCO.
type MACAddr struct {
	stix.CommonProperties

	Value string `json:"value"`
}

// NewMACAddr creates a new STIX MAC address SCO.
func NewMACAddr(value string) (*MACAddr, error) {

	value = strings.TrimSpace(value)

	mac, err := net.ParseMAC(value)

	if err != nil {
		return nil, errors.New("invalid MAC address")
	}

	// Normalize representation.
	value = mac.String()

	return &MACAddr{
		CommonProperties: stix.NewCommonProperties(
			stix.TypeMACAddr,
			stix.NewDeterministicID(
				stix.TypeMACAddr,
				value,
			),
		),
		Value: value,
	}, nil
}
