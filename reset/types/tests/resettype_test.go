package types_test

import (
	"testing"

	rt "github.com/aasanchez/ocpp16messages/reset/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	typeHardStr = "Hard"
	typeSoftStr = "Soft"
)

func TestResetType_IsValid_Hard(t *testing.T) {
	t.Parallel()

	if !rt.ResetTypeHard.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"ResetTypeHard",
		)
	}
}

func TestResetType_IsValid_Soft(t *testing.T) {
	t.Parallel()

	if !rt.ResetTypeSoft.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"ResetTypeSoft",
		)
	}
}

func TestResetType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	resetType := rt.ResetType("")
	if resetType.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ResetType(\"\")",
		)
	}
}

func TestResetType_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	resetType := rt.ResetType("Unknown")
	if resetType.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ResetType(\"Unknown\")",
		)
	}
}

func TestResetType_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	resetType := rt.ResetType("hard")
	if resetType.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ResetType(\"hard\")",
		)
	}
}

func TestResetType_String_Hard(t *testing.T) {
	t.Parallel()

	got := rt.ResetTypeHard.String()
	if got != typeHardStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ResetType.String()",
			got,
			typeHardStr,
		)
	}
}

func TestResetType_String_Soft(t *testing.T) {
	t.Parallel()

	got := rt.ResetTypeSoft.String()
	if got != typeSoftStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ResetType.String()",
			got,
			typeSoftStr,
		)
	}
}
