package types

// FirmwareStatus represents the status of a firmware download or installation
// as defined in OCPP 1.6.
type FirmwareStatus string

// FirmwareStatus enumeration values as defined in OCPP 1.6.
const (
	FirmwareStatusDownloaded         FirmwareStatus = "Downloaded"
	FirmwareStatusDownloadFailed     FirmwareStatus = "DownloadFailed"
	FirmwareStatusDownloading        FirmwareStatus = "Downloading"
	FirmwareStatusIdle               FirmwareStatus = "Idle"
	FirmwareStatusInstallationFailed FirmwareStatus = "InstallationFailed"
	FirmwareStatusInstalling         FirmwareStatus = "Installing"
	FirmwareStatusInstalled          FirmwareStatus = "Installed"
)

// IsValid checks if the FirmwareStatus value is valid per OCPP 1.6.
func (f FirmwareStatus) IsValid() bool {
	switch f {
	case FirmwareStatusDownloaded,
		FirmwareStatusDownloadFailed,
		FirmwareStatusDownloading,
		FirmwareStatusIdle,
		FirmwareStatusInstallationFailed,
		FirmwareStatusInstalling,
		FirmwareStatusInstalled:
		return true
	default:
		return false
	}
}

// String returns the string representation of FirmwareStatus.
func (f FirmwareStatus) String() string {
	return string(f)
}
