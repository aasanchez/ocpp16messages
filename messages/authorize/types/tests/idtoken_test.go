package types_test

import (
	"strings"
	"testing"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	validTokenLength   = 16
	invalidTokenLength = 21
)

func TestIdToken(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("A", validTokenLength)
	cstr, _ := st.NewCiString20Type(str)

	idtoken := mat.NewIdToken(cstr)

	if idtoken.String() != str {
		t.Errorf(st.ErrorMismatch, str, idtoken.String())
	}
}

func TestIdToken_value(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("C", validTokenLength)
	cstr, _ := st.NewCiString20Type(str)

	idtoken := mat.NewIdToken(cstr)

	if idtoken.Value() != cstr {
		t.Errorf(st.ErrorMismatch, cstr.Value(), idtoken.String())
	}
}

func TestIdToken_invalid(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("D", invalidTokenLength)

	_, err := st.NewCiString20Type(str)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

func TestIdToken_empty(t *testing.T) {
	t.Parallel()

	str := ""

	_, err := st.NewCiString20Type(str)
	if err == nil {
		t.Error("NewCiString20Type() error = nil, want error for empty string")
	}
}

func TestIdToken_String(t *testing.T) {
	t.Parallel()

	str := "RFID-TAG-123"
	cstr, _ := st.NewCiString20Type(str)
	idtoken := mat.NewIdToken(cstr)

	if idtoken.String() != str {
		t.Errorf(st.ErrorMismatch, str, idtoken.String())
	}
}
