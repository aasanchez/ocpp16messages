package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core/types"
)

func TestNewIdTokenValid(t *testing.T) {
	const validIdToken = "valid-id-token"

	idToken, err := types.NewIdToken(validIdToken)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if idToken.String() != validIdToken {
		t.Errorf("unexpected value: got %q, want %q", idToken.String(), validIdToken)
	}
}

func TestNewIdTokenTooLong(t *testing.T) {
	longValue := "this-id-tag-is-way-too-long-for-ocpp"
	_, err := types.NewIdToken(longValue)
	if err == nil {
		t.Fatal("expected error for too long idToken, got nil")
	}
}

func TestIdTokenStringNilReceiver(t *testing.T) {
	var idToken *types.IdToken
	if idToken.String() != "" {
		t.Errorf("expected empty string from nil receiver, got %q", idToken.String())
	}
}
