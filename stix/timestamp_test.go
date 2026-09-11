package stix_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/qensus-labs/go-stix/stix"
	"github.com/qensus-labs/go-stix/stix/sdo"
)

func TestTimestampMarshalsWithMillisecondPrecision(t *testing.T) {

	amsterdam := time.FixedZone("CEST", 2*60*60)

	tests := map[string]struct {
		in   time.Time
		want string
	}{
		"nanoseconds truncated": {
			in:   time.Date(2026, 9, 11, 8, 23, 55, 750120619, time.UTC),
			want: `"2026-09-11T08:23:55.750Z"`,
		},
		"whole seconds padded": {
			in:   time.Date(2026, 9, 11, 8, 23, 55, 0, time.UTC),
			want: `"2026-09-11T08:23:55.000Z"`,
		},
		"converted to UTC": {
			in:   time.Date(2026, 9, 11, 10, 23, 55, 0, amsterdam),
			want: `"2026-09-11T08:23:55.000Z"`,
		},
	}

	for name, tt := range tests {

		t.Run(name, func(t *testing.T) {

			got, err := json.Marshal(stix.Timestamp{Time: tt.in})

			if err != nil {
				t.Fatal(err)
			}

			if string(got) != tt.want {
				t.Fatalf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func TestTimestampUnmarshal(t *testing.T) {

	var got stix.Timestamp

	if err := json.Unmarshal([]byte(`"2026-09-11T08:23:55.750Z"`), &got); err != nil {
		t.Fatal(err)
	}

	want := time.Date(2026, 9, 11, 8, 23, 55, 750000000, time.UTC)

	if !got.Equal(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

// Regression test: nanosecond timestamps broke TAXII servers such as
// Medallion, and STIX 2.1 requires millisecond precision for created
// and modified.
func TestObjectTimestampsUseMillisecondPrecision(t *testing.T) {

	now := time.Date(2026, 9, 11, 8, 23, 55, 750120619, time.UTC)

	data, err := json.Marshal(sdo.NewObservedData(now, now, []string{"url--1"}))

	if err != nil {
		t.Fatal(err)
	}

	var fields map[string]any

	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}

	for _, key := range []string{"created", "modified", "first_observed", "last_observed"} {

		value, _ := fields[key].(string)

		if _, err := time.Parse("2006-01-02T15:04:05.000Z", value); err != nil || len(value) != 24 {
			t.Errorf("%s = %q, want millisecond precision", key, value)
		}
	}
}
