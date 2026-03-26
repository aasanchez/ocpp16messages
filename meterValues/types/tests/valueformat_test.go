package types_test

import (
	"testing"

	mt "github.com/aasanchez/ocpp16messages/metervalues/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	vfRawStr        = "Raw"
	vfSignedDataStr = "SignedData"
	vfTypeString    = "ValueFormat.String()"
)

func TestValueFormat_IsValid_Raw(t *testing.T) {
	t.Parallel()

	if !mt.Raw.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Raw")
	}
}

func TestValueFormat_IsValid_SignedData(t *testing.T) {
	t.Parallel()

	if !mt.SignedData.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "SignedData")
	}
}

func TestValueFormat_IsValid_Empty(t *testing.T) {
	t.Parallel()

	vf := mt.ValueFormat("")
	if vf.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ValueFormat(\"\")")
	}
}

func TestValueFormat_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	vf := mt.ValueFormat("Invalid")
	if vf.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ValueFormat(\"Invalid\")")
	}
}

func TestValueFormat_String_Raw(t *testing.T) {
	t.Parallel()

	got := mt.Raw.String()
	if got != vfRawStr {
		t.Errorf(st.ErrorMethodMismatch, vfTypeString, got, vfRawStr)
	}
}

func TestValueFormat_String_SignedData(t *testing.T) {
	t.Parallel()

	got := mt.SignedData.String()
	if got != vfSignedDataStr {
		t.Errorf(st.ErrorMethodMismatch, vfTypeString, got, vfSignedDataStr)
	}
}
