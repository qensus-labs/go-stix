package sco_test

import (
	"encoding/json"
	"testing"

	"github.com/qensus-labs/go-stix/stix/sco"
)

func TestNewDomainName(t *testing.T) {

	obj, err := sco.NewDomainName("example.com")
	if err != nil {
		t.Fatal(err)
	}

	if obj.Type != "domain-name" {
		t.Fatalf("expected type domain-name, got %q", obj.Type)
	}

	if obj.Value != "example.com" {
		t.Fatalf("expected value example.com, got %q", obj.Value)
	}

	if obj.ID == "" {
		t.Fatal("expected ID to be set")
	}
}

func TestNewDomainNameTrimSpace(t *testing.T) {

	obj, err := sco.NewDomainName("  example.com  ")
	if err != nil {
		t.Fatal(err)
	}

	if obj.Value != "example.com" {
		t.Fatalf("expected trimmed value, got %q", obj.Value)
	}
}

func TestDomainNameJSON(t *testing.T) {

	obj, err := sco.NewDomainName("example.com")
	if err != nil {
		t.Fatal(err)
	}

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}

	var decoded map[string]any

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}

	if decoded["type"] != "domain-name" {
		t.Fatalf("unexpected type %v", decoded["type"])
	}

	if decoded["value"] != "example.com" {
		t.Fatalf("unexpected value %v", decoded["value"])
	}
}
