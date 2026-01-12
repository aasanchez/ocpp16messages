// Package getLocalListVersion implements the OCPP 1.6 GetLocalListVersion
// message.
//
// # Handling Rules
//
// The Central System MAY request the version number of the Charge Point's
// Local Authorization List by sending GetLocalListVersion.req.
//
// Upon receipt, the Charge Point SHALL respond with GetLocalListVersion.conf
// containing the version number of its Local Authorization List.
//
// Version numbers have the following meaning:
//   - 0   : The local authorization list is empty.
//   - -1  : The Charge Point does not support Local Authorization Lists.
package getLocalListVersion
