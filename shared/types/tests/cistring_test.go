package types_test

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

// CiString20
func Test_sharedtypes_NewCiString20Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", st.CiString20Max)

	_, err := st.NewCiString20Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString20Max, err)
	}
}

func Test_sharedtypes_NewCiString20_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.NewCiString20Type(val)
	if err != nil {
		t.Errorf(errMsgEmptyAccepted, err)
	}
}

func Test_sharedtypes_NewCiString20_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", ciString20Over)

	_, err := st.NewCiString20Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString20Max)
	}
}

func Test_sharedtypes_NewCiString20_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", st.CiString20Max)

	ciStr, _ := st.NewCiString20Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString25
func Test_sharedtypes_NewCiString25Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", st.CiString25Max)

	_, err := st.NewCiString25Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString25Max, err)
	}
}

func Test_sharedtypes_NewCiString25_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.NewCiString25Type(val)
	if err != nil {
		t.Errorf(errMsgEmptyAccepted, err)
	}
}

func Test_sharedtypes_NewCiString25_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", ciString25Over)

	_, err := st.NewCiString25Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString25Max)
	}
}

func Test_sharedtypes_NewCiString25_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("F", st.CiString25Max)

	ciStr, _ := st.NewCiString25Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString50
func Test_sharedtypes_NewCiString50Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("G", st.CiString50Max)

	_, err := st.NewCiString50Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString50Max, err)
	}
}

func Test_sharedtypes_NewCiString50_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.NewCiString50Type(val)
	if err != nil {
		t.Errorf(errMsgEmptyAccepted, err)
	}
}

func Test_sharedtypes_NewCiString50_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("H", ciString50Over)

	_, err := st.NewCiString50Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString50Max)
	}
}

func Test_sharedtypes_NewCiString50_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("I", st.CiString50Max)

	ciStr, _ := st.NewCiString50Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString255
func Test_sharedtypes_NewCiString255Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("J", st.CiString255Max)

	_, err := st.NewCiString255Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString255Max, err)
	}
}

func Test_sharedtypes_NewCiString255_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.NewCiString255Type(val)
	if err != nil {
		t.Errorf(errMsgEmptyAccepted, err)
	}
}

func Test_sharedtypes_NewCiString255_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("K", ciString255Over)

	_, err := st.NewCiString255Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString255Max)
	}
}

func Test_sharedtypes_NewCiString255_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("L", st.CiString255Max)

	ciStr, _ := st.NewCiString255Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString500
func Test_sharedtypes_NewCiString500Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("M", st.CiString500Max)

	_, err := st.NewCiString500Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString500Max, err)
	}
}

func Test_sharedtypes_NewCiString500_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.NewCiString500Type(val)
	if err != nil {
		t.Errorf(errMsgEmptyAccepted, err)
	}
}

func Test_sharedtypes_NewCiString500_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("N", ciString500Over)

	_, err := st.NewCiString500Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString500Max)
	}
}

func Test_sharedtypes_NewCiString500_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("O", st.CiString500Max)

	ciStr, _ := st.NewCiString500Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}
