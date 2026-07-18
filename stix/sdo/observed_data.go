package sdo

import (
	"time"

	"github.com/qensus-labs/go-stix/stix"
)

// ObservedData represents a STIX 2.1 Observed Data object.
//
// ObservedData describes cyber observables that were seen
// during a specific time period.
type ObservedData struct {
	stix.CommonProperties

	FirstObserved time.Time `json:"first_observed"`

	LastObserved time.Time `json:"last_observed"`

	NumberObserved int `json:"number_observed"`

	ObjectRefs []string `json:"object_refs"`
}

// NewObservedData creates an ObservedData object.

func NewObservedData(
	firstObserved time.Time,
	lastObserved time.Time,
	objectRefs []string,
) *ObservedData {

	return &ObservedData{

		CommonProperties: stix.NewCommonProperties(
			"observed-data",
			stix.NewRandomID(
				"observed-data",
			),
		),

		FirstObserved: firstObserved.UTC(),

		LastObserved: lastObserved.UTC(),

		NumberObserved: len(objectRefs),

		ObjectRefs: objectRefs,
	}
}
