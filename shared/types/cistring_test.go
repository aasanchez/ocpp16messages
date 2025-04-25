package types

import (
	"strings"
	"testing"
)

func expectNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func expectError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestCiString_ReturnsError_WhenEmpty(t *testing.T) {
	t.Parallel()

	_, err := CiString("", 20)
	expectError(t, err)
}

func TestCiString_ReturnsError_WhenTooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString(strings.Repeat("X", 21), 20)
	expectError(t, err)
}

func TestCiString_ReturnsError_WhenNonPrintable(t *testing.T) {
	t.Parallel()

	_, err := CiString("OCPP\x01", 20)
	expectError(t, err)
}

func TestCiString_ReturnsNoError_WhenValid(t *testing.T) {
	t.Parallel()

	_, err := CiString("Hello123", 20)
	expectNoError(t, err)
}

func TestCiString_String_ReturnsOriginal(t *testing.T) {
	t.Parallel()

	str := "Hello123"
	cs, _ := CiString(str, 20)

	if cs.String() != str {
		t.Errorf("expected %q, got %q", str, cs.String())
	}
}

func TestCiString_Validate_ReturnsNil_WhenValid(t *testing.T) {
	t.Parallel()

	cs, _ := CiString("GoodOne!", 20)
	err := cs.Validate()
	expectNoError(t, err)
}

func TestCiString20Type_Create_Valid(t *testing.T) {
	t.Parallel()

	_, err := CiString20(strings.Repeat("A", 20))
	expectNoError(t, err)
}

func TestCiString20Type_Create_TooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString20(strings.Repeat("A", 21))
	expectError(t, err)
}

func TestCiString20Type_String(t *testing.T) {
	t.Parallel()

	s := strings.Repeat("A", 20)
	v, _ := CiString20(s)

	if v.String() != s {
		t.Errorf("expected %q, got %q", s, v.String())
	}
}

func TestCiString20Type_Validate(t *testing.T) {
	t.Parallel()

	v, _ := CiString20(strings.Repeat("A", 20))
	expectNoError(t, v.Validate())
}

func TestCiString20Type_IsValid(t *testing.T) {
	t.Parallel()

	v, _ := CiString20(strings.Repeat("A", 20))

	if !v.IsValid() {
		t.Error("expected IsValid to be true")
	}
}

func TestCiString25Type_Create_Valid(t *testing.T) {
	t.Parallel()

	_, err := CiString25(strings.Repeat("B", 25))
	expectNoError(t, err)
}

func TestCiString25Type_Create_TooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString25(strings.Repeat("B", 26))
	expectError(t, err)
}

func TestCiString25Type_String(t *testing.T) {
	t.Parallel()

	s := strings.Repeat("B", 25)
	v, _ := CiString25(s)

	if v.String() != s {
		t.Errorf("expected %q, got %q", s, v.String())
	}
}

func TestCiString25Type_Validate(t *testing.T) {
	t.Parallel()

	v, _ := CiString25(strings.Repeat("B", 25))
	expectNoError(t, v.Validate())
}

func TestCiString25Type_IsValid(t *testing.T) {
	t.Parallel()

	v, _ := CiString25(strings.Repeat("B", 25))

	if !v.IsValid() {
		t.Error("expected IsValid to be true")
	}
}

func TestCiString50Type_Create_Valid(t *testing.T) {
	t.Parallel()

	_, err := CiString50(strings.Repeat("C", 50))
	expectNoError(t, err)
}

func TestCiString50Type_Create_TooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString50(strings.Repeat("C", 51))
	expectError(t, err)
}

func TestCiString50Type_String(t *testing.T) {
	t.Parallel()

	s := strings.Repeat("C", 50)
	v, _ := CiString50(s)

	if v.String() != s {
		t.Errorf("expected %q, got %q", s, v.String())
	}
}

func TestCiString50Type_Validate(t *testing.T) {
	t.Parallel()

	v, _ := CiString50(strings.Repeat("C", 50))
	expectNoError(t, v.Validate())
}

func TestCiString50Type_IsValid(t *testing.T) {
	t.Parallel()

	v, _ := CiString50(strings.Repeat("C", 50))

	if !v.IsValid() {
		t.Error("expected IsValid to be true")
	}
}

func TestCiString255Type_Create_Valid(t *testing.T) {
	t.Parallel()

	_, err := CiString255(strings.Repeat("D", 255))
	expectNoError(t, err)
}

func TestCiString255Type_Create_TooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString255(strings.Repeat("D", 256))
	expectError(t, err)
}

func TestCiString255Type_String(t *testing.T) {
	t.Parallel()

	s := strings.Repeat("D", 255)
	v, _ := CiString255(s)

	if v.String() != s {
		t.Errorf("expected %q, got %q", s, v.String())
	}
}

func TestCiString255Type_Validate(t *testing.T) {
	t.Parallel()

	v, _ := CiString255(strings.Repeat("D", 255))
	expectNoError(t, v.Validate())
}

func TestCiString255Type_IsValid(t *testing.T) {
	t.Parallel()

	v, _ := CiString255(strings.Repeat("D", 255))

	if !v.IsValid() {
		t.Error("expected IsValid to be true")
	}
}

func TestCiString500Type_Create_Valid(t *testing.T) {
	t.Parallel()

	_, err := CiString500(strings.Repeat("E", 500))
	expectNoError(t, err)
}

func TestCiString500Type_Create_TooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString500(strings.Repeat("E", 501))
	expectError(t, err)
}

func TestCiString500Type_String(t *testing.T) {
	t.Parallel()

	s := strings.Repeat("E", 500)
	v, _ := CiString500(s)

	if v.String() != s {
		t.Errorf("expected %q, got %q", s, v.String())
	}
}

func TestCiString500Type_Validate(t *testing.T) {
	t.Parallel()

	v, _ := CiString500(strings.Repeat("E", 500))
	expectNoError(t, v.Validate())
}

func TestCiString500Type_IsValid(t *testing.T) {
	t.Parallel()

	v, _ := CiString500(strings.Repeat("E", 500))

	if !v.IsValid() {
		t.Error("expected IsValid to be true")
	}
}
