package otel

import (
	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/factory"
)

// Attribute names mapped from OpenTelemetry semantic attributes.
const (
	ClientAddress      = "client.address"
	ServerAddress      = "server.address"
	NetworkPeerAddress = "network.peer.address"
	URLFull            = "url.full"
)

// MapAttributes converts OpenTelemetry-style attributes
// into STIX objects.
func MapAttributes(
	builder *stix.Builder,
	attributes map[string]string,
) error {

	var refs []string

	for key, value := range attributes {

		switch key {

		case ClientAddress,
			ServerAddress,
			NetworkPeerAddress:

			ip, err := factory.IPv4(
				builder,
				value,
			)

			if err != nil {
				return err
			}

			refs = append(
				refs,
				ip.GetID(),
			)

		case URLFull:

			url, err := factory.URL(
				builder,
				value,
			)

			if err != nil {
				return err
			}

			refs = append(
				refs,
				url.GetID(),
			)
		}
	}

	if len(refs) > 0 {

		factory.ObservedData(
			builder,
			refs,
		)
	}

	return nil
}
