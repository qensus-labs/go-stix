package otel

import (
	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/factory"
)

// mapAddress converts an address value into the correct STIX SCO
// and returns the object ID.
func mapAddress(
	builder *stix.Builder,
	value string,
) (string, error) {

	switch ResolveAddress(value) {

	case IPv4:

		ip, err := factory.IPv4(
			builder,
			value,
		)

		if err != nil {
			return "", err
		}

		return ip.GetID(), nil

	case IPv6:

		ip, err := factory.IPv6(
			builder,
			value,
		)

		if err != nil {
			return "", err
		}

		return ip.GetID(), nil

	case Domain:

		domain, err := factory.DomainName(
			builder,
			value,
		)

		if err != nil {
			return "", err
		}

		return domain.GetID(), nil
	}

	return "", nil
}
