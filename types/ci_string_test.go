package types

import (
	"strings"
	"testing"
)

func TestCiString20Type_Valid(t *testing.T) {
	t.Parallel()
	cs,
		err := CiString20("ValidString")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if cs.String() != "ValidString" {
		t.Errorf("expected string output, got %s", cs.String())
	}
	if err := cs.Validate(); err != nil {
		t.Errorf("expected Validate to return nil, got %v", err)
	}
}

func TestCiString20Type_TooLong(t *testing.T) {
	t.Parallel()
	_, err := CiString20("ThisStringIsWayTooLongFor20Chars")
	if err == nil {
		t.Errorf("expected error for string too long")
	}
}

func TestCiString25Type_NonPrintable(t *testing.T) {
	t.Parallel()
	_, err := CiString25("Valid\x01String")
	if err == nil {
		t.Errorf("expected error for non-printable ASCII")
	}
}

func TestCiString25Type_Valid(t *testing.T) {
	t.Parallel()
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
}

func TestCiString50Type_Boundary(t *testing.T) {
	t.Parallel()
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
}

func TestCiString255Type_Validate(t *testing.T) {
	t.Parallel()
	s, _ := CiString255("abc")
	if s.String() != "abc" {
		t.Errorf("unexpected string output")
	}
	if err := s.Validate(); err != nil {
		t.Errorf("expected Validate to return nil, got %v", err)
	}
}

func TestCiString500Type_Valid(t *testing.T) {
	t.Parallel()
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
}

func TestCiString500Type_TooLong(t *testing.T) {
	t.Parallel()
	value := strings.Repeat("X", 501)
	_, err := CiString500(value)
	if err == nil {
		t.Errorf("expected error for value exceeding max length")
	}
}

func TestCiString255Type_ASCII(t *testing.T) {
	t.Parallel()
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
}
