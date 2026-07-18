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

func IPv6(
	builder *stix.Builder,
	value string,
) (*sco.IPv6Addr, error) {

	obj, err := sco.NewIPv6Addr(value)

	if err != nil {
		return nil, err
	}

	builder.Add(obj)

	return obj, nil
}

func MACAddr(
	builder *stix.Builder,
	value string,
) (*sco.MACAddr, error) {

	obj, err := sco.NewMACAddr(value)

	if err != nil {
		return nil, err
	}

	builder.Add(obj)

	return obj, nil
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

func DomainName(
	builder *stix.Builder,
	value string,
) (*sco.DomainName, error) {

	obj, err := sco.NewDomainName(value)
	if err != nil {
		return nil, err
	}

	return builder.Add(obj).(*sco.DomainName), nil
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
