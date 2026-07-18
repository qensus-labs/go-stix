package otel

import (
	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/factory"
	"github.com/qensus-labs/go-stix/stix/observation"
)

// MapLogEvent converts an OpenTelemetry-derived
// security event into STIX objects.
func MapObservation(
	builder *stix.Builder,
	obs observation.Observation,
) error {

	var refs []string

	if obs.ClientAddress != "" {

		ip, err := factory.IPv4(
			builder,
			obs.ClientAddress,
		)

		if err != nil {
			return err
		}

		refs = append(
			refs,
			ip.GetID(),
		)
	}

	if obs.ServerAddress != "" {

		ip, err := factory.IPv4(
			builder,
			obs.ServerAddress,
		)

		if err != nil {
			return err
		}

		refs = append(
			refs,
			ip.GetID(),
		)
	}

	if obs.NetworkPeerAddress != "" {

		ip, err := factory.IPv4(
			builder,
			obs.NetworkPeerAddress,
		)

		if err != nil {
			return err
		}

		refs = append(
			refs,
			ip.GetID(),
		)
	}

	if obs.URL != "" {

		url, err := factory.URL(
			builder,
			obs.URL,
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
