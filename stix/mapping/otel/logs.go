package otel

import (
	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/factory"
)

// MapLogEvent converts an OpenTelemetry-derived
// security event into STIX objects.
func MapLogEvent(
	builder *stix.Builder,
	event LogEvent,
) error {

	var refs []string

	if event.ClientAddress != "" {

		ip, err := factory.IPv4(
			builder,
			event.ClientAddress,
		)

		if err != nil {
			return err
		}

		refs = append(
			refs,
			ip.GetID(),
		)
	}

	if event.ServerAddress != "" {

		ip, err := factory.IPv4(
			builder,
			event.ServerAddress,
		)

		if err != nil {
			return err
		}

		refs = append(
			refs,
			ip.GetID(),
		)
	}

	if event.NetworkPeerAddress != "" {

		ip, err := factory.IPv4(
			builder,
			event.NetworkPeerAddress,
		)

		if err != nil {
			return err
		}

		refs = append(
			refs,
			ip.GetID(),
		)
	}

	if event.URL != "" {

		url, err := factory.URL(
			builder,
			event.URL,
		)

		if err != nil {
			return err
		}

		refs = append(
			refs,
			url.GetID(),
		)
	}

	if len(refs) > 0 {

		factory.ObservedData(
			builder,
			refs,
		)
	}

	return nil
}
