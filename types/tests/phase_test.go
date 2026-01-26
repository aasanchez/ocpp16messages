package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

const errPhaseValidFalse = "IsValid() = false, want true"

func TestPhase_String_L1(t *testing.T) {
	t.Parallel()

	got := types.PhaseL1.String()
	if got != "L1" {
		t.Errorf(types.ErrorMismatch, "L1", got)
	}

	if !types.PhaseL1.IsValid() {
		t.Error(errPhaseValidFalse)
	}
}

func TestPhase_String_L2(t *testing.T) {
	t.Parallel()

	got := types.PhaseL2.String()
	if got != "L2" {
		t.Errorf(types.ErrorMismatch, "L2", got)
	}

	if !types.PhaseL2.IsValid() {
		t.Error(errPhaseValidFalse)
	}
}

func TestPhase_String_L3(t *testing.T) {
	t.Parallel()

	got := types.PhaseL3.String()
	if got != "L3" {
		t.Errorf(types.ErrorMismatch, "L3", got)
	}

	if !types.PhaseL3.IsValid() {
		t.Error(errPhaseValidFalse)
	}
}

func TestPhase_String_N(t *testing.T) {
	t.Parallel()

	got := types.PhaseN.String()
	if got != "N" {
		t.Errorf(types.ErrorMismatch, "N", got)
	}

	if !types.PhaseN.IsValid() {
		t.Error(errPhaseValidFalse)
	}
}

func TestPhase_String_L1N(t *testing.T) {
	t.Parallel()

	got := types.PhaseL1N.String()
	if got != "L1-N" {
		t.Errorf(types.ErrorMismatch, "L1-N", got)
	}

	if !types.PhaseL1N.IsValid() {
		t.Error(errPhaseValidFalse)
	}
}

func TestPhase_IsValid_InvalidValue(t *testing.T) {
	t.Parallel()

	invalidPhase := types.Phase("InvalidPhase")

	if invalidPhase.IsValid() {
		t.Error("Phase.IsValid() = true, want false for invalid value")
	}
}
