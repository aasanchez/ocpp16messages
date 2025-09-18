package sharedtypes_test

import (
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// CiString20Type
func Test_sharedtypes_SetCiString20Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", 20)

	_, err := st.SetCiString20Type(val)
	if err != nil {
		t.Errorf("unexpected error for a 20-character string: %v", err)
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

	val := strings.Repeat("A", 21)

	_, err := st.SetCiString20Type(val)
	if err == nil {
		t.Error("expected an error for a string longer than 20 characters; got none")
	}
}

func Test_sharedtypes_SetCiString20Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 20)

	ciStr, _ := st.SetCiString20Type(input)
	if ciStr.Value() != input {
		t.Errorf("value mismatch: want %q, got %q", input, ciStr.Value())
	}
}

// CiString25Type
func Test_sharedtypes_SetCiString25Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 25)

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf("unexpected error for a 25-character string: %v", err)
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

	val := strings.Repeat("B", 26)

	_, err := st.SetCiString25Type(val)
	if err == nil {
		t.Error("expected an error for a string longer than 25 characters; got none")
	}
}

func Test_sharedtypes_SetCiString25Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	ciStr, _ := st.SetCiString25Type(input)
	if ciStr.Value() != input {
		t.Errorf("value mismatch: want %q, got %q", input, ciStr.Value())
	}
}

// CiString50Type
func Test_sharedtypes_SetCiString50Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 50)

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf("unexpected error for a 50-character string: %v", err)
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

	val := strings.Repeat("C", 51)

	_, err := st.SetCiString50Type(val)
	if err == nil {
		t.Error("expected an error for a string longer than 50 characters; got none")
	}
}

func Test_sharedtypes_SetCiString50Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	ciStr, _ := st.SetCiString50Type(input)
	if ciStr.Value() != input {
		t.Errorf("value mismatch: want %q, got %q", input, ciStr.Value())
	}
}

// CiString255Type
func Test_sharedtypes_SetCiString255Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 255)

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf("unexpected error for a 255-character string: %v", err)
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

	val := strings.Repeat("D", 256)

	_, err := st.SetCiString255Type(val)
	if err == nil {
		t.Error("expected an error for a string longer than 255 characters; got none")
	}
}

func Test_sharedtypes_SetCiString255Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	ciStr, _ := st.SetCiString255Type(input)
	if ciStr.Value() != input {
		t.Errorf("value mismatch: want %q, got %q", input, ciStr.Value())
	}
}

// CiString500Type
func Test_sharedtypes_SetCiString500Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 500)

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf("unexpected error for a 500-character string: %v", err)
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

	val := strings.Repeat("E", 501)

	_, err := st.SetCiString500Type(val)
	if err == nil {
		t.Error("expected an error for a string longer than 500 characters; got none")
	}
}

func Test_sharedtypes_SetCiString500Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	ciStr, _ := st.SetCiString500Type(input)
	if ciStr.Value() != input {
		t.Errorf("value mismatch: want %q, got %q", input, ciStr.Value())
	}
}
