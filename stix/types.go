package stix

import "time"

// STIXObject represents any STIX object.
type STIXObject interface {
	GetID() string
	GetType() string
}

// CommonProperties contains fields shared by STIX objects.
type CommonProperties struct {
	Type        string    `json:"type"`
	SpecVersion string    `json:"spec_version"`
	ID          string    `json:"id"`
	Created     time.Time `json:"created,omitempty"`
	Modified    time.Time `json:"modified,omitempty"`
}

// GetID returns the STIX identifier.
func (c CommonProperties) GetID() string {
	return c.ID
}

// GetType returns the STIX object type.
func (c CommonProperties) GetType() string {
	return c.Type
}

// NewCommonProperties creates common STIX metadata.
func NewCommonProperties(objectType string, id string) CommonProperties {

	now := time.Now().UTC()

	return CommonProperties{
		Type:        objectType,
		SpecVersion: STIXVersion,
		ID:          id,
		Created:     now,
		Modified:    now,
	}
}
