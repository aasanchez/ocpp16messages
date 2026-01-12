// RemoteStartTransaction handling rules:
//
// The Central System MAY request a Charge Point to start a transaction by sending
// RemoteStartTransaction.req. The Charge Point SHALL respond with
// RemoteStartTransaction.conf, indicating whether it accepts the request and
// will attempt to start a transaction.
//
// Behavior depends on the AuthorizeRemoteTxRequests configuration key:
//   - true: The Charge Point SHALL attempt to authorize the idTag using the Local
//     Authorization List, Authorization Cache, or Authorize.req. The transaction
//     will only start after authorization.
//   - false: The Charge Point SHALL attempt to start the transaction immediately.
//     Authorization will be checked when processing the StartTransaction.req
//     sent to the Central System.
//
// Typical use cases:
//   - Allow a CPO operator to assist an EV driver in starting a transaction.
//   - Enable mobile apps to control charging transactions via the Central System.
//   - Enable SMS-based control of charging transactions.
//
// Requirements for RemoteStartTransaction.req:
//   - idTag: Mandatory. Used by the Charge Point to start the transaction and
//     send StartTransaction.req to the Central System.
//   - connectorId: Optional. If provided, the transaction starts on the specified
//     connector. If omitted, the Charge Point chooses the connector.
//     A Charge Point MAY reject requests without a connectorId.
//   - ChargingProfile: Optional. If provided and supported, SHALL be applied with
//     purpose set to TxProfile. Unsupported profiles SHOULD be ignored.
package remoteStartTransaction
