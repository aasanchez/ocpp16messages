// Package meterValues implements the OCPP 1.6 MeterValues message.
//
// # Handling Rules
//
// A Charge Point MAY sample its electrical meter or other sensor/transducer
// hardware to provide meter values. The timing of these samples can be
// configured using ChangeConfiguration.req to set acquisition intervals
// and specify the data to report.
//
// The Charge Point SHALL send MeterValues.req to offload meter data. Each
// request SHALL include the following for each sample:
//
//  1. connectorId: The ID of the connector from which samples were taken.
//     0 indicates the entire Charge Point. If Measurand is energy-related,
//     the sample SHOULD come from the main energy meter.
//  2. transactionId: Optional. The transaction associated with these values.
//     May be omitted if no transaction is in progress or if the main meter
//     is used.
//  3. meterValue: One or more MeterValue elements, each representing
//     measurements taken at a specific timestamp.
//
// Each MeterValue contains:
//   - timestamp: The time the data was captured.
//   - sampledValue: One or more individual measurement values.
//
// SampledValue optional fields:
//   - measurand: Type of measurement.
//   - context: Reason or event triggering the reading.
//   - location: Measurement location (e.g., Inlet, Outlet).
//   - phase: Electrical phase(s) the measurement applies to. Values SHALL
//     be reported from the meter/grid perspective. Not applicable to all
//     Measurands.
//   - unit: Measurement unit.
//   - format (EXPERIMENTAL): "Raw" (default) or "SignedData", a digitally
//     signed binary block (hex). May be deprecated in future versions.
//
// # Notes
//
//   - Two special Measurands, Current.Offered and Power.Offered, indicate
//     the maximum current/power offered to the EV for smart charging.
//   - For connector phase rotation, the Central System MAY query the
//     ConnectorPhaseRotation configuration key via GetConfiguration.
//     Values per connector include: NotApplicable, Unknown, RST, RTS, SRT,
//     STR, TRS, TSR.
//   - Default interpretation of a sampledValue without optional fields is
//     a register reading of active import energy in Wh (Watt-hour).
//
// Upon receipt of MeterValues.req, the Central System SHALL respond with
// MeterValues.conf. Sanity checks MAY be applied, but SHALL NOT prevent
// sending the confirmation. Failure to respond would cause the Charge Point
// to retry the message according to transaction-related error handling.
package meterValues
