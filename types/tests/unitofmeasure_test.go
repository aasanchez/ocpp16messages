package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

func TestUnitOfMeasure_String_Wh(t *testing.T) {
	t.Parallel()

	got := types.UnitWh.String()
	want := "Wh"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestUnitOfMeasure_String_KWh(t *testing.T) {
	t.Parallel()

	got := types.UnitKWh.String()
	want := "kWh"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestUnitOfMeasure_String_W(t *testing.T) {
	t.Parallel()

	got := types.UnitW.String()
	want := "W"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestUnitOfMeasure_String_A(t *testing.T) {
	t.Parallel()

	got := types.UnitA.String()
	want := "A"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestUnitOfMeasure_String_V(t *testing.T) {
	t.Parallel()

	got := types.UnitV.String()
	want := "V"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestUnitOfMeasure_String_Percent(t *testing.T) {
	t.Parallel()

	got := types.UnitPercent.String()
	want := "Percent"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestUnitOfMeasure_IsValid_InvalidValue(t *testing.T) {
	t.Parallel()

	invalidUnit := types.UnitOfMeasure("InvalidUnit")

	if invalidUnit.IsValid() {
		t.Error("UnitOfMeasure.IsValid() = true, want false for invalid value")
	}
}
