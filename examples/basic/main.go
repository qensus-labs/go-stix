package main

import (
	"fmt"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/factory"
)

func main() {

	builder := stix.NewBuilder()

	ip, err := factory.IPv4(
		builder,
		"10.0.0.5",
	)

	if err != nil {
		panic(err)
	}

	url, err := factory.URL(
		builder,
		"https://example.com",
	)

	if err != nil {
		panic(err)
	}

	factory.ObservedData(
		builder,
		[]string{
			ip.GetID(),
			url.GetID(),
		},
	)

	data, err := builder.JSON()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
