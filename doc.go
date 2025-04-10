// Package ocpp16messages provides data types and validation logic for the Open Charge Point
// Protocol (OCPP) 1.6J. This package is focused on modeling the protocol message payloads,
// enforcing strict compliance with the specification's constraints such as maximum string
// lengths and printable ASCII requirements.
//
// This package is designed to be protocol-agnostic. It can be used independently of the
// serialization format (e.g., JSON, XML, SOAP), and is suitable for use in proxies, CSMS,
// test tools, and embedded implementations.
//
// All types are grouped by functional domain and modeled with validation-first behavior.
package ocpp16messages
