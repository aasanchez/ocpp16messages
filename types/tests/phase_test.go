package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

func TestPhase_String_L1(t *testing.T) {
	t.Parallel()

	got := types.PhaseL1.String()
	want := "L1"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestPhase_String_L2(t *testing.T) {
	t.Parallel()

	got := types.PhaseL2.String()
	want := "L2"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestPhase_String_L3(t *testing.T) {
	t.Parallel()

	got := types.PhaseL3.String()
	want := "L3"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestPhase_String_N(t *testing.T) {
	t.Parallel()

	got := types.PhaseN.String()
	want := "N"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestPhase_String_L1N(t *testing.T) {
	t.Parallel()

	got := types.PhaseL1N.String()
	want := "L1-N"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestPhase_IsValid_InvalidValue(t *testing.T) {
	t.Parallel()

	invalidPhase := types.Phase("InvalidPhase")

	if invalidPhase.IsValid() {
		t.Error("Phase.IsValid() = true, want false for invalid value")
	}
}
