package authorizetypes_test

import (
	"fmt"
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestIdToken_ConstructAndAccessors_MaxLen(t *testing.T) {
	t.Parallel()

	s := "nVIWxwYbHBmsRbI6"
	cs, _ := st.SetCiString20Type(s)

	idtoken, err := authorizetypes.SetIdToken(cs)
	if err != nil {
		t.Fatalf("SetIdToken error: %v", err)
	}

	if idtoken.String() != s {
		t.Errorf(st.ErrorStringMismatch, fmt.Sprintf("%q", s), fmt.Sprintf("%q", idtoken.String()))
	}

	if idtoken.Value() != cs {
		t.Errorf(st.ErrorValueMismatch, cs.Value(), idtoken.String())
	}

}
