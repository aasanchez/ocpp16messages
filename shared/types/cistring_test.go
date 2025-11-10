package sharedtypes

import (
	"errors"
	"strings"
	"testing"
)

func Test_sharedtypes_setCiString_Valid(t *testing.T) {
	t.Parallel()

	input := "OCPP-Test-123"
	maxLen := 20

	cis, err := setCiString(input, maxLen)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}

	if got := cis.val(); got != input {
		t.Errorf(ErrorValueMismatch, got, input)
	}
}

func Test_sharedtypes_setCiString_TooLong(t *testing.T) {
	t.Parallel()

	input := "ThisStringIsWayTooLong"
	maxLen := 5

	_, err := setCiString(input, maxLen)
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}

	if !errors.Is(err, errExceedsMaxLength) {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_setCiString_NonPrintableASCII(t *testing.T) {
	t.Parallel()

	input := "Valid\x01Invalid"

	_, err := setCiString(input, 20)
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}

	if !errors.Is(err, errInvalidASCII) {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_setCiString_EmptyIsValid(t *testing.T) {
	t.Parallel()

	input := ""
	maxLen := 10

	cis, err := setCiString(input, maxLen)
	if err != nil {
		t.Errorf(ErrorExpectedError, err)
	}

	if got := cis.val(); got != "" {
		t.Errorf(ErrorValueMismatch, input, got)
	}
}

func Test_sharedtypes_setCiString_ASCIIBoundary(t *testing.T) {
	t.Parallel()

	input := " ~"
	maxLen := 5

	cis, err := setCiString(input, maxLen)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}

	if got := cis.val(); got != input {
		t.Errorf(ErrorValueMismatch, got, input)
	}
}

func Test_sharedtypes_SetCiString20Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", 20)

	_, err := SetCiString20Type(val)
	if err != nil {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_SetCiString20Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := SetCiString20Type(val)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}
}

func Test_sharedtypes_SetCiString20Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", 21)

	_, err := SetCiString20Type(val)
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_SetCiString20Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 20)

	ciStr, _ := SetCiString20Type(input)
	if ciStr.Value() != input {
		t.Errorf(ErrorValueMismatch, input, ciStr.Value())
	}
}

func Test_sharedtypes_SetCiString25Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 25)

	_, err := SetCiString25Type(val)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}
}

func Test_sharedtypes_SetCiString25Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := SetCiString25Type(val)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}
}

func Test_sharedtypes_SetCiString25Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 26)

	_, err := SetCiString25Type(val)
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_SetCiString25Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	ciStr, _ := SetCiString25Type(input)
	if ciStr.Value() != input {
		t.Errorf(ErrorValueMismatch, input, ciStr.Value())
	}
}

func Test_sharedtypes_SetCiString50Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 50)

	_, err := SetCiString50Type(val)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}
}

func Test_sharedtypes_SetCiString50Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := SetCiString50Type(val)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}
}

func Test_sharedtypes_SetCiString50Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 51)

	_, err := SetCiString50Type(val)
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_SetCiString50Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	ciStr, _ := SetCiString50Type(input)
	if ciStr.Value() != input {
		t.Errorf(ErrorValueMismatch, input, ciStr.Value())
	}
}

func Test_sharedtypes_SetCiString255Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 255)

	_, err := SetCiString255Type(val)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}
}

func Test_sharedtypes_SetCiString255Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := SetCiString255Type(val)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}
}

func Test_sharedtypes_SetCiString255Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 256)

	_, err := SetCiString255Type(val)
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_SetCiString255Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	ciStr, _ := SetCiString255Type(input)
	if ciStr.Value() != input {
		t.Errorf(ErrorValueMismatch, input, ciStr.Value())
	}
}

func Test_sharedtypes_SetCiString500Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 500)

	_, err := SetCiString500Type(val)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}
}

func Test_sharedtypes_SetCiString500Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := SetCiString500Type(val)
	if err != nil {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_SetCiString500Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 501)

	_, err := SetCiString500Type(val)
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_SetCiString500Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	ciStr, _ := SetCiString500Type(input)
	if ciStr.Value() != input {
		t.Errorf(ErrorValueMismatch, input, ciStr.Value())
	}
}
