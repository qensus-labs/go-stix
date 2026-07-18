package otel

import "net"

type ObservableType int

const (
	ObservableUnknown ObservableType = iota
	ObservableIPv4
	ObservableIPv6
	ObservableDomain
	ObservableMAC
)

func DetectObservable(value string) ObservableType {

	if mac, err := net.ParseMAC(value); err == nil {
		if len(mac) == 6 {
			return ObservableMAC
		}
	}

	ip := net.ParseIP(value)

	if ip != nil {
		if ip.To4() != nil {
			return ObservableIPv4
		}

		return ObservableIPv6
	}

	if value != "" {
		return ObservableDomain
	}

	return ObservableUnknown
}
