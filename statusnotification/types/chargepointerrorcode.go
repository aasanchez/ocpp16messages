package types

// ChargePointErrorCode represents the error code reported in a
// StatusNotification as defined in OCPP 1.6.
type ChargePointErrorCode string

// ChargePointErrorCode enumeration values as defined in OCPP 1.6.
const (
	ErrCodeConnectorLockFailure ChargePointErrorCode = "ConnectorLockFailure"
	ErrCodeEVCommunicationError ChargePointErrorCode = "EVCommunicationError"
	ErrCodeGroundFailure        ChargePointErrorCode = "GroundFailure"
	ErrCodeHighTemperature      ChargePointErrorCode = "HighTemperature"
	ErrCodeInternalError        ChargePointErrorCode = "InternalError"
	ErrCodeLocalListConflict    ChargePointErrorCode = "LocalListConflict"
	ErrCodeNoError              ChargePointErrorCode = "NoError"
	ErrCodeOtherError           ChargePointErrorCode = "OtherError"
	ErrCodeOverCurrentFailure   ChargePointErrorCode = "OverCurrentFailure"
	ErrCodeOverVoltage          ChargePointErrorCode = "OverVoltage"
	ErrCodePowerMeterFailure    ChargePointErrorCode = "PowerMeterFailure"
	ErrCodePowerSwitchFailure   ChargePointErrorCode = "PowerSwitchFailure"
	ErrCodeReaderFailure        ChargePointErrorCode = "ReaderFailure"
	ErrCodeResetFailure         ChargePointErrorCode = "ResetFailure"
	ErrCodeUnderVoltage         ChargePointErrorCode = "UnderVoltage"
	ErrCodeWeakSignal           ChargePointErrorCode = "WeakSignal"
)

// IsValid checks if the ChargePointErrorCode value is valid per OCPP 1.6.
func (c ChargePointErrorCode) IsValid() bool {
	switch c {
	case ErrCodeConnectorLockFailure,
		ErrCodeEVCommunicationError,
		ErrCodeGroundFailure,
		ErrCodeHighTemperature,
		ErrCodeInternalError,
		ErrCodeLocalListConflict,
		ErrCodeNoError,
		ErrCodeOtherError,
		ErrCodeOverCurrentFailure,
		ErrCodeOverVoltage,
		ErrCodePowerMeterFailure,
		ErrCodePowerSwitchFailure,
		ErrCodeReaderFailure,
		ErrCodeResetFailure,
		ErrCodeUnderVoltage,
		ErrCodeWeakSignal:
		return true
	default:
		return false
	}
}

// String returns the string representation of ChargePointErrorCode.
func (c ChargePointErrorCode) String() string {
	return string(c)
}
