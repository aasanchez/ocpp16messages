package test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestCiString20_Valid(t *testing.T) {
	valid := "12345678901234567890"
	cis := core.CiString20Type(valid)
	if err := cis.Validate(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(cis) != valid {
		t.Errorf("expected value %s, got %s", valid, string(cis))
	}
}

func TestCiString20_TooLong(t *testing.T) {
	tooLong := "123456789012345678901" // 21 characters
	cis := core.CiString20Type(tooLong)
	if err := cis.Validate(); err == nil {
		t.Fatal("expected error for string exceeding 20 characters, got nil")
	}
}

func TestCiString20_Empty(t *testing.T) {
	cis := core.CiString20Type("")
	if err := cis.Validate(); err != nil {
		t.Fatal("expected no error for empty string, got", err)
	}
}

func TestCiString20_StringMethod(t *testing.T) {
	s := "hello"
	cis := core.CiString20Type(s)
	if err := cis.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(cis) != s {
		t.Errorf("String() returned %q, expected %q", string(cis), s)
	}
}
