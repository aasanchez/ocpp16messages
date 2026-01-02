package types_test

import (
	"testing"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func Test_IdTokenFromCiString(t *testing.T) {
	t.Parallel()

	validStr := "ABC1234567890123456" // 20 characters

	str, err := st.NewCiString20Type(validStr)
	if err != nil {
		t.Fatalf("failed to construct CiString20: %v", err)
	}

	_, err = mat.NewIdToken(str)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func Test_IdTokenFromCiString_TooLong(t *testing.T) {
	t.Parallel()

	invalidStr := "ABC1234567890123456789" // 21 characters

	_, err := st.NewCiString20Type(invalidStr)
	if err == nil {
		t.Fatal(
			"expected error creating CiString20 from over-length string",
		)
	}
}

func Test_IdTokenFromCiString_Empty(t *testing.T) {
	t.Parallel()

	str, err := st.NewCiString20Type("")
	if err != nil {
		t.Fatalf("failed to construct empty CiString20: %v", err)
	}

	_, err = mat.NewIdToken(str)
	if err == nil {
		t.Error("expected error from NewIdToken with empty CiString20")
	}
}
