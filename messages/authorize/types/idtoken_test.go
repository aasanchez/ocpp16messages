package types

import (
	"strings"
	"testing"

	sharedTypes "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	validTokenLength   = 16
	invalidTokenLength = 21
)

func TestIDToken(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("A", validTokenLength)
	cstr, _ := sharedTypes.NewCiString20(str)

	_, err := NewIDToken(cstr)
	if err != nil {
		t.Errorf("NewIDToken error: %v", err)
	}
}

func TestIDToken_string(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("B", validTokenLength)
	cstr, _ := sharedTypes.NewCiString20(str)

	idtoken, _ := NewIDToken(cstr)

	if idtoken.String() != str {
		t.Errorf(sharedTypes.ErrorMismatch, str, idtoken.String())
	}

	if idtoken.Value() != cstr {
		t.Errorf(sharedTypes.ErrorMismatch, cstr.Value(), idtoken.String())
	}
}

func TestIDToken_value(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("C", validTokenLength)
	cstr, _ := sharedTypes.NewCiString20(str)

	idtoken, _ := NewIDToken(cstr)

	if idtoken.Value() != cstr {
		t.Errorf(sharedTypes.ErrorMismatch, cstr.Value(), idtoken.String())
	}
}

func TestIDToken_invalid(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("D", invalidTokenLength)

	_, err := sharedTypes.NewCiString20(str)
	if err == nil {
		t.Errorf(sharedTypes.ErrorExpectedError, "length overflow")
	}
}

func TestIDToken_empty(t *testing.T) {
	t.Parallel()

	str := ""
	cstr, _ := sharedTypes.NewCiString20(str)

	_, err := NewIDToken(cstr)
	if err == nil {
		t.Errorf(sharedTypes.ErrorExpectedError, err)
	}
}
