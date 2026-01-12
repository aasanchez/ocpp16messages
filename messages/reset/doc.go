// Package reset implements the OCPP 1.6 Reset message.
//
// # Handling Rules
//
// The Central System MAY request a Charge Point to reset by sending Reset.req.
// The request can specify a soft or hard reset. The Charge Point SHALL respond
// with Reset.conf indicating whether it will attempt the reset.
//
// # Transaction Handling
//
//   - Before performing any reset, the Charge Point SHALL send
//     StopTransaction.req for ongoing transactions. If StopTransaction.conf
//     is not received, the StopTransaction.req SHALL be queued.
//
// # Soft Reset
//
//   - Gracefully stop ongoing transactions and send StopTransaction.req
//     for each.
//   - Restart application software if possible, otherwise restart the
//     processor/controller.
//
// # Hard Reset
//
//   - Restart all hardware. Graceful stopping of ongoing transactions is
//     not required.
//   - If possible, send StopTransaction.req for previously ongoing
//     transactions after reboot and BootNotification.conf has been accepted.
//   - Use only as a last resort; queued information may be lost.
//
// # Persistent States
//
//   - Certain states, e.g., a connector set to Unavailable, SHALL persist
//     across resets.
package reset
