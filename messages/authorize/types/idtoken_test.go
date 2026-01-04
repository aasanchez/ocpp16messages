package types

import (
	"strings"
	"testing"

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

	_, err := NewIdToken(cstr)
	if err != nil {
		t.Errorf("NewIdToken error: %v", err)
	}
}

func TestIdToken_value(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("C", validTokenLength)
	cstr, _ := st.NewCiString20Type(str)

	idtoken, _ := NewIdToken(cstr)

	if idtoken.Value() != cstr {
		t.Errorf(st.ErrorMismatch, cstr.Value(), idtoken.value.Value())
	}
}

func TestIdToken_invalid(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("D", invalidTokenLength)

	_, err := st.NewCiString20Type(str)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, "length overflow")
	}
}

func TestIdToken_empty(t *testing.T) {
	t.Parallel()

	str := ""
	cstr, _ := st.NewCiString20Type(str)

	_, err := NewIdToken(cstr)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}
