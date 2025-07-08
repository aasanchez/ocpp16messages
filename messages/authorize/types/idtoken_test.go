package authorizetypes

import (
	"testing"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func Test_IdTokenFromCiString(t *testing.T) {
	t.Parallel()

	validStr := "ABC1234567890123456" // 20 characters

	str, err := sharedtypes.SetCiString20Type(validStr)
	if err != nil {
		t.Fatalf("failed to construct CiString20Type: %v", err)
	}

	_, err = IdToken(str)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func Test_IdTokenFromCiString_TooLong(t *testing.T) {
	t.Parallel()

	invalidStr := "ABC1234567890123456789" // 21 characters

	_, err := sharedtypes.SetCiString20Type(invalidStr)
	if err == nil {
		t.Fatalf("expected error when creating CiString20Type from over-length string, got nil")
	}
}
