package sharedtypes_test

import (
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	ciString20Max  = 20
	ciString20Over = 21
	ciString25Max  = 25
	ciString25Over = 26
	ciString50Max  = 50
	ciString50Over = 51
	ciString255Max  = 255
	ciString255Over = 256
	ciString500Max  = 500
	ciString500Over = 501

	errMsgValueMismatch = "value mismatch: want %q, got %q"
)

// CiString20Type
func Test_sharedtypes_SetCiString20Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", ciString20Max)

	_, err := st.SetCiString20Type(val)
	if err != nil {
		t.Errorf("unexpected error for a %d-character string: %v", ciString20Max, err)
	}
}

func Test_sharedtypes_SetCiString20Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString20Type(val)
	if err != nil {
		t.Errorf("empty string should be accepted; got error: %v", err)
	}
}

func Test_sharedtypes_SetCiString20Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", ciString20Over)

	_, err := st.SetCiString20Type(val)
	if err == nil {
		t.Errorf("expected an error for a string longer than %d characters; got none", ciString20Max)
	}
}

func Test_sharedtypes_SetCiString20Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", ciString20Max)

	ciStr, _ := st.SetCiString20Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString25Type
func Test_sharedtypes_SetCiString25Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", ciString25Max)

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf("unexpected error for a %d-character string: %v", ciString25Max, err)
	}
}

func Test_sharedtypes_SetCiString25Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf("empty string should be accepted; got error: %v", err)
	}
}

func Test_sharedtypes_SetCiString25Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", ciString25Over)

	_, err := st.SetCiString25Type(val)
	if err == nil {
		t.Errorf("expected an error for a string longer than %d characters; got none", ciString25Max)
	}
}

func Test_sharedtypes_SetCiString25Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("F", ciString25Max)

	ciStr, _ := st.SetCiString25Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString50Type
func Test_sharedtypes_SetCiString50Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("G", ciString50Max)

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf("unexpected error for a %d-character string: %v", ciString50Max, err)
	}
}

func Test_sharedtypes_SetCiString50Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf("empty string should be accepted; got error: %v", err)
	}
}

func Test_sharedtypes_SetCiString50Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("H", ciString50Over)

	_, err := st.SetCiString50Type(val)
	if err == nil {
		t.Errorf("expected an error for a string longer than %d characters; got none", ciString50Max)
	}
}

func Test_sharedtypes_SetCiString50Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("I", ciString50Max)

	ciStr, _ := st.SetCiString50Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString255Type
func Test_sharedtypes_SetCiString255Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("J", ciString255Max)

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf(
			"unexpected error for a %d-character string: %v",
			ciString255Max, err,
		)
	}
}

func Test_sharedtypes_SetCiString255Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf("empty string should be accepted; got error: %v", err)
	}
}

func Test_sharedtypes_SetCiString255Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("K", ciString255Over)

	_, err := st.SetCiString255Type(val)
	if err == nil {
		t.Errorf(
			"expected error for string exceeding %d characters",
			ciString255Max,
		)
	}
}

func Test_sharedtypes_SetCiString255Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("L", ciString255Max)

	ciStr, _ := st.SetCiString255Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}

// CiString500Type
func Test_sharedtypes_SetCiString500Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("M", ciString500Max)

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf(
			"unexpected error for a %d-character string: %v",
			ciString500Max, err,
		)
	}
}

func Test_sharedtypes_SetCiString500Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf("empty string should be accepted; got error: %v", err)
	}
}

func Test_sharedtypes_SetCiString500Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("N", ciString500Over)

	_, err := st.SetCiString500Type(val)
	if err == nil {
		t.Errorf(
			"expected error for string exceeding %d characters",
			ciString500Max,
		)
	}
}

func Test_sharedtypes_SetCiString500Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("O", ciString500Max)

	ciStr, _ := st.SetCiString500Type(input)
	if ciStr.Value() != input {
		t.Errorf(errMsgValueMismatch, input, ciStr.Value())
	}
}
