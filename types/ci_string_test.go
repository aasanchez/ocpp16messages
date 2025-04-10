package types

import (
	"strings"
	"testing"
)

func TestCiString20Type_Valid(t *testing.T) {
	t.Parallel()
	CiString20StrValue := "ValidString"
	CiString20Str, err := CiString20(CiString20StrValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if CiString20Str.String() != CiString20StrValue {
		t.Errorf("expected string output, got %s", CiString20Str.String())
	}
	err = CiString20Str.Validate()
	if err != nil {
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

func TestCiString25Type_Valid(t *testing.T) {
	t.Parallel()
	CiString25StrValue := "ValidString"
	CiString25Str, err := CiString25(CiString25StrValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if CiString25Str.String() != CiString25StrValue {
		t.Errorf("expected string output, got %s", CiString25Str.String())
	}
	err = CiString25Str.Validate()
	if err != nil {
		t.Errorf("expected Validate to return nil, got %v", err)
	}
}

func TestCiString25Type_NonPrintable(t *testing.T) {
	t.Parallel()
	_, err := CiString25("Valid\x01String")
	if err == nil {
		t.Errorf("expected error for non-printable ASCII")
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
	err = cs.Validate()
	if err != nil {
		t.Errorf("expected Validate to return nil, got %v", err)
	}
}

func TestCiString255Type_Validate(t *testing.T) {
	t.Parallel()
	cs, err := CiString25("abc")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if cs.String() != "abc" {
		t.Errorf("expected string output, got %s", cs.String())
	}
	err = cs.Validate()
	if err != nil {
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
	err = s.Validate()
	if err != nil {
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

func TestCiStringType_ASCII(t *testing.T) {
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

func TestCiStringType_ChineseCharacters(t *testing.T) {
	t.Parallel()
	_, err := CiString20("你好，OCPP")
	if err == nil {
		t.Errorf("expected error for non-ASCII input (Chinese), got nil")
	}
}

func TestCiStringType_UkrainianCharacters(t *testing.T) {
	t.Parallel()
	_, err := CiString20("Привіт, OCPP")
	if err == nil {
		t.Errorf("expected error for non-ASCII input (Ukrainian), got nil")
	}
}

func TestCiStringType_ArabicCharacters(t *testing.T) {
	t.Parallel()
	_, err := CiString20("مرحبًا OCPP")
	if err == nil {
		t.Errorf("expected error for non-ASCII input (Arabic), got nil")
	}
}

func TestCiStringType_NonPrintable(t *testing.T) {
	t.Parallel()
	_, err := CiString20("Valid\x01String")
	if err == nil {
		t.Errorf("expected error for non-printable ASCII")
	}
}

func TestCiStringType_Empty(t *testing.T) {
	t.Parallel()
	_, err := CiString20("")
	if err != nil {
		t.Errorf("unexpected error for empty string: %v", err)
	}
}

func TestCiStringType_NullEquivalent(t *testing.T) {
	t.Parallel()

	var s string // Go does not allow null strings, but this simulates the zero-value ("null")
	_, err := CiString20(s)
	if err != nil {
		t.Errorf("unexpected error for zero-value string: %v", err)
	}
}
