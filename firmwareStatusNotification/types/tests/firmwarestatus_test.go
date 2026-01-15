package types_test

import (
	"testing"

	fn "github.com/aasanchez/ocpp16messages/firmwareStatusNotification/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusIdleStr               = "Idle"
	statusDownloadingStr        = "Downloading"
	statusDownloadedStr         = "Downloaded"
	statusDownloadFailedStr     = "DownloadFailed"
	statusInstallingStr         = "Installing"
	statusInstalledStr          = "Installed"
	statusInstallationFailedStr = "InstallationFailed"
	methodStringName            = "FirmwareStatus.String()"
)

func TestFirmwareStatus_IsValid_Idle(t *testing.T) {
	t.Parallel()

	if !fn.FirmwareStatusIdle.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "FirmwareStatusIdle")
	}
}

func TestFirmwareStatus_IsValid_Downloading(t *testing.T) {
	t.Parallel()

	if !fn.FirmwareStatusDownloading.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "FirmwareStatusDownloading")
	}
}

func TestFirmwareStatus_IsValid_Downloaded(t *testing.T) {
	t.Parallel()

	if !fn.FirmwareStatusDownloaded.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "FirmwareStatusDownloaded")
	}
}

func TestFirmwareStatus_IsValid_DownloadFailed(t *testing.T) {
	t.Parallel()

	if !fn.FirmwareStatusDownloadFailed.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "FirmwareStatusDownloadFailed")
	}
}

func TestFirmwareStatus_IsValid_Installing(t *testing.T) {
	t.Parallel()

	if !fn.FirmwareStatusInstalling.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "FirmwareStatusInstalling")
	}
}

func TestFirmwareStatus_IsValid_Installed(t *testing.T) {
	t.Parallel()

	if !fn.FirmwareStatusInstalled.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "FirmwareStatusInstalled")
	}
}

func TestFirmwareStatus_IsValid_InstallationFailed(t *testing.T) {
	t.Parallel()

	if !fn.FirmwareStatusInstallationFailed.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "FirmwareStatusInstallationFailed")
	}
}

func TestFirmwareStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := fn.FirmwareStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "FirmwareStatus(\"\")")
	}
}

func TestFirmwareStatus_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	status := fn.FirmwareStatus("Invalid")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "FirmwareStatus(\"Invalid\")")
	}
}

func TestFirmwareStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := fn.FirmwareStatus("idle")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "FirmwareStatus(\"idle\")")
	}
}

func TestFirmwareStatus_String_Idle(t *testing.T) {
	t.Parallel()

	got := fn.FirmwareStatusIdle.String()
	if got != statusIdleStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusIdleStr,
		)
	}
}

func TestFirmwareStatus_String_Downloading(t *testing.T) {
	t.Parallel()

	got := fn.FirmwareStatusDownloading.String()
	if got != statusDownloadingStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusDownloadingStr,
		)
	}
}

func TestFirmwareStatus_String_Downloaded(t *testing.T) {
	t.Parallel()

	got := fn.FirmwareStatusDownloaded.String()
	if got != statusDownloadedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusDownloadedStr,
		)
	}
}

func TestFirmwareStatus_String_DownloadFailed(t *testing.T) {
	t.Parallel()

	got := fn.FirmwareStatusDownloadFailed.String()
	if got != statusDownloadFailedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusDownloadFailedStr,
		)
	}
}

func TestFirmwareStatus_String_Installing(t *testing.T) {
	t.Parallel()

	got := fn.FirmwareStatusInstalling.String()
	if got != statusInstallingStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusInstallingStr,
		)
	}
}

func TestFirmwareStatus_String_Installed(t *testing.T) {
	t.Parallel()

	got := fn.FirmwareStatusInstalled.String()
	if got != statusInstalledStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusInstalledStr,
		)
	}
}

func TestFirmwareStatus_String_InstallationFailed(t *testing.T) {
	t.Parallel()

	got := fn.FirmwareStatusInstallationFailed.String()
	if got != statusInstallationFailedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusInstallationFailedStr,
		)
	}
}
