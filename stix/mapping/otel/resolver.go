package otel

import "net"

type AddressType int

const (
	Unknown AddressType = iota
	IPv4
	IPv6
	Domain
)

func ResolveAddress(value string) AddressType {

	ip := net.ParseIP(value)

	if ip != nil {

		if ip.To4() != nil {
			return IPv4
		}

		return IPv6
	}

	return Domain
}
