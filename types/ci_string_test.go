package types

import (
	"strings"
	"testing"
)

// Define constants for repeated error messages.
const (
	errExpectedNoError      = "expected no error, got %v"
	errExpectedStringOutput = "expected string output, got %s"
	errExpectedErrorTooLong = "expected error for string too long"
)

func TestCiString20TypeValid(t *testing.T) {
	t.Parallel()

	CiString20StrValue := strings.Repeat("A", 20)
	CiString20Str, err := CiString20(CiString20StrValue)

	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	if CiString20Str.String() != CiString20StrValue {
		t.Errorf(errExpectedStringOutput, CiString20Str.String())
	}

	err = CiString20Str.Validate()

	if err != nil {
		t.Errorf(errExpectedValidate, err)
	}
}

func TestCiString20TypeTooLong(t *testing.T) {
	t.Parallel()

	CiString20StrValue := strings.Repeat("A", 21)

	_, err := CiString20(CiString20StrValue)

	if err == nil {
		t.Errorf(errExpectedErrorTooLong)
	}
}

func TestCiString25TypeValid(t *testing.T) {
	t.Parallel()

	CiString25StrValue := strings.Repeat("A", 25)
	CiString25Str, err := CiString25(CiString25StrValue)

	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	if CiString25Str.String() != CiString25StrValue {
		t.Errorf(errExpectedStringOutput, CiString25Str.String())
	}

	err = CiString25Str.Validate()

	if err != nil {
		t.Errorf(errExpectedValidate, err)
	}
}

func TestCiString25TypeTooLong(t *testing.T) {
	t.Parallel()

	CiString25StrValue := strings.Repeat("A", 26)

	_, err := CiString25(CiString25StrValue)

	if err == nil {
		t.Errorf(errExpectedErrorTooLong)
	}
}

func TestCiString50TypeValid(t *testing.T) {
	t.Parallel()

	CiString50StrValue := strings.Repeat("A", 50)
	CiString50Str, err := CiString50(CiString50StrValue)

	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	if CiString50Str.String() != CiString50StrValue {
		t.Errorf(errExpectedStringOutput, CiString50Str.String())
	}

	err = CiString50Str.Validate()

	if err != nil {
		t.Errorf(errExpectedValidate, err)
	}
}

func TestCiString50TypeTooLong(t *testing.T) {
	t.Parallel()

	CiString50StrValue := strings.Repeat("A", 51)

	_, err := CiString50(CiString50StrValue)

	if err == nil {
		t.Errorf(errExpectedErrorTooLong)
	}
}

func TestCiString255TypeValid(t *testing.T) {
	t.Parallel()

	CiString255StrValue := strings.Repeat("A", 255)
	CiString255Str, err := CiString255(CiString255StrValue)

	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	if CiString255Str.String() != CiString255StrValue {
		t.Errorf(errExpectedStringOutput, CiString255Str.String())
	}

	err = CiString255Str.Validate()

	if err != nil {
		t.Errorf(errExpectedValidate, err)
	}
}

func TestCiString255TypeTooLong(t *testing.T) {
	t.Parallel()

	CiString255StrValue := strings.Repeat("A", 256)

	_, err := CiString255(CiString255StrValue)

	if err == nil {
		t.Errorf(errExpectedErrorTooLong)
	}
}

func TestCiString500TypeValid(t *testing.T) {
	t.Parallel()

	CiString500StrValue := strings.Repeat("A", 500)
	CiString500Str, err := CiString500(CiString500StrValue)

	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	if CiString500Str.String() != CiString500StrValue {
		t.Errorf(errExpectedStringOutput, CiString500Str.String())
	}

	err = CiString500Str.Validate()

	if err != nil {
		t.Errorf(errExpectedValidate, err)
	}
}

func TestCiString500TypeTooLong(t *testing.T) {
	t.Parallel()

	CiString500StrValue := strings.Repeat("A", 501)

	_, err := CiString500(CiString500StrValue)

	if err == nil {
		t.Errorf(errExpectedErrorTooLong)
	}
}

func TestCiStringTypeASCII(t *testing.T) {
	t.Parallel()

	var builder strings.Builder

	for i := 32; i <= 126; i++ {
		builder.WriteRune(rune(i))
	}

	input := builder.String()
	String255, err := CiString255(input)

	if err != nil {
		t.Errorf("expected all printable ASCII to pass validation, got %v", err)
	}

	if String255.String() != input {
		t.Errorf("unexpected output from String()")
	}

	if err := String255.Validate(); err != nil {
		t.Errorf(errExpectedValidate, err)
	}
}

func TestCiStringTypeChineseCharacters(t *testing.T) {
	t.Parallel()

	_, err := CiString20("你好，OCPP")

	if err == nil {
		t.Errorf("expected error for non-ASCII input (Chinese), got nil")
	}
}

func TestCiStringTypeUkrainianCharacters(t *testing.T) {
	t.Parallel()

	str, err := CiString20("Привіт, OCPP")

	if err == nil {
		t.Errorf("expected error for non-ASCII input (Ukrainian), got %s", str)
	}
}

func TestCiStringTypeArabicCharacters(t *testing.T) {
	t.Parallel()

	str, err := CiString20("مرحبًا OCPP")

	if err == nil {
		t.Errorf("expected error for non-ASCII input (Arabic), got %s", str)
	}
}

func TestCiStringTypeNonPrintable(t *testing.T) {
	t.Parallel()

	str, err := CiString20("Valid\x01String")

	if err == nil {
		t.Errorf("expected error for non-printable ASCII, got %s", str)
	}
}

func TestCiStringTypeEmpty(t *testing.T) {
	t.Parallel()

	_, err := CiString20("")

	if err == nil {
		t.Errorf("expected error for empty string, got nil")
	}
}

func TestCiStringTypeNullEquivalent(t *testing.T) {
	t.Parallel()

	var nullString string
	_, err := CiString20(nullString)

	if err == nil {
		t.Errorf("expected error for zero-value string (null equivalent), got nil")
	}
}
