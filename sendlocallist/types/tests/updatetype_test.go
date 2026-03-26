package types_test

import (
	"testing"

	slt "github.com/aasanchez/ocpp16messages/sendlocallist/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	typeFullStr         = "Full"
	typeDifferentialStr = "Differential"
)

func TestUpdateType_IsValid_Full(t *testing.T) {
	t.Parallel()

	if !slt.UpdateTypeFull.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"UpdateTypeFull",
		)
	}
}

func TestUpdateType_IsValid_Differential(t *testing.T) {
	t.Parallel()

	if !slt.UpdateTypeDifferential.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"UpdateTypeDifferential",
		)
	}
}

func TestUpdateType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	updateType := slt.UpdateType("")
	if updateType.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"UpdateType(\"\")",
		)
	}
}

func TestUpdateType_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	updateType := slt.UpdateType("Unknown")
	if updateType.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"UpdateType(\"Unknown\")",
		)
	}
}

func TestUpdateType_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	updateType := slt.UpdateType("full")
	if updateType.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"UpdateType(\"full\")",
		)
	}
}

func TestUpdateType_String_Full(t *testing.T) {
	t.Parallel()

	got := slt.UpdateTypeFull.String()
	if got != typeFullStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"UpdateType.String()",
			got,
			typeFullStr,
		)
	}
}

func TestUpdateType_String_Differential(t *testing.T) {
	t.Parallel()

	got := slt.UpdateTypeDifferential.String()
	if got != typeDifferentialStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"UpdateType.String()",
			got,
			typeDifferentialStr,
		)
	}
}
