package observation

import "time"

// Observation represents a normalized security observation.
//
// It is intentionally vendor-agnostic and can be populated from
// OpenTelemetry, Syslog, Zeek, Suricata, OCSF, ECS or other
// telemetry sources.
type Observation struct {
	Timestamp time.Time

	ClientAddress      string
	ServerAddress      string
	NetworkPeerAddress string

	URL string

	Hostname string

	UserAgent string

	ProcessName        string
	ProcessPID         int64
	ProcessCommandLine string
}
