package types_test

import (
	"testing"

	ct "github.com/aasanchez/ocpp16messages/clearChargingProfile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	purposeChargePointMaxProfileStr = "ChargePointMaxProfile"
	purposeTxDefaultProfileStr      = "TxDefaultProfile"
	purposeTxProfileStr             = "TxProfile"
	purposeTypeString               = "ChargingProfilePurposeType.String()"
)

func TestChargingProfilePurposeType_IsValid_CPMaxProfile(t *testing.T) {
	t.Parallel()

	if !ct.ChargePointMaxProfile.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargePointMaxProfile")
	}
}

func TestChargingProfilePurposeType_IsValid_TxDefaultProfile(t *testing.T) {
	t.Parallel()

	if !ct.TxDefaultProfile.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "TxDefaultProfile")
	}
}

func TestChargingProfilePurposeType_IsValid_TxProfile(t *testing.T) {
	t.Parallel()

	if !ct.TxProfile.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "TxProfile")
	}
}

func TestChargingProfilePurposeType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	purpose := ct.ChargingProfilePurposeType("")
	if purpose.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargingProfilePurposeType(\"\")")
	}
}

func TestChargingProfilePurposeType_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	purpose := ct.ChargingProfilePurposeType("Invalid")
	if purpose.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargingProfilePurposeType(\"Invalid\")")
	}
}

func TestChargingProfilePurposeType_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	purpose := ct.ChargingProfilePurposeType("txprofile")
	if purpose.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ChargingProfilePurposeType(\"txprofile\")",
		)
	}
}

func TestChargingProfilePurposeType_String_ChargePointMaxProfile(t *testing.T) {
	t.Parallel()

	got := ct.ChargePointMaxProfile.String()
	if got != purposeChargePointMaxProfileStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			purposeTypeString,
			got,
			purposeChargePointMaxProfileStr,
		)
	}
}

func TestChargingProfilePurposeType_String_TxDefaultProfile(t *testing.T) {
	t.Parallel()

	got := ct.TxDefaultProfile.String()
	if got != purposeTxDefaultProfileStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			purposeTypeString,
			got,
			purposeTxDefaultProfileStr,
		)
	}
}

func TestChargingProfilePurposeType_String_TxProfile(t *testing.T) {
	t.Parallel()

	got := ct.TxProfile.String()
	if got != purposeTxProfileStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			purposeTypeString,
			got,
			purposeTxProfileStr,
		)
	}
}
