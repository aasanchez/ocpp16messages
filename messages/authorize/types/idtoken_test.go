package types

import (
	"fmt"
	"strings"
	"testing"

	sharedTypes "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	validTokenLength   = 16
	invalidTokenLength = 21
)

func TestIdToken(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("A", validTokenLength)
	cstr, _ := sharedTypes.NewCiString20(str)

	_, err := SetIdToken(cstr)
	if err != nil {
		t.Errorf("SetIdToken error: %v", err)
	}
}

func TestIdToken_string(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("B", validTokenLength)
	cstr, _ := sharedTypes.NewCiString20(str)

	idtoken, _ := SetIdToken(cstr)

	if idtoken.String() != str {
		t.Errorf(sharedTypes.ErrorStringMismatch,
			fmt.Sprintf("%q", str),
			fmt.Sprintf("%q", idtoken.String()))
	}

	if idtoken.Value() != cstr {
		t.Errorf(sharedTypes.ErrorValueMismatch, cstr.Value(), idtoken.String())
	}
}

func TestIdToken_value(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("C", validTokenLength)
	cstr, _ := sharedTypes.NewCiString20(str)

	idtoken, _ := SetIdToken(cstr)

	if idtoken.Value() != cstr {
		t.Errorf(sharedTypes.ErrorValueMismatch, cstr.Value(), idtoken.String())
	}
}

func TestIdToken_invalid(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("D", invalidTokenLength)

	_, err := sharedTypes.NewCiString20(str)
	if err == nil {
		t.Errorf(sharedTypes.ErrorExpectedError, "length overflow")
	}
}

func TestIdToken_empty(t *testing.T) {
	t.Parallel()

	str := ""
	cstr, _ := sharedTypes.NewCiString20(str)

	_, err := SetIdToken(cstr)
	if err == nil {
		t.Errorf(sharedTypes.ErrorExpectedError, err)
	}
}
