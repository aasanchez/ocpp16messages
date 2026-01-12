// RemoteStopTransaction handling rules:
//
// The Central System MAY request a Charge Point to stop an ongoing transaction
// by sending RemoteStopTransaction.req. This is equivalent to a local action
// at the Charge Point to stop the transaction.
//
// Upon stopping the transaction, the Charge Point SHALL:
//   - Send StopTransaction.req to the Central System.
//   - Unlock the connector, if applicable.
//
// Typical use cases:
//   - Allow a CPO operator to assist an EV driver who has difficulty stopping a transaction.
//   - Enable mobile apps to control the transaction via the Central System.
package remoteStopTransaction
