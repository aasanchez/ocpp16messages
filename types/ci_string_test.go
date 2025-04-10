package types

import (
	"strings"
	"testing"
)

func TestCiString20Type_Valid(t *testing.T) {
	t.Parallel()

	CiString20StrValue := strings.Repeat("A", 20)
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

	CiString20StrValue := strings.Repeat("A", 21)

	_, err := CiString20(CiString20StrValue)

	if err == nil {
		t.Errorf("expected error for string too long")
	}
}

// func TestCiStringType_ASCII(t *testing.T) {
// 	t.Parallel()
// 	var builder strings.Builder
// 	for i := 32; i <= 126; i++ {
// 		builder.WriteRune(rune(i))
// 	}
// 	input := builder.String()
// 	s, err := CiString255(input)
// 	if err != nil {
// 		t.Errorf("expected all printable ASCII to pass validation, got %v", err)
// 	}
// 	if s.String() != input {
// 		t.Errorf("unexpected output from String()")
// 	}
// 	if err := s.Validate(); err != nil {
// 		t.Errorf("expected Validate to return nil, got %v", err)
// 	}
// }

func TestCiStringType_ChineseCharacters(t *testing.T) {
	t.Parallel()

	_, err := CiString20("你好，OCPP")

	if err == nil {
		t.Errorf("expected error for non-ASCII input (Chinese), got nil")
	}
}

func TestCiStringType_UkrainianCharacters(t *testing.T) {
	t.Parallel()

	str, err := CiString20("Привіт, OCPP")

	if err == nil {
		t.Errorf("expected error for non-ASCII input (Ukrainian), got %s", str)
	}
}

func TestCiStringType_ArabicCharacters(t *testing.T) {
	t.Parallel()

	str, err := CiString20("مرحبًا OCPP")

	if err == nil {
		t.Errorf("expected error for non-ASCII input (Arabic), got %s", str)
	}
}

func TestCiStringType_NonPrintable(t *testing.T) {
	t.Parallel()

	str, err := CiString20("Valid\x01String")

	if err == nil {
		t.Errorf("expected error for non-printable ASCII, got %s", str)
	}
}

func TestCiStringType_Empty(t *testing.T) {
	t.Parallel()

	str, err := CiString20("")

	if err != nil {
		t.Errorf("unexpected error for empty string(%s): %v", str, err)
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
