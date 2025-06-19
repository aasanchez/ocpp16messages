package types

import (
	"strings"
	"testing"
)

const (
	errMsgExpectedValue   = "expected %q, got %q"
	errMsgExpectedIsValid = "expected IsValid to be true"
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

func TestCiStringReturnsErrorWhenEmpty(t *testing.T) {
	t.Parallel()

	_, err := CiString("", 20)
	expectError(t, err)
}

func TestCiStringReturnsErrorWhenTooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString(strings.Repeat("X", 21), 20)
	expectError(t, err)
}

func TestCiStringReturnsErrorWhenNonPrintable(t *testing.T) {
	t.Parallel()

	_, err := CiString("OCPP\x01", 20)
	expectError(t, err)
}

func TestCiStringNonASCII(t *testing.T) {
	t.Parallel()

	_, err := CiString("你好世界", 20) //nolint:gosmopolitan // Intentional: test Unicode handling
	expectError(t, err)
}

func TestCiStringReturnsNoErrorWhenValid(t *testing.T) {
	t.Parallel()

	_, err := CiString("Hello123", 20)
	expectNoError(t, err)
}

func TestCiStringStringReturnsOriginal(t *testing.T) {
	t.Parallel()

	str := "Hello123"
	cs, _ := CiString(str, 20)

	if cs.Value() != str {
		t.Errorf(errMsgExpectedValue, str, cs.Value())
	}
}

func TestCiStringValidateReturnsNilWhenValid(t *testing.T) {
	t.Parallel()

	cs, _ := CiString("GoodOne!", 20)
	err := cs.Validate()
	expectNoError(t, err)
}

func TestCiString20TypeCreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString20(strings.Repeat("A", 20))
	expectNoError(t, err)
}

func TestCiString20TypeCreateTooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString20(strings.Repeat("A", 21))
	expectError(t, err)
}

func TestCiString20TypeString(t *testing.T) {
	t.Parallel()

	s := strings.Repeat("A", 20)
	v, _ := CiString20(s)

	if v.Value() != s {
		t.Errorf(errMsgExpectedValue, s, v.Value())
	}
}

func TestCiString20TypeValidate(t *testing.T) {
	t.Parallel()

	v, _ := CiString20(strings.Repeat("A", 20))
	expectNoError(t, v.Validate())
}

func TestCiString20TypeIsValid(t *testing.T) {
	t.Parallel()

	v, _ := CiString20(strings.Repeat("A", 20))

	if !v.IsValid() {
		t.Error(errMsgExpectedIsValid)
	}
}

func TestCiString25TypeCreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString25(strings.Repeat("B", 25))
	expectNoError(t, err)
}

func TestCiString25TypeCreateTooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString25(strings.Repeat("B", 26))
	expectError(t, err)
}

func TestCiString25TypeString(t *testing.T) {
	t.Parallel()

	s := strings.Repeat("B", 25)
	v, _ := CiString25(s)

	if v.Value() != s {
		t.Errorf(errMsgExpectedValue, s, v.Value())
	}
}

func TestCiString25TypeValidate(t *testing.T) {
	t.Parallel()

	v, _ := CiString25(strings.Repeat("B", 25))
	expectNoError(t, v.Validate())
}

func TestCiString25TypeIsValid(t *testing.T) {
	t.Parallel()

	v, _ := CiString25(strings.Repeat("B", 25))

	if !v.IsValid() {
		t.Error(errMsgExpectedIsValid)
	}
}

func TestCiString50TypeCreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString50(strings.Repeat("C", 50))
	expectNoError(t, err)
}

func TestCiString50TypeCreateTooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString50(strings.Repeat("C", 51))
	expectError(t, err)
}

func TestCiString50TypeString(t *testing.T) {
	t.Parallel()

	s := strings.Repeat("C", 50)
	v, _ := CiString50(s)

	if v.Value() != s {
		t.Errorf(errMsgExpectedValue, s, v.Value())
	}
}

func TestCiString50TypeValidate(t *testing.T) {
	t.Parallel()

	v, _ := CiString50(strings.Repeat("C", 50))
	expectNoError(t, v.Validate())
}

func TestCiString50TypeIsValid(t *testing.T) {
	t.Parallel()

	v, _ := CiString50(strings.Repeat("C", 50))

	if !v.IsValid() {
		t.Error(errMsgExpectedIsValid)
	}
}

func TestCiString255TypeCreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString255(strings.Repeat("D", 255))
	expectNoError(t, err)
}

func TestCiString255TypeCreateTooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString255(strings.Repeat("D", 256))
	expectError(t, err)
}

func TestCiString255TypeString(t *testing.T) {
	t.Parallel()

	s := strings.Repeat("D", 255)
	v, _ := CiString255(s)

	if v.Value() != s {
		t.Errorf(errMsgExpectedValue, s, v.Value())
	}
}

func TestCiString255TypeValidate(t *testing.T) {
	t.Parallel()

	v, _ := CiString255(strings.Repeat("D", 255))
	expectNoError(t, v.Validate())
}

func TestCiString255TypeIsValid(t *testing.T) {
	t.Parallel()

	v, _ := CiString255(strings.Repeat("D", 255))

	if !v.IsValid() {
		t.Error(errMsgExpectedIsValid)
	}
}

func TestCiString500TypeCreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString500(strings.Repeat("E", 500))
	expectNoError(t, err)
}

func TestCiString500TypeCreateTooLong(t *testing.T) {
	t.Parallel()

	_, err := CiString500(strings.Repeat("E", 501))
	expectError(t, err)
}

func TestCiString500TypeString(t *testing.T) {
	t.Parallel()

	s := strings.Repeat("E", 500)
	v, _ := CiString500(s)

	if v.Value() != s {
		t.Errorf(errMsgExpectedValue, s, v.Value())
	}
}

func TestCiString500TypeValidate(t *testing.T) {
	t.Parallel()

	v, _ := CiString500(strings.Repeat("E", 500))
	expectNoError(t, v.Validate())
}

func TestCiString500TypeIsValid(t *testing.T) {
	t.Parallel()

	v, _ := CiString500(strings.Repeat("E", 500))

	if !v.IsValid() {
		t.Error(errMsgExpectedIsValid)
	}
}
