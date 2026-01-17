package types_test

import (
	"testing"

	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	locBodyStr   = "Body"
	locCableStr  = "Cable"
	locEVStr     = "EV"
	locInletStr  = "Inlet"
	locOutletStr = "Outlet"
	locTypeStr   = "Location.String()"
)

func TestLocation_IsValid_Body(t *testing.T) {
	t.Parallel()

	if !mt.Body.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Body")
	}
}

func TestLocation_IsValid_Cable(t *testing.T) {
	t.Parallel()

	if !mt.Cable.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Cable")
	}
}

func TestLocation_IsValid_EV(t *testing.T) {
	t.Parallel()

	if !mt.EV.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "EV")
	}
}

func TestLocation_IsValid_Inlet(t *testing.T) {
	t.Parallel()

	if !mt.Inlet.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Inlet")
	}
}

func TestLocation_IsValid_Outlet(t *testing.T) {
	t.Parallel()

	if !mt.Outlet.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Outlet")
	}
}

func TestLocation_IsValid_Empty(t *testing.T) {
	t.Parallel()

	loc := mt.Location("")
	if loc.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Location(\"\")")
	}
}

func TestLocation_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	loc := mt.Location("Invalid")
	if loc.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Location(\"Invalid\")")
	}
}

func TestLocation_String_Body(t *testing.T) {
	t.Parallel()

	got := mt.Body.String()
	if got != locBodyStr {
		t.Errorf(st.ErrorMethodMismatch, locTypeStr, got, locBodyStr)
	}
}

func TestLocation_String_Cable(t *testing.T) {
	t.Parallel()

	got := mt.Cable.String()
	if got != locCableStr {
		t.Errorf(st.ErrorMethodMismatch, locTypeStr, got, locCableStr)
	}
}

func TestLocation_String_EV(t *testing.T) {
	t.Parallel()

	got := mt.EV.String()
	if got != locEVStr {
		t.Errorf(st.ErrorMethodMismatch, locTypeStr, got, locEVStr)
	}
}

func TestLocation_String_Inlet(t *testing.T) {
	t.Parallel()

	got := mt.Inlet.String()
	if got != locInletStr {
		t.Errorf(st.ErrorMethodMismatch, locTypeStr, got, locInletStr)
	}
}

func TestLocation_String_Outlet(t *testing.T) {
	t.Parallel()

	got := mt.Outlet.String()
	if got != locOutletStr {
		t.Errorf(st.ErrorMethodMismatch, locTypeStr, got, locOutletStr)
	}
}
