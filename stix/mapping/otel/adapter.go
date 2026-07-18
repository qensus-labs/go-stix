package otel

// Attributes represents OpenTelemetry attributes
// relevant for STIX conversion.
type Attributes map[string]string

// FromAttributes converts telemetry attributes
// into the internal LogEvent model.
func FromAttributes(
	attributes Attributes,
) LogEvent {

	event := LogEvent{}

	event.ClientAddress =
		attributes[ClientAddress]

	event.ServerAddress =
		attributes[ServerAddress]

	event.NetworkPeerAddress =
		attributes[NetworkPeerAddress]

	event.URL =
		attributes[URLFull]

	return event
}
