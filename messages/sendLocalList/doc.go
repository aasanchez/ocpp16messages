// SendLocalList handling rules:
//
// The Central System MAY send a Local Authorization List to a Charge Point
// using SendLocalList.req. The list can be:
//   - Full: Replace the current list entirely.
//   - Differential: Apply updates to the existing list.
//
// The request SHALL include:
//   - updateType: Full or Differential.
//   - versionNumber: The version to associate with the local authorization list
//     after the update.
//
// Upon receipt, the Charge Point SHALL respond with SendLocalList.conf
// indicating whether the update was accepted.
//
// If the response status is Failed or VersionMismatch and the updateType
// was Differential, the Central System SHOULD retry by sending the full
// list with updateType set to Full.
package sendLocalList
