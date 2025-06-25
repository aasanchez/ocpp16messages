package types

import (
	"strings"
	"testing"
)

func TestCiString_EmptyFails(t *testing.T) {
	t.Parallel()

	_, err := CiString("", 20)
	if err == nil {
		t.Errorf(ErrExpectedNonNilError, "empty string")
	}
}

func TestCiString_TooLongFails(t *testing.T) {
	t.Parallel()

	_, err := CiString(strings.Repeat("A", 21), 20)
	if err == nil {
		t.Errorf(ErrExpectedNonNilError, "too long string")
	}
}

func TestCiString_NonPrintableFails(t *testing.T) {
	t.Parallel()
	//lint:ignore ST1018 intentional control character to test validation of non-printable input
	_, err := CiString("OCPP", 20) //nolint:staticcheck
	if err == nil {
		t.Errorf(ErrExpectedNonNilError, "non-printable characters")
	}
}

func TestCiString_ValidPasses(t *testing.T) {
	t.Parallel()

	_, err := CiString("OCPP16-Test", 20)
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString_ValueReturnsRaw(t *testing.T) {
	t.Parallel()

	input := "OCPP-Valid"
	cs, _ := CiString(input, 20)

	if cs.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, cs.Value())
	}
}

func TestCiString20Type_CreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString20("ABCDEFGHIJKLMNOPQRST")
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString20Type_CreateTooLongFails(t *testing.T) {
	t.Parallel()

	_, err := CiString20(strings.Repeat("X", 21))
	if err == nil {
		t.Errorf(ErrExpectedNonNilError, "> 20 characters")
	}
}

func TestCiString25Type_CreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString25(strings.Repeat("B", 25))
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString50Type_CreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString50(strings.Repeat("C", 50))
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString255Type_CreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString255(strings.Repeat("D", 255))
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString500Type_CreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString500(strings.Repeat("E", 500))
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString20Type_ValidatePasses(t *testing.T) {
	t.Parallel()

	cs, _ := CiString20("TestString20")
	if err := cs.Validate(); err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString25Type_ValidatePasses(t *testing.T) {
	t.Parallel()

	cs, _ := CiString25("TestString25ValidDataHere")
	if err := cs.Validate(); err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString50Type_ValidatePasses(t *testing.T) {
	t.Parallel()

	cs, _ := CiString50("This string is valid and under 50 chars")
	if err := cs.Validate(); err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString255Type_ValidatePasses(t *testing.T) {
	t.Parallel()

	cs, _ := CiString255(strings.Repeat("A", 100))
	if err := cs.Validate(); err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString500Type_ValidatePasses(t *testing.T) {
	t.Parallel()

	cs, _ := CiString500(strings.Repeat("B", 200))
	if err := cs.Validate(); err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}
