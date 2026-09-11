package stix

import "time"

// timestampLayout is RFC 3339 in UTC with exactly three fractional
// digits, the precision STIX 2.1 requires for created and modified.
const timestampLayout = "2006-01-02T15:04:05.000Z"

// Timestamp is a STIX 2.1 timestamp. It marshals to JSON in UTC with
// millisecond precision and unmarshals any RFC 3339 timestamp.
type Timestamp struct {
	time.Time
}

// NewTimestamp returns t as a Timestamp.
func NewTimestamp(t time.Time) Timestamp {
	return Timestamp{Time: t.UTC()}
}

// MarshalJSON implements json.Marshaler.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.UTC().Format(timestampLayout) + `"`), nil
}
