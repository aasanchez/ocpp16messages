package sharedtypes_test

import (
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	ciString20Over  = 21
	ciString25Over  = 26
	ciString50Over  = 51
	ciString255Over = 256
	ciString500Over = 501

	emptyString = ""

	errMsgValueMismatch   = "value mismatch: want %q, got %q"
	errMsgStringTooLong   = "expected error for string exceeding %d characters"
	errMsgUnexpectedError = "unexpected error for a %d-character string: %v"
	errMsgEmptyAccepted   = "empty string should be accepted; got error: %v"
)

// CiString20Type
func Test_sharedtypes_SetCiString20Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", st.CiString20Max)

	_, err := st.SetCiString20Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString20Max, err)
	}
}

func Test_sharedtypes_SetCiString20Type_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.SetCiString20Type(val)
	if err != nil {
		t.Errorf(errMsgEmptyAccepted, err)
	}
}

func Test_sharedtypes_SetCiString20Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", ciString20Over)

	_, err := st.SetCiString20Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString20Max)
	}
}

func Test_sharedtypes_SetCiString20Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", st.CiString20Max)

	ciStr, _ := st.SetCiString20Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString25Type
func Test_sharedtypes_SetCiString25Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", st.CiString25Max)

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString25Max, err)
	}
}

func Test_sharedtypes_SetCiString25Type_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf(errMsgEmptyAccepted, err)
	}
}

func Test_sharedtypes_SetCiString25Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", ciString25Over)

	_, err := st.SetCiString25Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString25Max)
	}
}

func Test_sharedtypes_SetCiString25Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("F", st.CiString25Max)

	ciStr, _ := st.SetCiString25Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString50Type
func Test_sharedtypes_SetCiString50Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("G", st.CiString50Max)

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString50Max, err)
	}
}

func Test_sharedtypes_SetCiString50Type_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf(errMsgEmptyAccepted, err)
	}
}

func Test_sharedtypes_SetCiString50Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("H", ciString50Over)

	_, err := st.SetCiString50Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString50Max)
	}
}

func Test_sharedtypes_SetCiString50Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("I", st.CiString50Max)

	ciStr, _ := st.SetCiString50Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString255Type
func Test_sharedtypes_SetCiString255Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("J", st.CiString255Max)

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString255Max, err)
	}
}

func Test_sharedtypes_SetCiString255Type_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf(errMsgEmptyAccepted, err)
	}
}

func Test_sharedtypes_SetCiString255Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("K", ciString255Over)

	_, err := st.SetCiString255Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString255Max)
	}
}

func Test_sharedtypes_SetCiString255Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("L", st.CiString255Max)

	ciStr, _ := st.SetCiString255Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString500Type
func Test_sharedtypes_SetCiString500Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("M", st.CiString500Max)

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString500Max, err)
	}
}

func Test_sharedtypes_SetCiString500Type_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf(errMsgEmptyAccepted, err)
	}
}

func Test_sharedtypes_SetCiString500Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("N", ciString500Over)

	_, err := st.SetCiString500Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString500Max)
	}
}

func Test_sharedtypes_SetCiString500Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("O", st.CiString500Max)

	ciStr, _ := st.SetCiString500Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}
