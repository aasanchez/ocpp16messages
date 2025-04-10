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
	CiString25StrValue := "Ab1Cd2Ef3Gh4Ij5Kl6Mn7Op8Q"
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

func TestCiString25Type_TooLong(t *testing.T) {
	t.Parallel()
	_, err := CiString25("Ab1Cd2Ef3Gh4Ij5Kl6Mn7Op8Qr9St")
	if err == nil {
		t.Errorf("expected error for string too long")
	}
}

func TestCiString50Type_Valid(t *testing.T) {
	t.Parallel()
	CiString50StrValue := "Ab1Cd2Ef3Gh4Ij5Kl6Mn7Op8QAb1Cd2Ef3Gh4Ij5Kl6Mn7Op8Q"
	CiString50Str, err := CiString50(CiString50StrValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if CiString50Str.String() != CiString50StrValue {
		t.Errorf("expected string output, got %s", CiString50Str.String())
	}
	err = CiString50Str.Validate()
	if err != nil {
		t.Errorf("expected Validate to return nil, got %v", err)
	}
}

func TestCiString50Type_TooLong(t *testing.T) {
	t.Parallel()
	_, err := CiString50("Ab1Cd2Ef3Gh4Ij5Kl6Mn7Op8QAb1Cd2Ef3Gh4Ij5Kl6Mn7Op8Q1234")
	if err == nil {
		t.Errorf("expected error for string too long")
	}
}

func TestCiString255Type_Valid(t *testing.T) {
	t.Parallel()
	CiString255StrValue := strings.Repeat("A", 255)
	CiString255Str, err := CiString255(CiString255StrValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if CiString255Str.String() != CiString255StrValue {
		t.Errorf("expected string output, got %s", CiString255Str.String())
	}
	err = CiString255Str.Validate()
	if err != nil {
		t.Errorf("expected Validate to return nil, got %v", err)
	}
}

func TestCiString255Type_TooLong(t *testing.T) {
	t.Parallel()
	CiString255StrValue := strings.Repeat("A", 256)
	_, err := CiString255(CiString255StrValue)
	if err == nil {
		t.Errorf("expected error for string too long")
	}
}

func TestCiString500Type_Valid(t *testing.T) {
	t.Parallel()
	CiString500StrValue := strings.Repeat("A", 255)
	CiString500Str, err := CiString500(CiString500StrValue)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if CiString500Str.String() != CiString500StrValue {
		t.Errorf("expected string output, got %s", CiString500Str.String())
	}
	err = CiString500Str.Validate()
	if err != nil {
		t.Errorf("expected Validate to return nil, got %v", err)
	}
}

func TestCiString500Type_TooLong(t *testing.T) {
	t.Parallel()
	CiString500StrValue := strings.Repeat("A", 501)
	_, err := CiString500(CiString500StrValue)
	if err == nil {
		t.Errorf("expected error for string too long")
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
	_, err := CiString20("你好, OCPP")
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
	_, err := CiString255("Valid\x01String")
	if err == nil {
		t.Errorf("expected error for non-printable ASCII")
	}
}

func TestCiStringType_Empty(t *testing.T) {
	t.Parallel()
	_, err := CiString25("")
	if err != nil {
		t.Errorf("unexpected error for empty string: %v", err)
	}
}

func TestCiStringType_NullEquivalent(t *testing.T) {
	t.Parallel()
	var nullString string // Go does not allow null strings, but this simulates the zero-value ("null")
	_, err := CiString20(nullString)
	if err != nil {
		t.Errorf("unexpected error for zero-value string: %v", err)
	}
}
