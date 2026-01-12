// Package unlockConnector implements the OCPP 1.6 UnlockConnector message.
//
// # Handling Rules
//
// The Central System MAY request a Charge Point to unlock a connector using
// UnlockConnector.req. This is intended to assist EV drivers who cannot unplug
// the cable due to a connector or cable retention malfunction.
//
// # Purpose
//
//   - Allows a CPO operator to manually trigger the unlock in case of a
//     connector malfunction, enabling the EV driver to remove the cable.
//   - SHOULD NOT be used to remotely stop a running transaction; use
//     RemoteStopTransaction for that purpose.
//
// # Charge Point Behavior
//
//   - Upon receipt of UnlockConnector.req, the Charge Point SHALL respond
//     with UnlockConnector.conf indicating whether it successfully unlocked
//     the connector.
//   - If a transaction is in progress on the connector, the Charge Point
//     SHALL finish the transaction first, as per StopTransaction rules.
//
// # Notes
//
//   - UnlockConnector.req only affects the cable retention lock on the
//     connector. It does not unlock a connector access door or other
//     enclosure.
package unlockConnector
