// Package types provides OCPP 1.6 StatusNotification message type definitions.
//
// This package includes:
//   - ChargePointStatus: enumeration of connector/Charge Point states
//     (Available, Preparing, Charging, SuspendedEV, SuspendedEVSE,
//     Finishing, Reserved, Unavailable, Faulted)
//   - ChargePointErrorCode: enumeration of error codes
//     (NoError, ConnectorLockFailure, EVCommunicationError, GroundFailure,
//     HighTemperature, InternalError, LocalListConflict, OtherError,
//     OverCurrentFailure, OverVoltage, PowerMeterFailure, PowerSwitchFailure,
//     ReaderFailure, ResetFailure, UnderVoltage, WeakSignal)
package types
