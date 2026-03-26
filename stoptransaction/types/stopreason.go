package types

// Reason represents the reason for stopping a transaction as defined in
// OCPP 1.6 specification. The Charge Point SHALL send the reason if the
// transaction was not stopped normally (i.e., not stopped via EV driver
// presenting an idTag).
type Reason string

// Reason enumeration values as defined in OCPP 1.6.
const (
	// ReasonDeAuthorized indicates the transaction was stopped because of the
	// authorization status in a StartTransaction.conf.
	ReasonDeAuthorized Reason = "DeAuthorized"
	// ReasonEmergencyStop indicates an emergency stop button was used.
	ReasonEmergencyStop Reason = "EmergencyStop"
	// ReasonEVDisconnected indicates disconnection of cable on EV side.
	ReasonEVDisconnected Reason = "EVDisconnected"
	// ReasonHardReset indicates a hard reset command was received.
	ReasonHardReset Reason = "HardReset"
	// ReasonLocal indicates a normal stop by the EV driver presenting idTag.
	ReasonLocal Reason = "Local"
	// ReasonOther indicates an unspecified reason.
	ReasonOther Reason = "Other"
	// ReasonPowerLoss indicates a power outage occurred.
	ReasonPowerLoss Reason = "PowerLoss"
	// ReasonReboot indicates a local reboot command.
	ReasonReboot Reason = "Reboot"
	// ReasonRemote indicates a remote stop transaction request was received.
	ReasonRemote Reason = "Remote"
	// ReasonSoftReset indicates a soft reset command was received.
	ReasonSoftReset Reason = "SoftReset"
	// ReasonUnlockCommand indicates an unlock connector command was received.
	ReasonUnlockCommand Reason = "UnlockCommand"
)

// IsValid checks if the Reason value is valid per OCPP 1.6.
func (r Reason) IsValid() bool {
	switch r {
	case ReasonDeAuthorized,
		ReasonEmergencyStop,
		ReasonEVDisconnected,
		ReasonHardReset,
		ReasonLocal,
		ReasonOther,
		ReasonPowerLoss,
		ReasonReboot,
		ReasonRemote,
		ReasonSoftReset,
		ReasonUnlockCommand:
		return true
	default:
		return false
	}
}

// String returns the string representation of Reason.
func (r Reason) String() string {
	return string(r)
}
