package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/changeAvailability/types"
)

const (
	errMismatch     = "Expected %v, got %v"
	typeOperative   = "Operative"
	typeInoperative = "Inoperative"
)

func TestAvailabilityType_IsValid_Operative(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityType(typeOperative)
	if !at.IsValid() {
		t.Errorf(errMismatch, true, at.IsValid())
	}
}

func TestAvailabilityType_IsValid_Inoperative(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityType(typeInoperative)
	if !at.IsValid() {
		t.Errorf(errMismatch, true, at.IsValid())
	}
}

func TestAvailabilityType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityType("")
	if at.IsValid() {
		t.Errorf(errMismatch, false, at.IsValid())
	}
}

func TestAvailabilityType_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityType("Unknown")
	if at.IsValid() {
		t.Errorf(errMismatch, false, at.IsValid())
	}
}

func TestAvailabilityType_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityType("operative")
	if at.IsValid() {
		t.Errorf(errMismatch, false, at.IsValid())
	}
}

func TestAvailabilityType_String_Operative(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityTypeOperative
	if at.String() != typeOperative {
		t.Errorf(errMismatch, typeOperative, at.String())
	}
}

func TestAvailabilityType_String_Inoperative(t *testing.T) {
	t.Parallel()

	at := types.AvailabilityTypeInoperative
	if at.String() != typeInoperative {
		t.Errorf(errMismatch, typeInoperative, at.String())
	}
}

func TestAvailabilityType_Constants(t *testing.T) {
	t.Parallel()

	if types.AvailabilityTypeOperative != typeOperative {
		t.Errorf(errMismatch, typeOperative, types.AvailabilityTypeOperative)
	}

	if types.AvailabilityTypeInoperative != typeInoperative {
		t.Errorf(
			errMismatch,
			typeInoperative,
			types.AvailabilityTypeInoperative,
		)
	}
}
