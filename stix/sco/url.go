package sco

import (
	"net/url"

	"github.com/qensus-labs/go-stix/stix"
)

// URL represents a STIX URL Cyber Observable Object.
type URL struct {
	stix.CommonProperties

	Value string `json:"value"`
}

// NewURL creates a STIX URL object.

func NewURL(value string) (*URL, error) {

	parsed, err := url.Parse(value)

	if err != nil {
		return nil, stix.ErrInvalidURL
	}

	if parsed.Scheme == "" ||
		parsed.Host == "" {

		return nil, stix.ErrInvalidURL
	}

	return &URL{

		CommonProperties: stix.NewCommonProperties(
			stix.TypeURL,
			stix.NewDeterministicID(
				stix.TypeURL,
				value,
			),
		),

		Value: value,
	}, nil
}
