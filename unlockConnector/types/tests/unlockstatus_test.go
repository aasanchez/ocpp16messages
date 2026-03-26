package types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
	ut "github.com/aasanchez/ocpp16messages/unlockconnector/types"
)

const (
	statusUnlockedStr     = "Unlocked"
	statusUnlockFailedStr = "UnlockFailed"
	statusNotSupportedStr = "NotSupported"

	unlockStatusStringMethod = "UnlockStatus.String()"
)

func TestUnlockStatus_IsValid_Unlocked(t *testing.T) {
	t.Parallel()

	if !ut.UnlockStatusUnlocked.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"UnlockStatusUnlocked",
		)
	}
}

func TestUnlockStatus_IsValid_UnlockFailed(t *testing.T) {
	t.Parallel()

	if !ut.UnlockStatusUnlockFailed.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"UnlockStatusUnlockFailed",
		)
	}
}

func TestUnlockStatus_IsValid_NotSupported(t *testing.T) {
	t.Parallel()

	if !ut.UnlockStatusNotSupported.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"UnlockStatusNotSupported",
		)
	}
}

func TestUnlockStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := ut.UnlockStatus("")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"UnlockStatus(\"\")",
		)
	}
}

func TestUnlockStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := ut.UnlockStatus("Unknown")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"UnlockStatus(\"Unknown\")",
		)
	}
}

func TestUnlockStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := ut.UnlockStatus("unlocked")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"UnlockStatus(\"unlocked\")",
		)
	}
}

func TestUnlockStatus_String_Unlocked(t *testing.T) {
	t.Parallel()

	got := ut.UnlockStatusUnlocked.String()
	if got != statusUnlockedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unlockStatusStringMethod,
			got,
			statusUnlockedStr,
		)
	}
}

func TestUnlockStatus_String_UnlockFailed(t *testing.T) {
	t.Parallel()

	got := ut.UnlockStatusUnlockFailed.String()
	if got != statusUnlockFailedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unlockStatusStringMethod,
			got,
			statusUnlockFailedStr,
		)
	}
}

func TestUnlockStatus_String_NotSupported(t *testing.T) {
	t.Parallel()

	got := ut.UnlockStatusNotSupported.String()
	if got != statusNotSupportedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			unlockStatusStringMethod,
			got,
			statusNotSupportedStr,
		)
	}
}
