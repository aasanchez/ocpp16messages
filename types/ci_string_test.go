package types

import (
	"strings"
	"testing"
)

func TestCiStringTypes_Validation(t *testing.T) {
	t.Run("valid CiString20Type", func(t *testing.T) {
		cs, err := CiString20("ValidString")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if cs.String() != "ValidString" {
			t.Errorf("expected string output, got %s", cs.String())
		}
		if err := cs.Validate(); err != nil {
			t.Errorf("expected Validate to return nil, got %v", err)
		}
	})

	t.Run("too long CiString20Type", func(t *testing.T) {
		_, err := CiString20("ThisStringIsWayTooLongFor20Chars")
		if err == nil {
			t.Errorf("expected error for string too long")
		}
	})

	t.Run("non-printable ASCII in CiString25Type", func(t *testing.T) {
		_, err := CiString25("Valid\x01String")
		if err == nil {
			t.Errorf("expected error for non-printable ASCII")
		}
	})

	t.Run("valid CiString25Type", func(t *testing.T) {
		cs, err := CiString25("AnotherValid")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if cs.String() != "AnotherValid" {
			t.Errorf("expected string output, got %s", cs.String())
		}
		if err := cs.Validate(); err != nil {
			t.Errorf("expected Validate to return nil, got %v", err)
		}
	})

	t.Run("boundary max length CiString50Type", func(t *testing.T) {
		value := strings.Repeat("A", 50)
		cs, err := CiString50(value)
		if err != nil {
			t.Errorf("expected no error for exact max length, got %v", err)
		}
		if cs.String() != value {
			t.Errorf("unexpected string output")
		}
		if err := cs.Validate(); err != nil {
			t.Errorf("expected Validate to return nil, got %v", err)
		}
	})

	t.Run("validate method check on CiString255Type", func(t *testing.T) {
		s, _ := CiString255("abc")
		if s.String() != "abc" {
			t.Errorf("unexpected string output")
		}
		if err := s.Validate(); err != nil {
			t.Errorf("expected Validate to return nil, got %v", err)
		}
	})

	t.Run("construct CiString500Type successfully", func(t *testing.T) {
		s, err := CiString500("This is a test")
		if err != nil {
			t.Errorf("expected valid CiString500Type, got %v", err)
		}
		if s.String() != "This is a test" {
			t.Errorf("unexpected string output")
		}
		if err := s.Validate(); err != nil {
			t.Errorf("expected Validate to return nil, got %v", err)
		}
	})

	t.Run("invalid CiString500Type length", func(t *testing.T) {
		value := strings.Repeat("X", 501)
		_, err := CiString500(value)
		if err == nil {
			t.Errorf("expected error for value exceeding max length")
		}
	})

	t.Run("printable ASCII full loop", func(t *testing.T) {
		var builder strings.Builder
		for i := 32; i <= 126; i++ {
			builder.WriteRune(rune(i))
		}
		input := builder.String()
		s, err := CiString255(input)
		if err != nil {
			t.Errorf("expected all printable ASCII to pass validation, got %v", err)
		}
		if s.String() != input {
			t.Errorf("unexpected output from String()")
		}
		if err := s.Validate(); err != nil {
			t.Errorf("expected Validate to return nil, got %v", err)
		}
	})
}
