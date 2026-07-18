package sco

import (
	"errors"
	"strings"

	"github.com/qensus-labs/go-stix/stix"
)

// DomainName represents the STIX 2.1 domain-name SCO.
//
// Specification:
// https://docs.oasis-open.org/cti/stix/v2.1/stix-v2.1.html#_Toc31107557
type DomainName struct {
	stix.CommonProperties

	Value string `json:"value"`
}

// NewDomainName creates a new STIX domain-name object.
func NewDomainName(value string) (*DomainName, error) {

	value = strings.TrimSpace(value)

	if value == "" {
		return nil, errors.New("domain-name value is required")
	}

	return &DomainName{

		CommonProperties: stix.NewCommonProperties(
			stix.TypeDomainName,
			stix.NewDeterministicID(
				stix.TypeDomainName,
				value,
			),
		),

		Value: value,
	}, nil

}
