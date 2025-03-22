// Package core provides shared data types, enumerations, validation utilities,
// and extensibility mechanisms used throughout the OCPP 1.6J message schema.
//
// This package is part of the ocpp16_messages module and is intended to be imported
// by other packages such as `authorize`, `bootnotification`, and other message types
// that rely on common structures defined in the OCPP 1.6J specification.
//
// Features included in this package:
//   - CiString types (CiString20Type, CiString25Type, etc.) for constrained-length strings
//   - The IdToken type used in Authorize.req and other messages
//   - Standard enums such as AuthorizationStatus
//   - Support for parsing and representing CALLERROR messages
//   - Plugin infrastructure for registering custom message validators
//   - Validation hook interface for observing validation outcomes
//
// This package is intentionally transport-agnostic and focused on schema compliance.
// It does not manage message transport, WebSockets, or SOAP/JSON specifics.
package core
