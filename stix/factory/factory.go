package factory

import (
	"time"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/sco"
	"github.com/qensus-labs/go-stix/stix/sdo"
)

func IPv4(
	builder *stix.Builder,
	value string,
) (*sco.IPv4Addr, error) {

	obj, err := sco.NewIPv4Addr(value)

	if err != nil {
		return nil, err
	}

	return builder.Add(obj).(*sco.IPv4Addr), nil
}

func URL(
	builder *stix.Builder,
	value string,
) (*sco.URL, error) {

	obj, err := sco.NewURL(value)

	if err != nil {
		return nil, err
	}

	return builder.Add(obj).(*sco.URL), nil
}

func ObservedData(
	builder *stix.Builder,
	refs []string,
) *sdo.ObservedData {

	now := time.Now().UTC()

	obj := sdo.NewObservedData(
		now,
		now,
		refs,
	)

	return builder.Add(obj).(*sdo.ObservedData)
}
