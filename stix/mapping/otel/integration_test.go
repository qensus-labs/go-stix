package otel_test

import (
	"testing"

	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/mapping/otel"
)

func TestCollectorLogRecordToSTIX(t *testing.T) {

	record := plog.NewLogRecord()

	record.Attributes().PutStr(
		otel.ClientAddress,
		"10.0.0.5",
	)

	record.Attributes().PutStr(
		otel.ServerAddress,
		"2001:db8::1",
	)

	record.Attributes().PutStr(
		otel.NetworkPeerAddress,
		"api.example.com",
	)

	record.Attributes().PutStr(
		otel.URLFull,
		"https://example.com",
	)

	record.Attributes().PutStr(
		otel.ProcessName,
		"nginx",
	)

	record.Attributes().PutInt(
		otel.ProcessPID,
		1234,
	)

	obs := otel.FromLogRecord(record)

	if obs.ProcessName != "nginx" {
		t.Fatal("missing process name")
	}

	if obs.ProcessPID != 1234 {
		t.Fatal("missing process pid")
	}

	builder := stix.NewBuilder()

	err := otel.MapObservation(builder, obs)
	if err != nil {
		t.Fatal(err)
	}

	if builder.Count() < 3 {
		t.Fatalf(
			"expected at least 3 STIX objects, got %d",
			builder.Count(),
		)
	}

	foundIPv4 := false
	foundIPv6 := false
	foundDomain := false
	foundURL := false
	foundProcess := false
	foundObservedData := false

	for _, obj := range builder.Objects() {

		switch obj.GetType() {

		case stix.TypeIPv4Addr:
			foundIPv4 = true

		case stix.TypeIPv6Addr:
			foundIPv6 = true

		case stix.TypeDomainName:
			foundDomain = true

		case stix.TypeURL:
			foundURL = true

		case stix.TypeProcess:
			foundProcess = true

		case stix.TypeObservedData:
			foundObservedData = true
		}
	}

	if !foundIPv4 {
		t.Fatal("missing ipv4-addr SCO")
	}

	if !foundIPv6 {
		t.Fatal("missing ipv6-addr SCO")
	}

	if !foundDomain {
		t.Fatal("missing domain-name SCO")
	}

	if !foundURL {
		t.Fatal("missing url SCO")
	}

	if !foundProcess {
		t.Fatal("missing process SCO")
	}

	if !foundObservedData {
		t.Fatal("missing observed-data SDO")
	}
}
