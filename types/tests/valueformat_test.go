package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

func TestValueFormat_String_Raw(t *testing.T) {
	t.Parallel()

	got := types.ValueFormatRaw.String()
	want := "Raw"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestValueFormat_String_SignedData(t *testing.T) {
	t.Parallel()

	got := types.ValueFormatSignedData.String()
	want := "SignedData"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestValueFormat_IsValid_InvalidValue(t *testing.T) {
	t.Parallel()

	invalidFormat := types.ValueFormat("InvalidFormat")

	if invalidFormat.IsValid() {
		t.Error("ValueFormat.IsValid() = true, want false for invalid value")
	}
}
