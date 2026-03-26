package types_test

import (
	"testing"

	tt "github.com/aasanchez/ocpp16messages/triggermessage/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	bootnotificationStr              = "BootNotification"
	diagnosticsstatusnotificationStr = "DiagnosticsStatusNotification"
	firmwarestatusnotificationStr    = "FirmwareStatusNotification"
	heartbeatStr                     = "Heartbeat"
	metervaluesStr                   = "MeterValues"
	statusnotificationStr            = "StatusNotification"
	methodStringName                 = "MessageTrigger.String()"
)

func TestMessageTrigger_IsValid_BootNotification(t *testing.T) {
	t.Parallel()

	if !tt.MessageTriggerBootNotification.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"MessageTriggerBootNotification",
		)
	}
}

func TestMessageTrigger_IsValid_DiagnosticsStatusNotification(t *testing.T) {
	t.Parallel()

	if !tt.MessageTriggerDiagnosticsStatusNotification.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"MessageTriggerDiagnosticsStatusNotification",
		)
	}
}

func TestMessageTrigger_IsValid_FirmwareStatusNotification(t *testing.T) {
	t.Parallel()

	if !tt.MessageTriggerFirmwareStatusNotification.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"MessageTriggerFirmwareStatusNotification",
		)
	}
}

func TestMessageTrigger_IsValid_Heartbeat(t *testing.T) {
	t.Parallel()

	if !tt.MessageTriggerHeartbeat.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"MessageTriggerHeartbeat",
		)
	}
}

func TestMessageTrigger_IsValid_MeterValues(t *testing.T) {
	t.Parallel()

	if !tt.MessageTriggerMeterValues.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"MessageTriggerMeterValues",
		)
	}
}

func TestMessageTrigger_IsValid_StatusNotification(t *testing.T) {
	t.Parallel()

	if !tt.MessageTriggerStatusNotification.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"MessageTriggerStatusNotification",
		)
	}
}

func TestMessageTrigger_IsValid_Empty(t *testing.T) {
	t.Parallel()

	trigger := tt.MessageTrigger("")
	if trigger.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"MessageTrigger(\"\")",
		)
	}
}

func TestMessageTrigger_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	trigger := tt.MessageTrigger("Unknown")
	if trigger.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"MessageTrigger(\"Unknown\")",
		)
	}
}

func TestMessageTrigger_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	trigger := tt.MessageTrigger("heartbeat")
	if trigger.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"MessageTrigger(\"heartbeat\")",
		)
	}
}

func TestMessageTrigger_String_BootNotification(t *testing.T) {
	t.Parallel()

	got := tt.MessageTriggerBootNotification.String()
	if got != bootnotificationStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			bootnotificationStr,
		)
	}
}

func TestMessageTrigger_String_DiagnosticsStatusNotification(t *testing.T) {
	t.Parallel()

	got := tt.MessageTriggerDiagnosticsStatusNotification.String()
	if got != diagnosticsstatusnotificationStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			diagnosticsstatusnotificationStr,
		)
	}
}

func TestMessageTrigger_String_FirmwareStatusNotification(t *testing.T) {
	t.Parallel()

	got := tt.MessageTriggerFirmwareStatusNotification.String()
	if got != firmwarestatusnotificationStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			firmwarestatusnotificationStr,
		)
	}
}

func TestMessageTrigger_String_Heartbeat(t *testing.T) {
	t.Parallel()

	got := tt.MessageTriggerHeartbeat.String()
	if got != heartbeatStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			heartbeatStr,
		)
	}
}

func TestMessageTrigger_String_MeterValues(t *testing.T) {
	t.Parallel()

	got := tt.MessageTriggerMeterValues.String()
	if got != metervaluesStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			metervaluesStr,
		)
	}
}

func TestMessageTrigger_String_StatusNotification(t *testing.T) {
	t.Parallel()

	got := tt.MessageTriggerStatusNotification.String()
	if got != statusnotificationStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusnotificationStr,
		)
	}
}
