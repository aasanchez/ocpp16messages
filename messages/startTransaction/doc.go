// StartTransaction handling rules:
//
// The Charge Point SHALL send StartTransaction.req to the Central System to
// notify that a transaction has started. If the transaction ends a reservation,
// the StartTransaction.req MUST include the reservationId.
//
// Upon receipt, the Central System SHOULD respond with StartTransaction.conf,
// which SHALL include:
//   - transactionId: The assigned transaction identifier.
//   - authorizationStatus: Status of the idTag authorization.
//
// Central System responsibilities:
//   - MUST verify the identifier in StartTransaction.req, as it may have been
//     authorized locally by the Charge Point with outdated information.
//   - For example, the idTag may have been blocked since being added to the
//     Charge Point’s Authorization Cache.
//
// Charge Point responsibilities:
//   - If Authorization Cache is implemented, the Charge Point SHALL update the
//     cache entry with the IdTagInfo from StartTransaction.conf, if the idTag
//     is not in the Local Authorization List.
//
// Notes:
//   - The Central System may perform sanity checks on StartTransaction.req data,
//     but SHALL NOT withhold StartTransaction.conf based on these checks.
//   - Failing to respond with StartTransaction.conf will cause the Charge Point
//     to retry according to transaction-related error handling rules.

package startTransaction
