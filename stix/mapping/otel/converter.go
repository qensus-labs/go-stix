package otel

import (
	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/factory"
	"github.com/qensus-labs/go-stix/stix/observation"
)

// MapObservation converts an OpenTelemetry-derived
// security observation into STIX objects.
func MapObservation(
	builder *stix.Builder,
	obs observation.Observation,
) error {

	var refs []string

	addresses := []string{
		obs.ClientAddress,
		obs.ServerAddress,
		obs.NetworkPeerAddress,
	}

	for _, address := range addresses {

		if address == "" {
			continue
		}

		id, err := mapAddress(
			builder,
			address,
		)

		if err != nil {
			return err
		}

		if id != "" {
			refs = append(refs, id)
		}
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

	if obs.ProcessName != "" || obs.ProcessPID != 0 {

		process, err := factory.Process(
			builder,
			obs.ProcessPID,
			obs.ProcessName,
			obs.ProcessCommandLine,
		)

		if err != nil {
			return err
		}

		refs = append(
			refs,
			process.GetID(),
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
