package otel

import "github.com/qensus-labs/go-stix/stix/observation"

// OpenTelemetry semantic attribute names that are relevant for STIX mapping.
const (
	ClientAddress      = "client.address"
	ServerAddress      = "server.address"
	NetworkPeerAddress = "network.peer.address"
	URLFull            = "url.full"

	HostName          = "host.name"
	UserAgentOriginal = "user_agent.original"

	ProcessName = "process.name"
	ProcessPID  = "process.pid"
)

// Attributes represents a simplified set of OpenTelemetry attributes.
type Attributes map[string]string

// FromAttributes converts OpenTelemetry attributes into a canonical
// Observation model.
func FromAttributes(attrs Attributes) observation.Observation {

	return observation.Observation{
		ClientAddress:      attrs[ClientAddress],
		ServerAddress:      attrs[ServerAddress],
		NetworkPeerAddress: attrs[NetworkPeerAddress],

		URL: attrs[URLFull],

		Hostname: attrs[HostName],

		UserAgent: attrs[UserAgentOriginal],

		ProcessName: attrs[ProcessName],

		// ProcessPID will be added when we support typed values.
	}
}
