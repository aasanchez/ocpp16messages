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

func TestIDToken(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("A", validTokenLength)
	cstr, _ := st.NewCiString20(str)

	_, err := NewIDToken(cstr)
	if err != nil {
		t.Errorf("NewIDToken error: %v", err)
	}
}

func TestIDToken_string(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("B", validTokenLength)
	cstr, _ := st.NewCiString20(str)

	idtoken, _ := NewIDToken(cstr)

	if idtoken.String() != str {
		t.Errorf(st.ErrorMismatch, str, idtoken.String())
	}

	if idtoken.Value() != cstr {
		t.Errorf(st.ErrorMismatch, cstr.Value(), idtoken.String())
	}
}

func TestIDToken_value(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("C", validTokenLength)
	cstr, _ := st.NewCiString20(str)

	idtoken, _ := NewIDToken(cstr)

	if idtoken.Value() != cstr {
		t.Errorf(st.ErrorMismatch, cstr.Value(), idtoken.String())
	}
}

func TestIDToken_invalid(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("D", invalidTokenLength)

	_, err := st.NewCiString20(str)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, "length overflow")
	}
}

func TestIDToken_empty(t *testing.T) {
	t.Parallel()

	str := ""
	cstr, _ := st.NewCiString20(str)

	_, err := NewIDToken(cstr)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}
