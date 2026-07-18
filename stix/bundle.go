package stix

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Bundle represents a STIX 2.1 Bundle.
//
// A bundle is a container for STIX objects.
type Bundle struct {
	Type    string       `json:"type"`
	ID      string       `json:"id"`
	Objects []STIXObject `json:"objects"`
}

// NewBundle creates an empty STIX bundle.
func NewBundle() *Bundle {

	return &Bundle{
		Type: TypeBundle,
		ID:   "bundle--" + uuid.NewString(),
		Objects: make(
			[]STIXObject,
			0,
		),
	}
}

// Add adds an object to the bundle.
func (b *Bundle) Add(object STIXObject) {

	if object == nil {
		return
	}

	b.Objects = append(
		b.Objects,
		object,
	)
}

// Count returns the number of objects.
func (b *Bundle) Count() int {

	return len(b.Objects)
}

// JSON returns serialized STIX JSON.
func (b *Bundle) JSON() ([]byte, error) {

	return json.MarshalIndent(
		b,
		"",
		"  ",
	)
}
