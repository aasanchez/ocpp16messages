package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/changeConfiguration/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAccepted     = "Accepted"
	statusRejected     = "Rejected"
	statusReboot       = "RebootRequired"
	statusNotSupported = "NotSupported"
)

func TestConfigurationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatus(statusAccepted)
	if !cs.IsValid() {
		t.Errorf(st.ErrorMismatchValue, true, cs.IsValid())
	}
}

func TestConfigurationStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatus(statusRejected)
	if !cs.IsValid() {
		t.Errorf(st.ErrorMismatchValue, true, cs.IsValid())
	}
}

func TestConfigurationStatus_IsValid_RebootRequired(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatus(statusReboot)
	if !cs.IsValid() {
		t.Errorf(st.ErrorMismatchValue, true, cs.IsValid())
	}
}

func TestConfigurationStatus_IsValid_NotSupported(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatus(statusNotSupported)
	if !cs.IsValid() {
		t.Errorf(st.ErrorMismatchValue, true, cs.IsValid())
	}
}

func TestConfigurationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatus("")
	if cs.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, cs.IsValid())
	}
}

func TestConfigurationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatus("Unknown")
	if cs.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, cs.IsValid())
	}
}

func TestConfigurationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatus("accepted")
	if cs.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, cs.IsValid())
	}
}

func TestConfigurationStatus_IsValid_Scheduled(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatus("Scheduled")
	if cs.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, cs.IsValid())
	}
}

func TestConfigurationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatusAccepted
	if cs.String() != statusAccepted {
		t.Errorf(st.ErrorMismatchValue, statusAccepted, cs.String())
	}
}

func TestConfigurationStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatusRejected
	if cs.String() != statusRejected {
		t.Errorf(st.ErrorMismatchValue, statusRejected, cs.String())
	}
}

func TestConfigurationStatus_String_RebootRequired(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatusRebootRequired
	if cs.String() != statusReboot {
		t.Errorf(st.ErrorMismatchValue, statusReboot, cs.String())
	}
}

func TestConfigurationStatus_String_NotSupported(t *testing.T) {
	t.Parallel()

	cs := types.ConfigurationStatusNotSupported
	if cs.String() != statusNotSupported {
		t.Errorf(st.ErrorMismatchValue, statusNotSupported, cs.String())
	}
}

func TestConfigurationStatus_Constants(t *testing.T) {
	t.Parallel()

	if types.ConfigurationStatusAccepted != statusAccepted {
		t.Errorf(
			st.ErrorMismatchValue,
			statusAccepted,
			types.ConfigurationStatusAccepted,
		)
	}

	if types.ConfigurationStatusRejected != statusRejected {
		t.Errorf(
			st.ErrorMismatchValue,
			statusRejected,
			types.ConfigurationStatusRejected,
		)
	}

	if types.ConfigurationStatusRebootRequired != statusReboot {
		t.Errorf(
			st.ErrorMismatchValue,
			statusReboot,
			types.ConfigurationStatusRebootRequired,
		)
	}

	if types.ConfigurationStatusNotSupported != statusNotSupported {
		t.Errorf(
			st.ErrorMismatchValue,
			statusNotSupported,
			types.ConfigurationStatusNotSupported,
		)
	}
}
