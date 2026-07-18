package otel

// LogEvent represents the security-relevant fields
// extracted from an OpenTelemetry log record.
//
// This model intentionally avoids a dependency on the
// OpenTelemetry SDK. It keeps the mapper lightweight and
// testable.
type LogEvent struct {
	ClientAddress string

	ServerAddress string

	NetworkPeerAddress string

	URL string
}
