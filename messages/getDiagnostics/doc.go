// Package getDiagnostics implements the OCPP 1.6 GetDiagnostics message.
//
// # Handling Rules
//
// The Central System MAY request diagnostic information from a Charge Point
// by sending GetDiagnostics.req, including a location for uploading the data
// and optional begin and end timestamps.
//
// If diagnostic information is available, the Charge Point SHALL respond
// with GetDiagnostics.conf containing the name of the file to be uploaded.
// Only a single diagnostics file SHALL be uploaded, and its format is not
// prescribed.
//
// If no diagnostic information is available, the response SHALL NOT include
// a file name.
//
// During the upload, the Charge Point MUST send
// DiagnosticsStatusNotification.req messages to keep the Central System
// informed of the upload status.
package getDiagnostics
