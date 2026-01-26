// Package ocpp16messages provides Open Charge Point Protocol (OCPP) 1.6 message
// constructors and validation for Charge Points (EVSE) and Central Systems
// (CSMS/backends). It offers typed Req()/Conf() builders for all 28 OCPP 1.6
// operations, plus shared validated types for EV charging networks.
//
// Keywords: OCPP 1.6, Open Charge Point Protocol, EV charging, EVSE, CSMS,
// Charge Point, Central System, MeterValues, Authorize, BootNotification.
//
// This library implements strict type validation for OCPP 1.6 protocol fields,
// including CiString types (case-insensitive strings with length validation),
// DateTime types (RFC3339-compliant timestamps that must already be in UTC),
// and Integer types (validated uint16 values). All types use the constructor
// pattern with validation, returning errors for invalid inputs rather than
// panicking. Types are designed to be thread-safe with immutable fields and
// value receivers.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/aasanchez/ocpp16messages/authorize"
//	)
//
//	func main() {
//		req, err := authorize.Req(authorize.ReqInput{
//			IdTag: "RFID-ABC123",
//		})
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(req.IdTag.String())
//	}
package ocpp16messages
