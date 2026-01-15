package types

// DiagnosticsStatus represents the status of a diagnostics upload
// as defined in OCPP 1.6.
type DiagnosticsStatus string

// DiagnosticsStatus enumeration values as defined in OCPP 1.6.
const (
	DiagnosticsStatusIdle         DiagnosticsStatus = "Idle"
	DiagnosticsStatusUploaded     DiagnosticsStatus = "Uploaded"
	DiagnosticsStatusUploadFailed DiagnosticsStatus = "UploadFailed"
	DiagnosticsStatusUploading    DiagnosticsStatus = "Uploading"
)

// IsValid checks if the DiagnosticsStatus value is valid per OCPP 1.6.
func (d DiagnosticsStatus) IsValid() bool {
	switch d {
	case DiagnosticsStatusIdle,
		DiagnosticsStatusUploaded,
		DiagnosticsStatusUploadFailed,
		DiagnosticsStatusUploading:
		return true
	default:
		return false
	}
}

// String returns the string representation of DiagnosticsStatus.
func (d DiagnosticsStatus) String() string {
	return string(d)
}
