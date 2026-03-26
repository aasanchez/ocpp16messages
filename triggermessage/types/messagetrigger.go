package types

// MessageTrigger represents the type of message to trigger from a Charge Point.
type MessageTrigger string

// Alias for shorter constant names within this package.
type mt = MessageTrigger

// MessageTrigger enumeration values as defined in OCPP 1.6.
const (
	// MessageTriggerBootNotification triggers a BootNotification message.
	MessageTriggerBootNotification mt = "BootNotification"
	// MessageTriggerDiagnosticsStatusNotification triggers a
	// DiagnosticsStatusNotification.
	//nolint:revive // Long constant name required for OCPP compliance
	MessageTriggerDiagnosticsStatusNotification mt = ("DiagnosticsStatusNotification")
	// MessageTriggerFirmwareStatusNotification triggers a
	// FirmwareStatusNotification.
	MessageTriggerFirmwareStatusNotification mt = "FirmwareStatusNotification"
	// MessageTriggerHeartbeat triggers a Heartbeat message.
	MessageTriggerHeartbeat mt = "Heartbeat"
	// MessageTriggerMeterValues triggers a MeterValues message.
	MessageTriggerMeterValues mt = "MeterValues"
	// MessageTriggerStatusNotification triggers a StatusNotification message.
	MessageTriggerStatusNotification mt = "StatusNotification"
)

// IsValid checks if the MessageTrigger value is valid per OCPP 1.6.
func (m MessageTrigger) IsValid() bool {
	switch m {
	case MessageTriggerBootNotification,
		MessageTriggerDiagnosticsStatusNotification,
		MessageTriggerFirmwareStatusNotification,
		MessageTriggerHeartbeat,
		MessageTriggerMeterValues,
		MessageTriggerStatusNotification:
		return true
	default:
		return false
	}
}

// String returns the string representation of MessageTrigger.
func (m MessageTrigger) String() string {
	return string(m)
}
