package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

func TestLocation_String_Body(t *testing.T) {
	t.Parallel()

	got := types.LocationBody.String()
	want := "Body"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestLocation_String_Cable(t *testing.T) {
	t.Parallel()

	got := types.LocationCable.String()
	want := "Cable"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestLocation_String_EV(t *testing.T) {
	t.Parallel()

	got := types.LocationEV.String()
	want := "EV"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestLocation_String_Inlet(t *testing.T) {
	t.Parallel()

	got := types.LocationInlet.String()
	want := "Inlet"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestLocation_String_Outlet(t *testing.T) {
	t.Parallel()

	got := types.LocationOutlet.String()
	want := "Outlet"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestLocation_IsValid_InvalidValue(t *testing.T) {
	t.Parallel()

	invalidLocation := types.Location("InvalidLocation")

	if invalidLocation.IsValid() {
		t.Error("Location.IsValid() = true, want false for invalid value")
	}
}
