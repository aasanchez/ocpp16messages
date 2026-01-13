package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/changeAvailability/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	typeOperative   = "Operative"
	typeInoperative = "Inoperative"
)

func TestAvailabilityType_IsValid_Operative(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityType(typeOperative)
	if !at.IsValid() {
		t.Errorf(st.ErrorMismatchValue, true, at.IsValid())
	}
}

func TestAvailabilityType_IsValid_Inoperative(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityType(typeInoperative)
	if !at.IsValid() {
		t.Errorf(st.ErrorMismatchValue, true, at.IsValid())
	}
}

func TestAvailabilityType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityType("")
	if at.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, at.IsValid())
	}
}

func TestAvailabilityType_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityType("Unknown")
	if at.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, at.IsValid())
	}
}

func TestAvailabilityType_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityType("operative")
	if at.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, at.IsValid())
	}
}

func TestAvailabilityType_String_Operative(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityTypeOperative
	if at.String() != typeOperative {
		t.Errorf(st.ErrorMismatchValue, typeOperative, at.String())
	}
}

func TestAvailabilityType_String_Inoperative(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityTypeInoperative
	if at.String() != typeInoperative {
		t.Errorf(st.ErrorMismatchValue, typeInoperative, at.String())
	}
}

func TestAvailabilityType_Constants(t *testing.T) {
	t.Parallel()

	if types.AvailabilityTypeOperative != typeOperative {
		t.Errorf(
			st.ErrorMismatchValue,
			typeOperative,
			types.AvailabilityTypeOperative,
		)
	}

	if types.AvailabilityTypeInoperative != typeInoperative {
		t.Errorf(
			st.ErrorMismatchValue,
			typeInoperative,
			types.AvailabilityTypeInoperative,
		)
	}
}
