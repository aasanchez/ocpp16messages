package sharedtypes_test

import (
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// CiString20Type
func TestSetCiString20Type_Valid(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", 20)

	_, err := st.SetCiString20Type(val)
	if err != nil {
		t.Errorf("expected no error for 20-char string, got %v", err)
	}
}

func TestSetCiString20Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString20Type(val)
	if err != nil {
		t.Errorf("Empty value should be allowed")
	}
}

func TestSetCiString20Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", 21)

	_, err := st.SetCiString20Type(val)
	if err == nil {
		t.Errorf("expected error for string longer than 20 chars, got nil")
	}
}

func TestSetCiString20Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 20)

	ciStr, _ := st.SetCiString20Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// CiString25Type
func TestSetCiString25Type_Valid(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 25)

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf("expected no error for 25-char string, got %v", err)
	}
}

func TestSetCiString25Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString25Type(val)
	if err != nil {
		t.Errorf("Empty value should be allowed")
	}
}

func TestSetCiString25Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 26)

	_, err := st.SetCiString25Type(val)
	if err == nil {
		t.Errorf("expected error for string longer than 25 chars, got nil")
	}
}

func TestSetCiString25Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	ciStr, _ := st.SetCiString25Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// CiString50Type
func TestSetCiString50Type_Valid(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 50)

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf("expected no error for 50-char string, got %v", err)
	}
}

func TestSetCiString50Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString50Type(val)
	if err != nil {
		t.Errorf("Empty value should be allowed")
	}
}

func TestSetCiString50Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 51)

	_, err := st.SetCiString50Type(val)
	if err == nil {
		t.Errorf("expected error for string longer than 50 chars, got nil")
	}
}

func TestSetCiString50Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	ciStr, _ := st.SetCiString50Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// CiString255Type
func TestSetCiString255Type_Valid(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 255)

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf("expected no error for 255-char string, got %v", err)
	}
}

func TestSetCiString255Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString255Type(val)
	if err != nil {
		t.Errorf("Empty value should be allowed")
	}
}

func TestSetCiString255Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 256)

	_, err := st.SetCiString255Type(val)
	if err == nil {
		t.Errorf("expected error for string longer than 50 chars, got nil")
	}
}

func TestSetCiString255Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	ciStr, _ := st.SetCiString255Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// CiString500Type
func TestSetCiString500Type_Valid(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 500)

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf("expected no error for 255-char string, got %v", err)
	}
}

func TestSetCiString500Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := st.SetCiString500Type(val)
	if err != nil {
		t.Errorf("Empty value should be allowed")
	}
}

func TestSetCiString500Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 501)

	_, err := st.SetCiString500Type(val)
	if err == nil {
		t.Errorf("expected error for string longer than 50 chars, got nil")
	}
}

func TestSetCiString500Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	ciStr, _ := st.SetCiString500Type(input)
	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}
