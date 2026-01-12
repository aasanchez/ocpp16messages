// ChangeConfiguration handling rules:
//
// The Central System MAY request a Charge Point to change configuration
// parameters by sending a ChangeConfiguration.req containing a key-value
// pair, where "key" identifies the configuration setting and "value"
// specifies the new setting.
//
// Upon receipt, the Charge Point SHALL respond with ChangeConfiguration.conf
// indicating whether the configuration change could be applied.
//
// The response status SHALL be set as follows:
//   - Accepted:        Change applied successfully and effective immediately.
//   - RebootRequired:  Change applied successfully but requires a reboot
//     to become effective.
//   - NotSupported:    The specified key is not supported by the Charge Point.
//   - Rejected:        Change was not applied and none of the above statuses
//     apply (e.g. invalid format or out-of-range value).
//
// The content and format of "key" and "value" are not prescribed by OCPP.
//
// If a key represents a CSL (Comma-Separated List), it MAY be accompanied
// by a [KeyName]MaxLength configuration key indicating the maximum number
// of items allowed. If not present, a safe default of one (1) item SHOULD
// be assumed.
package changeConfiguration
