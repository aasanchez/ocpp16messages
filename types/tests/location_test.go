package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

func checkLocation(t *testing.T, location types.Location, want string) {
	t.Helper()

	if location.String() != want {
		t.Errorf(types.ErrorMismatch, want, location.String())
	}

	if !location.IsValid() {
		t.Errorf("Location.IsValid() = false, want true for %s", location)
	}
}

func TestLocation_String_Body(t *testing.T) {
	t.Parallel()

	checkLocation(t, types.LocationBody, "Body")
}

func TestLocation_String_Cable(t *testing.T) {
	t.Parallel()

	checkLocation(t, types.LocationCable, "Cable")
}

func TestLocation_String_EV(t *testing.T) {
	t.Parallel()

	checkLocation(t, types.LocationEV, "EV")
}

func TestLocation_String_Inlet(t *testing.T) {
	t.Parallel()

	checkLocation(t, types.LocationInlet, "Inlet")
}

func TestLocation_String_Outlet(t *testing.T) {
	t.Parallel()

	checkLocation(t, types.LocationOutlet, "Outlet")
}

func TestLocation_IsValid_InvalidValue(t *testing.T) {
	t.Parallel()

	invalidLocation := types.Location("InvalidLocation")

	if invalidLocation.IsValid() {
		t.Error("Location.IsValid() = true, want false for invalid value")
	}
}
