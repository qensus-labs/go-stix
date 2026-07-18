package otel

import (
	"strconv"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/qensus-labs/go-stix/stix/observation"
)

// FromLogRecord converts a Collector LogRecord into a canonical Observation.
func FromLogRecord(record plog.LogRecord) observation.Observation {

	obs := observation.Observation{}

	record.Attributes().Range(func(k string, v pcommon.Value) bool {

		switch k {

		case ClientAddress:
			obs.ClientAddress = v.Str()

		case ServerAddress:
			obs.ServerAddress = v.Str()

		case NetworkPeerAddress:
			obs.NetworkPeerAddress = v.Str()

		case URLFull:
			obs.URL = v.Str()

		case HostName:
			obs.Hostname = v.Str()

		case UserAgentOriginal:
			obs.UserAgent = v.Str()

		case ProcessName:
			obs.ProcessName = v.Str()

		case ProcessPID:
			switch v.Type() {
			case pcommon.ValueTypeInt:
				obs.ProcessPID = v.Int()
			case pcommon.ValueTypeStr:
				if pid, err := strconv.ParseInt(v.Str(), 10, 64); err == nil {
					obs.ProcessPID = pid
				}
			}

		case ProcessCommandLine:
			obs.ProcessCommandLine = v.Str()

		}

		return true
	})

	if ts := record.Timestamp(); ts != 0 {
		obs.Timestamp = ts.AsTime()
	}

	return obs
}
