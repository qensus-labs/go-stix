package stix

import "errors"

var (
	ErrInvalidObject = errors.New("invalid STIX object")

	ErrInvalidIPv4 = errors.New("invalid IPv4 address")

	ErrInvalidURL = errors.New("invalid URL")
)
