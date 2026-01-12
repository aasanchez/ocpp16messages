// UpdateFirmware handling rules:
//
// The Central System MAY instruct a Charge Point to update its firmware using
// UpdateFirmware.req. The request contains:
//   - The earliest date and time when the Charge Point is allowed to retrieve the firmware.
//   - The location (URL or path) from which the firmware can be downloaded.
//
// Charge Point behavior:
//   - Upon receipt of UpdateFirmware.req, the Charge Point SHALL respond with
//     UpdateFirmware.conf.
//   - The Charge Point SHOULD start retrieving the firmware as soon as possible
//     after the allowed retrieve-date.
//   - During download and installation, the Charge Point MUST send
//     FirmwareStatusNotification.req PDUs to keep the Central System updated on
//     progress.
//
// Installation rules:
//   - The Charge Point SHALL install the new firmware once it has verified the
//     firmware image is valid.
//   - If installation cannot safely occur during an ongoing charging session, it
//     is RECOMMENDED to wait until the session ends.
//   - Connectors not in use SHOULD be set to UNAVAILABLE while waiting for idle
//     conditions.
//
// Optional best practices:
//   - Rebooting the Charge Point after installation before sending "Installed"
//     status is recommended to verify connectivity and firmware integrity, but
//     is not mandatory.
package updateFirmware
