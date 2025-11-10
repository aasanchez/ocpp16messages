package authorizetypes

import (
	"fmt"
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestIdToken(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("A", 16)
	cstr, _ := st.SetCiString20Type(str)

	_, err := SetIdToken(cstr)
	if err != nil {
		t.Errorf("SetIdToken error: %v", err)
	}
}

func TestIdToken_string(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("B", 16)
	cstr, _ := st.SetCiString20Type(str)

	idtoken, _ := SetIdToken(cstr)

	if idtoken.String() != str {
		t.Errorf(st.ErrorStringMismatch,
			fmt.Sprintf("%q", str),
			fmt.Sprintf("%q", idtoken.String()))
	}

	if idtoken.Value() != cstr {
		t.Errorf(st.ErrorValueMismatch, cstr.Value(), idtoken.String())
	}
}

func TestIdToken_value(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("C", 16)
	cstr, _ := st.SetCiString20Type(str)

	idtoken, _ := SetIdToken(cstr)

	if idtoken.Value() != cstr {
		t.Errorf(st.ErrorValueMismatch, cstr.Value(), idtoken.String())
	}
}

func TestIdToken_invalid(t *testing.T) {
	t.Parallel()

	str := strings.Repeat("D", 21)

	_, err := st.SetCiString20Type(str)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, "length overflow")
	}
}

func TestIdToken_empty(t *testing.T) {
	t.Parallel()

	str := ""
	cstr, _ := st.SetCiString20Type(str)

	_, err := SetIdToken(cstr)
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}
