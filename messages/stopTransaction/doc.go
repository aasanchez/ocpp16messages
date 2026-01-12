// StopTransaction handling rules:
//
// The Charge Point SHALL send StopTransaction.req to the Central System when a
// transaction is stopped. This request MAY include an optional TransactionData
// element containing MeterValues, using the same structure as MeterValues.req.
//
// Upon receipt, the Central System SHALL respond with StopTransaction.conf.
//
// Central System notes:
//   - Cannot prevent a transaction from stopping.
//   - MAY provide information about the idTag used to stop the transaction,
//     which SHOULD be used to update the Authorization Cache, if implemented.
//
// Charge Point behavior:
//   - idTag MAY be omitted if the Charge Point itself stops the transaction
//     (e.g., during a reset).
//   - If a transaction ends normally (e.g., EV driver presented idTag), Reason
//     MAY be omitted and assumed "Local"; otherwise, it SHOULD reflect the actual reason.
//   - The Charge Point SHALL unlock the cable (if not permanently attached) as part of
//     normal transaction termination.
//
// Cable-disconnection behavior:
//   - Controlled by configuration keys:
//   - StopTransactionOnEVSideDisconnect: If true, transaction stops on EV-side disconnect.
//   - UnlockConnectorOnEVSideDisconnect: If true, connector unlocks on EV-side disconnect.
//   - StopTransactionOnEVSideDisconnect = false has priority over UnlockConnectorOnEVSideDisconnect.
//   - Prevents unauthorized energy flow stoppage by unplugging the cable.
//
// Cache updates:
//   - If Authorization Cache is implemented, the Charge Point SHALL update the cache
//     with IdTagInfo from StopTransaction.conf if the idTag is not in the Local Authorization List.
//
// Notes:
//   - Central System MAY perform sanity checks on StopTransaction.req, but SHALL NOT
//     withhold StopTransaction.conf. Failing to respond triggers Charge Point retries
//     according to transaction-related error handling rules.
package stopTransaction
