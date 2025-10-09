package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestIdToken_ConstructAndAccessors_ABC(t *testing.T) {
	s := "abc"

	cs, err := st.SetCiString20Type(s)
	if err != nil {
		t.Fatalf("SetCiString20Type(%q) error: %v", s, err)
	}

	tok, err := authorizetypes.SetIdToken(cs)
	if err != nil {
		t.Fatalf("SetIdToken error: %v", err)
	}

	if got := tok.String(); got != s {
		t.Errorf("String() = %q, want %q", got, s)
	}

	if tok.Value() != cs {
		t.Errorf("Value() = %#v, want %#v", tok.Value(), cs)
	}

}

func TestIdToken_ConstructAndAccessors_MaxLen(t *testing.T) {
	s := "nVIWxwYbHBmsRbI6"
	cs, _ := st.SetCiString20Type(s)

	idtoken, err := authorizetypes.SetIdToken(cs)
	if err != nil {
		t.Fatalf("SetIdToken error: %v", err)
	}

	if idtoken.String() != s {
		t.Errorf("String() = %q, want %q", idtoken.String(), s)
	}

	if idtoken.Value() != cs {
		t.Errorf("Value() = %#v, want %#v", idtoken.Value(), cs)
	}

}

func TestIdToken_ZeroValueAccessors(t *testing.T) {
	var zero st.CiString20Type
	ok, err := authorizetypes.SetIdToken(zero)
	if err != nil {
		t.Fatalf("SetIdToken(zero) error: %v", err)
	}
	if ok.String() != "" {
		t.Errorf("String() zero = %q, want empty", ok.String())
	}
	if ok.Value().Value() != "" {
		t.Errorf("Value().Value() zero = %q, want empty", ok.Value().Value())
	}
}
