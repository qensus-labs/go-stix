package stix_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/sdo"
)

func TestTimestampJSONUsesWholeSecondsUTC(t *testing.T) {
	input := time.Date(
		2026, time.October, 18,
		15, 17, 57, 120527490,
		time.FixedZone("test-zone", 2*60*60),
	)

	got, err := json.Marshal(stix.Timestamp(input))
	if err != nil {
		t.Fatalf("marshal timestamp: %v", err)
	}

	const want = `"2026-10-18T13:17:57Z"`
	if string(got) != want {
		t.Fatalf("unexpected timestamp: got %s, want %s", got, want)
	}
}

func TestSTIXObjectTimestampsUseWholeSeconds(t *testing.T) {
	input := time.Date(
		2026, time.October, 18,
		15, 17, 57, 120527490,
		time.UTC,
	)
	obj := sdo.NewObservedData(input, input, []string{"domain-name--example"})

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("marshal observed-data: %v", err)
	}

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatalf("unmarshal JSON: %v", err)
	}

	for _, field := range []string{
		"created",
		"modified",
		"first_observed",
		"last_observed",
	} {
		var value string
		if err := json.Unmarshal(fields[field], &value); err != nil {
			t.Fatalf("decode %s: %v", field, err)
		}

		parsed, err := time.Parse("2006-01-02T15:04:05Z", value)
		if err != nil {
			t.Errorf("%s is not in whole-second UTC format: %q: %v", field, value, err)
			continue
		}

		if parsed.Format("2006-01-02T15:04:05Z") != value {
			t.Errorf("%s has unexpected format: %q", field, value)
		}
	}

	if got := string(fields["first_observed"]); got != `"2026-10-18T15:17:57Z"` {
		t.Errorf("unexpected first_observed: %s", got)
	}

	if got := string(fields["last_observed"]); got != `"2026-10-18T15:17:57Z"` {
		t.Errorf("unexpected last_observed: %s", got)
	}
}
