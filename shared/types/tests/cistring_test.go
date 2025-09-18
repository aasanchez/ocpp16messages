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
		t.Errorf("expected no error for 20-char string, got %v", err)
	}
}

func Test_sharedtypes_SetCiString20Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString20Type(val)
	if err != nil {
		t.Errorf("Empty value should be allowed")
	}
}

func Test_sharedtypes_SetCiString20Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", 21)

	_, err := st.SetCiString20Type(val)
	if err == nil {
		t.Errorf("expected error for string longer than 20 chars, got nil")
	}
}

func Test_sharedtypes_SetCiString20Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 20)

	ciStr, _ := st.SetCiString20Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// CiString25Type
func Test_sharedtypes_SetCiString25Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 25)

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf("expected no error for 25-char string, got %v", err)
	}
}

func Test_sharedtypes_SetCiString25Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf("Empty value should be allowed")
	}
}

func Test_sharedtypes_SetCiString25Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 26)

	_, err := st.SetCiString25Type(val)
	if err == nil {
		t.Errorf("expected error for string longer than 25 chars, got nil")
	}
}

func Test_sharedtypes_SetCiString25Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	ciStr, _ := st.SetCiString25Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// CiString50Type
func Test_sharedtypes_SetCiString50Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 50)

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf("expected no error for 50-char string, got %v", err)
	}
}

func Test_sharedtypes_SetCiString50Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf("Empty value should be allowed")
	}
}

func Test_sharedtypes_SetCiString50Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 51)

	_, err := st.SetCiString50Type(val)
	if err == nil {
		t.Errorf("expected error for string longer than 50 chars, got nil")
	}
}

func Test_sharedtypes_SetCiString50Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	ciStr, _ := st.SetCiString50Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// CiString255Type
func Test_sharedtypes_SetCiString255Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 255)

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf("expected no error for 255-char string, got %v", err)
	}
}

func Test_sharedtypes_SetCiString255Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf("Empty value should be allowed")
	}
}

func Test_sharedtypes_SetCiString255Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 256)

	_, err := st.SetCiString255Type(val)
	if err == nil {
		t.Errorf("expected error for string longer than 50 chars, got nil")
	}
}

func Test_sharedtypes_SetCiString255Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	ciStr, _ := st.SetCiString255Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// CiString500Type
func Test_sharedtypes_SetCiString500Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 500)

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf("expected no error for 255-char string, got %v", err)
	}
}

func Test_sharedtypes_SetCiString500Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf("Empty value should be allowed")
	}
}

func Test_sharedtypes_SetCiString500Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 501)

	_, err := st.SetCiString500Type(val)
	if err == nil {
		t.Errorf("expected error for string longer than 50 chars, got nil")
	}
}

func Test_sharedtypes_SetCiString500Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	ciStr, _ := st.SetCiString500Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}
