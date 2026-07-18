package stix

import (
	"github.com/google/uuid"
)

var stixNamespace = uuid.MustParse(
	"4e4f7c36-5d6f-4f8e-bb19-9d9f3d9f8c71",
)

// NewDeterministicID creates a stable STIX identifier.

func NewDeterministicID(objectType string, value string) string {

	id := uuid.NewSHA1(
		stixNamespace,
		[]byte(objectType+":"+value),
	)

	return objectType + "--" + id.String()
}

// NewRandomID creates a random STIX identifier.
//
// Used later for SDO objects such as Indicators and Relationships.
func NewRandomID(objectType string) string {

	return objectType + "--" + uuid.NewString()
}
