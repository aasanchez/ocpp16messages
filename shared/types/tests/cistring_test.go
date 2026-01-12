package types_test

import (
	"errors"
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// String length test values (one character over maximum allowed).
const (
	ciString20Over  = 21
	ciString25Over  = 26
	ciString50Over  = 51
	ciString255Over = 256
	ciString500Over = 501

	emptyString = ""
)

// Test error message formats.
const (
	errMsgValueMismatch   = "value mismatch: want %q, got %q"
	errMsgStringTooLong   = "expected error for string exceeding %d characters"
	errMsgUnexpectedError = "unexpected error for a %d-character string: %v"
	errMsgEmptyWantError  = "empty string should return error, got nil"
	errMsgWantErrEmptyVal = "expected ErrEmptyValue, got: %v"
)

// CiString20
func TestNewCiString20Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", st.CiString20Max)

	_, err := st.NewCiString20Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString20Max, err)
	}
}

func TestNewCiString20_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.NewCiString20Type(val)
	if err == nil {
		t.Error(errMsgEmptyWantError)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(errMsgWantErrEmptyVal, err)
	}
}

func TestNewCiString20_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", ciString20Over)

	_, err := st.NewCiString20Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString20Max)
	}
}

func TestNewCiString20_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", st.CiString20Max)

	ciStr, _ := st.NewCiString20Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString25
func TestNewCiString25Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", st.CiString25Max)

	_, err := st.NewCiString25Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString25Max, err)
	}
}

func TestNewCiString25_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.NewCiString25Type(val)
	if err == nil {
		t.Error(errMsgEmptyWantError)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(errMsgWantErrEmptyVal, err)
	}
}

func TestNewCiString25_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", ciString25Over)

	_, err := st.NewCiString25Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString25Max)
	}
}

func TestNewCiString25_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("F", st.CiString25Max)

	ciStr, _ := st.NewCiString25Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString50
func TestNewCiString50Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("G", st.CiString50Max)

	_, err := st.NewCiString50Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString50Max, err)
	}
}

func TestNewCiString50_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.NewCiString50Type(val)
	if err == nil {
		t.Error(errMsgEmptyWantError)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(errMsgWantErrEmptyVal, err)
	}
}

func TestNewCiString50_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("H", ciString50Over)

	_, err := st.NewCiString50Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString50Max)
	}
}

func TestNewCiString50_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("I", st.CiString50Max)

	ciStr, _ := st.NewCiString50Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString255
func TestNewCiString255Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("J", st.CiString255Max)

	_, err := st.NewCiString255Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString255Max, err)
	}
}

func TestNewCiString255_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.NewCiString255Type(val)
	if err == nil {
		t.Error(errMsgEmptyWantError)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(errMsgWantErrEmptyVal, err)
	}
}

func TestNewCiString255_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("K", ciString255Over)

	_, err := st.NewCiString255Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString255Max)
	}
}

func TestNewCiString255_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("L", st.CiString255Max)

	ciStr, _ := st.NewCiString255Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString500
func TestNewCiString500Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("M", st.CiString500Max)

	_, err := st.NewCiString500Type(val)
	if err != nil {
		t.Errorf(errMsgUnexpectedError, st.CiString500Max, err)
	}
}

func TestNewCiString500_Empty(t *testing.T) {
	t.Parallel()

	val := emptyString

	_, err := st.NewCiString500Type(val)
	if err == nil {
		t.Error(errMsgEmptyWantError)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(errMsgWantErrEmptyVal, err)
	}
}

func TestNewCiString500_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("N", ciString500Over)

	_, err := st.NewCiString500Type(val)
	if err == nil {
		t.Errorf(errMsgStringTooLong, st.CiString500Max)
	}
}

func TestNewCiString500_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("O", st.CiString500Max)

	ciStr, _ := st.NewCiString500Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}
