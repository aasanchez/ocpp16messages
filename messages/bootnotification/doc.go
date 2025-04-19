// Package bootnotification implements the BootNotification.req and BootNotification.conf
// messages as defined in the OCPP 1.6J specification.
//
// BootNotification is the first message a Charge Point must send after startup or reboot.
// It informs the Central System about the Charge Point’s identity, vendor, model,
// and software/firmware version.
//
// The Central System responds with a BootNotification.conf message containing a
// registration status and a heartbeat interval. This response determines whether the
// Charge Point can begin normal operation.
//
// # Protocol Rules:
//
//   - A Charge Point MUST send a BootNotification.req every time it powers on or reboots.
//   - Until it receives a response with status "Accepted" or "Pending", the Charge Point
//     MUST NOT send any other OCPP messages. This includes previously cached messages.
//   - If the response status is "Accepted":
//       - The Charge Point updates its heartbeat interval to match the value in the response.
//       - It is RECOMMENDED to synchronize the internal clock using the returned current time.
//   - If the status is "Rejected":
//       - The Charge Point MUST stop all communication until the retry interval expires.
//       - During this time, the connection may be closed by either side.
//       - The Charge Point MUST NOT respond to any messages from the Central System.
//   - If the status is "Pending":
//       - The Central System intends to perform additional configuration or checks before accepting.
//       - The connection SHOULD remain open.
//       - The Charge Point MUST NOT send any requests unless explicitly triggered by the Central System.
//
// While in Pending state, the following requests from the Central System are NOT allowed:
// RemoteStartTransaction.req and RemoteStopTransaction.req.
//
// This package provides message types, constructors, and validation tools for implementing
// a correct boot notification flow in line with OCPP 1.6J.

package bootnotification
