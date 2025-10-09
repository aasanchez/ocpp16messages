package authorizetypes_test

import (
	"strings"
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

	if got := tok.Value().Value(); got != s {
		t.Errorf("Value().Value() = %q, want %q", got, s)
	}
}

func TestIdToken_ConstructAndAccessors_MaxLen(t *testing.T) {
	s := strings.Repeat("x", 20) // boundary for CiString20Type

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

	if got := tok.Value().Value(); got != s {
		t.Errorf("Value().Value() = %q, want %q", got, s)
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
