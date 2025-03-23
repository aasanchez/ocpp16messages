package test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestCiString20_IsValid(t *testing.T) {
	valid := core.CiString20("shortstring")
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := core.CiString20("this_string_is_way_too_long_to_be_valid")
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString20_MarshalText(t *testing.T) {
	s := core.CiString20("valid")
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != "valid" {
		t.Errorf("Expected 'valid', got: %s", data)
	}

	s = core.CiString20("this_is_more_than_twenty_chars")
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 20 characters")
	}
}

func TestCiString20_UnmarshalText(t *testing.T) {
	var s core.CiString20
	err := s.UnmarshalText([]byte("valid"))
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s != "valid" {
		t.Errorf("Expected 'valid', got: %q", s)
	}

	err = s.UnmarshalText([]byte("this_is_more_than_twenty_chars"))
	if err == nil {
		t.Error("Expected error for string exceeding 20 characters")
	}
}
