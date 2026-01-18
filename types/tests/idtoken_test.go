package types_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

const (
	validTokenLength   = 16
	invalidTokenLength = 21
)

func TestIdToken(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("A", validTokenLength)
	cstr, _ := types.NewCiString20Type(str)

	idtoken := types.NewIdToken(cstr)

	if idtoken.String() != str {
		t.Errorf(types.ErrorMismatch, str, idtoken.String())
	}
}

func TestIdToken_value(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("C", validTokenLength)
	cstr, _ := types.NewCiString20Type(str)

	idtoken := types.NewIdToken(cstr)

	if idtoken.Value() != cstr {
		t.Errorf(types.ErrorMismatch, cstr.Value(), idtoken.String())
	}
}

func TestIdToken_invalid(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("D", invalidTokenLength)

	_, err := types.NewCiString20Type(str)
	if err == nil {
		t.Errorf(types.ErrorExpectedError, err)
	}
}

func TestIdToken_empty(t *testing.T) {
	t.Parallel()

	str := ""

	_, err := types.NewCiString20Type(str)
	if err == nil {
		t.Error("NewCiString20Type() error = nil, want error for empty string")
	}
}

func TestIdToken_String(t *testing.T) {
	t.Parallel()

	str := "RFID-TAG-123"
	cstr, _ := types.NewCiString20Type(str)
	idtoken := types.NewIdToken(cstr)

	if idtoken.String() != str {
		t.Errorf(types.ErrorMismatch, str, idtoken.String())
	}
}
