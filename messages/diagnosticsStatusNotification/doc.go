// GetDiagnostics handling rules:
//
// The Central System MAY request diagnostic information from a Charge Point
// by sending a GetDiagnostics.req, including a location where the diagnostics
// data MUST be uploaded and optional begin and end timestamps.
//
// Upon receipt, if diagnostic information is available, the Charge Point
// SHALL respond with GetDiagnostics.conf containing the name of the file
// that will be uploaded. A single diagnostics file SHALL be uploaded and
// its format is not prescribed.
//
// If no diagnostics information is available, the response SHALL NOT
// include a file name.
//
// During the upload process, the Charge Point MUST send
// DiagnosticsStatusNotification.req messages to keep the Central System
// informed of the upload status.
package diagnosticsStatusNotification
