package types

// DataTransferStatus represents the result of a DataTransfer request
// as defined in OCPP 1.6.
type DataTransferStatus string

// DataTransferStatus enumeration values as defined in OCPP 1.6.
const (
	DataTransferStatusAccepted         DataTransferStatus = "Accepted"
	DataTransferStatusRejected         DataTransferStatus = "Rejected"
	DataTransferStatusUnknownMessageId DataTransferStatus = "UnknownMessageId"
	DataTransferStatusUnknownVendor    DataTransferStatus = "UnknownVendor"
)

// IsValid checks if the DataTransferStatus value is valid per OCPP 1.6.
func (d DataTransferStatus) IsValid() bool {
	switch d {
	case DataTransferStatusAccepted,
		DataTransferStatusRejected,
		DataTransferStatusUnknownMessageId,
		DataTransferStatusUnknownVendor:
		return true
	default:
		return false
	}
}

// String returns the string representation of DataTransferStatus.
func (d DataTransferStatus) String() string {
	return string(d)
}
