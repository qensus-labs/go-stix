package stix

import (
	"encoding/json"
	"time"
)

// Timestamp is a STIX timestamp serialized in UTC with whole-second precision.
//
// Go's time.Time JSON marshaler emits fractional seconds when present. Some TAXII
// servers, including the Medallion deployment this library targets, require the
// timestamp format YYYY-MM-DDTHH:MM:SSZ. Timestamp therefore truncates
// fractional seconds during JSON serialization.
type Timestamp time.Time

// MarshalJSON serializes the timestamp as UTC with no fractional seconds.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	value := time.Time(t).UTC().Truncate(time.Second)
	return json.Marshal(value.Format("2006-01-02T15:04:05Z"))
}

// UnmarshalJSON accepts RFC 3339 timestamps, including fractional seconds.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return err
	}

	*t = Timestamp(parsed.UTC())
	return nil
}
