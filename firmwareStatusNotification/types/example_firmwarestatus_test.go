package types_test

import (
	"fmt"

	fn "github.com/aasanchez/ocpp16messages/firmwareStatusNotification/types"
)

const (
	labelStatus  = "Status:"
	labelIsValid = "IsValid:"
)

// ExampleFirmwareStatusIdle demonstrates using the Idle status,
// sent when no firmware download/installation is in progress.
func ExampleFirmwareStatusIdle() {
	status := fn.FirmwareStatusIdle
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Idle
	// IsValid: true
}

// ExampleFirmwareStatusDownloading demonstrates using the Downloading status,
// sent when firmware download is in progress.
func ExampleFirmwareStatusDownloading() {
	status := fn.FirmwareStatusDownloading
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Downloading
	// IsValid: true
}

// ExampleFirmwareStatusDownloaded demonstrates using the Downloaded status,
// sent when firmware download has completed successfully.
func ExampleFirmwareStatusDownloaded() {
	status := fn.FirmwareStatusDownloaded
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Downloaded
	// IsValid: true
}

// ExampleFirmwareStatusDownloadFailed demonstrates using the DownloadFailed
// status, sent when firmware download has failed.
func ExampleFirmwareStatusDownloadFailed() {
	status := fn.FirmwareStatusDownloadFailed
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: DownloadFailed
	// IsValid: true
}

// ExampleFirmwareStatusInstalling demonstrates using the Installing status,
// sent when firmware installation is in progress.
func ExampleFirmwareStatusInstalling() {
	status := fn.FirmwareStatusInstalling
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Installing
	// IsValid: true
}

// ExampleFirmwareStatusInstalled demonstrates using the Installed status,
// sent when firmware installation has completed successfully.
func ExampleFirmwareStatusInstalled() {
	status := fn.FirmwareStatusInstalled
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Installed
	// IsValid: true
}

// ExampleFirmwareStatusInstallationFailed demonstrates using the
// InstallationFailed status, sent when firmware installation has failed.
func ExampleFirmwareStatusInstallationFailed() {
	status := fn.FirmwareStatusInstallationFailed
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: InstallationFailed
	// IsValid: true
}

// ExampleFirmwareStatus_IsValid_invalid demonstrates that invalid status
// values are rejected.
func ExampleFirmwareStatus_IsValid_invalid() {
	status := fn.FirmwareStatus("InvalidStatus")
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// IsValid: false
}
