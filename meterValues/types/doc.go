// Package types provides Open Charge Point Protocol (OCPP) 1.6
// MeterValues type definitions for EV charging.
//
// This package includes:
//   - ReadingContext: enumeration of reading context values
//     (Interruption.Begin, Interruption.End, Other, Sample.Clock,
//     Sample.Periodic, Transaction.Begin, Transaction.End, Trigger)
//   - ValueFormat: enumeration of value format options (Raw, SignedData)
//   - Measurand: enumeration of measurand types (Energy, Power, Current, etc.)
//   - Phase: enumeration of electrical phases (L1, L2, L3, N, L1-N, etc.)
//   - Location: enumeration of measurement locations
//     (Body, Cable, EV, Inlet, Outlet)
//   - UnitOfMeasure: enumeration of measurement units
//     (Wh, kWh, varh, kvarh, W, kW, VA, kVA, var, kvar, A, V, Celsius,
//     Fahrenheit, K, Percent)
//   - SampledValue: structure for individual measurement values
//   - MeterValue: structure for a set of measurements at a timestamp
package types
