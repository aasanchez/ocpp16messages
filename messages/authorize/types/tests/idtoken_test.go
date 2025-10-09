package authorizetypes_test

import (
	"fmt"
	"testing"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestIdToken_ConstructAndAccessors_MaxLen(t *testing.T) {
	t.Parallel()

	str := "nVIWxwYbHBmsRbI6"
	cstr, _ := st.SetCiString20Type(str)

	idtoken, err := mat.SetIdToken(cstr)
	if err != nil {
		t.Fatalf("SetIdToken error: %v", err)
	}

	if idtoken.String() != str {
		t.Errorf(st.ErrorStringMismatch,
			fmt.Sprintf("%q", str),
			fmt.Sprintf("%q", idtoken.String()))
	}

	if idtoken.Value() != cstr {
		t.Errorf(st.ErrorValueMismatch, cstr.Value(), idtoken.String())
	}
}
