package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/stoptransaction/types"
)

const errReasonStringFormat = "Reason.String() = %v, want %v"

func TestReason_DeAuthorized_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonDeAuthorized

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for DeAuthorized")
	}
}

func TestReason_EmergencyStop_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonEmergencyStop

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for EmergencyStop")
	}
}

func TestReason_EVDisconnected_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonEVDisconnected

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for EVDisconnected")
	}
}

func TestReason_HardReset_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonHardReset

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for HardReset")
	}
}

func TestReason_Local_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonLocal

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for Local")
	}
}

func TestReason_Other_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonOther

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for Other")
	}
}

func TestReason_PowerLoss_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonPowerLoss

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for PowerLoss")
	}
}

func TestReason_Reboot_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonReboot

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for Reboot")
	}
}

func TestReason_Remote_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonRemote

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for Remote")
	}
}

func TestReason_SoftReset_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonSoftReset

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for SoftReset")
	}
}

func TestReason_UnlockCommand_IsValid(t *testing.T) {
	t.Parallel()

	reason := types.ReasonUnlockCommand

	if !reason.IsValid() {
		t.Error("Reason.IsValid() = false, want true for UnlockCommand")
	}
}

func TestReason_Invalid_IsNotValid(t *testing.T) {
	t.Parallel()

	reason := types.Reason("InvalidReason")

	if reason.IsValid() {
		t.Error("Reason.IsValid() = true, want false for invalid reason")
	}
}

func TestReason_Empty_IsNotValid(t *testing.T) {
	t.Parallel()

	reason := types.Reason("")

	if reason.IsValid() {
		t.Error("Reason.IsValid() = true, want false for empty string")
	}
}

func TestReason_DeAuthorized_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonDeAuthorized
	expected := "DeAuthorized"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}

func TestReason_EmergencyStop_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonEmergencyStop
	expected := "EmergencyStop"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}

func TestReason_EVDisconnected_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonEVDisconnected
	expected := "EVDisconnected"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}

func TestReason_HardReset_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonHardReset
	expected := "HardReset"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}

func TestReason_Local_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonLocal
	expected := "Local"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}

func TestReason_Other_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonOther
	expected := "Other"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}

func TestReason_PowerLoss_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonPowerLoss
	expected := "PowerLoss"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}

func TestReason_Reboot_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonReboot
	expected := "Reboot"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}

func TestReason_Remote_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonRemote
	expected := "Remote"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}

func TestReason_SoftReset_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonSoftReset
	expected := "SoftReset"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}

func TestReason_UnlockCommand_String(t *testing.T) {
	t.Parallel()

	reason := types.ReasonUnlockCommand
	expected := "UnlockCommand"

	if reason.String() != expected {
		t.Errorf(errReasonStringFormat, reason.String(), expected)
	}
}
